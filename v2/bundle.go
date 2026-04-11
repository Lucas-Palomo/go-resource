package resource

import (
	"errors"
	"fmt"
	"maps"
	"sort"
	"strings"
	"sync"

	"golang.org/x/text/language"
)

// Catalog stores a flat key/value map for one locale.
type Catalog map[string]string

// Clone returns a defensive copy of the catalog.
func (c Catalog) Clone() Catalog {
	return maps.Clone(c)
}

// Bundle stores decoded catalogs and runtime lookup configuration.
type Bundle struct {
	mu sync.RWMutex

	catalogs map[language.Tag]Catalog
	decoders map[string]Decoder
	loaded   bool

	fallbackLocale       language.Tag
	currentLocale        language.Tag
	missingKeyStrategy   MissingKeyStrategy
	duplicateKeyStrategy DuplicateKeyStrategy
}

// New creates a new resource bundle with sane defaults and built-in decoders.
func New(opts ...Option) *Bundle {
	cfg := defaultConfig()
	for _, opt := range opts {
		if opt != nil {
			opt(&cfg)
		}
	}

	b := &Bundle{
		catalogs:             make(map[language.Tag]Catalog),
		decoders:             make(map[string]Decoder),
		fallbackLocale:       cfg.fallbackLocale,
		currentLocale:        cfg.currentLocale,
		missingKeyStrategy:   cfg.missingKeyStrategy,
		duplicateKeyStrategy: cfg.duplicateKeyStrategy,
	}

	if b.currentLocale == language.Und {
		b.currentLocale = b.fallbackLocale
	}

	b.RegisterDecoder(".json", JSONDecoder{})
	b.RegisterDecoder(".yaml", YAMLDecoder{})
	b.RegisterDecoder(".yml", YAMLDecoder{})
	b.RegisterDecoder(".toml", TOMLDecoder{})
	b.RegisterDecoder(".properties", PropertiesDecoder{})

	return b
}

// RegisterDecoder registers a decoder for a file extension such as ".json" or "yaml".
func (b *Bundle) RegisterDecoder(ext string, decoder Decoder) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.decoders[normalizeExt(ext)] = decoder
}

// SetLocale changes the current lookup locale.
func (b *Bundle) SetLocale(locale language.Tag) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.currentLocale = locale
}

// Locale returns the current lookup locale.
func (b *Bundle) Locale() language.Tag {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.currentLocale
}

// FallbackLocale returns the configured default fallback locale.
func (b *Bundle) FallbackLocale() language.Tag {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.fallbackLocale
}

// Loaded reports whether at least one successful load operation has completed.
func (b *Bundle) Loaded() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.loaded
}

// Locales returns the set of loaded locales sorted lexicographically.
func (b *Bundle) Locales() []language.Tag {
	b.mu.RLock()
	defer b.mu.RUnlock()

	locales := make([]language.Tag, 0, len(b.catalogs))
	for locale := range b.catalogs {
		locales = append(locales, locale)
	}

	sort.Slice(locales, func(i, j int) bool {
		return locales[i].String() < locales[j].String()
	})

	return locales
}

// Catalog returns a defensive copy of the flat catalog for a locale.
func (b *Bundle) Catalog(locale language.Tag) Catalog {
	b.mu.RLock()
	defer b.mu.RUnlock()

	catalog, ok := b.catalogs[locale]
	if !ok {
		return nil
	}

	return catalog.Clone()
}

// Reset clears all loaded catalogs while preserving configuration and registered decoders.
func (b *Bundle) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.catalogs = make(map[language.Tag]Catalog)
	b.loaded = false
}

// Lookup resolves a key using the current locale and the configured fallback chain.
func (b *Bundle) Lookup(key string, args ...any) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.lookupLocked(b.currentLocale, key, args...)
}

// LookupFor resolves a key for a specific locale without mutating the bundle state.
func (b *Bundle) LookupFor(locale language.Tag, key string, args ...any) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.lookupLocked(locale, key, args...)
}

// Has reports whether a key is resolvable using the current locale and fallback chain.
func (b *Bundle) Has(key string) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if !b.loaded {
		return false
	}

	_, ok := b.resolveLocked(b.currentLocale, key)
	return ok
}

// HasFor reports whether a key is resolvable for a specific locale and fallback chain.
func (b *Bundle) HasFor(locale language.Tag, key string) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if !b.loaded {
		return false
	}

	_, ok := b.resolveLocked(locale, key)
	return ok
}

// Get is a convenience wrapper around Lookup. It preserves the old ergonomic style while avoiding panics.
func (b *Bundle) Get(key string, args ...any) string {
	value, err := b.Lookup(key, args...)
	if err != nil {
		if errors.Is(err, ErrMissingKey) {
			return key
		}
		return key
	}

	return value
}

// GetFor is a convenience wrapper around LookupFor.
func (b *Bundle) GetFor(locale language.Tag, key string, args ...any) string {
	value, err := b.LookupFor(locale, key, args...)
	if err != nil {
		if errors.Is(err, ErrMissingKey) {
			return key
		}
		return key
	}

	return value
}

func (b *Bundle) lookupLocked(locale language.Tag, key string, args ...any) (string, error) {
	if !b.loaded {
		return "", ErrBundleNotLoaded
	}

	if value, ok := b.resolveLocked(locale, key); ok {
		if len(args) == 0 {
			return value, nil
		}
		return fmt.Sprintf(value, args...), nil
	}

	switch b.missingKeyStrategy {
	case ReturnEmptyOnMissing:
		return "", nil
	case ErrorOnMissing:
		return "", KeyNotFoundError{
			Key:      key,
			Locale:   locale,
			Fallback: b.fallbackLocale,
		}
	default:
		return key, nil
	}
}

func (b *Bundle) resolveLocked(locale language.Tag, key string) (string, bool) {
	for _, candidate := range localeChain(locale, b.fallbackLocale) {
		catalog, ok := b.catalogs[candidate]
		if !ok {
			continue
		}

		value, ok := catalog[key]
		if ok {
			return value, true
		}
	}

	return "", false
}

func localeChain(locale language.Tag, fallback language.Tag) []language.Tag {
	seen := make(map[string]struct{})
	chain := make([]language.Tag, 0, 8)

	appendLocaleTree := func(tag language.Tag) {
		for current := tag; current != language.Und; current = current.Parent() {
			id := current.String()
			if _, ok := seen[id]; ok {
				break
			}

			seen[id] = struct{}{}
			chain = append(chain, current)
		}
	}

	appendLocaleTree(locale)
	appendLocaleTree(fallback)

	if _, ok := seen[language.Und.String()]; !ok {
		chain = append(chain, language.Und)
	}

	return chain
}

func joinKey(parts ...string) string {
	filtered := parts[:0]
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			filtered = append(filtered, part)
		}
	}

	return strings.Join(filtered, ".")
}
