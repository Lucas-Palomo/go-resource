package resource

import (
	"strings"

	"golang.org/x/text/language"
)

// MissingKeyStrategy defines how the bundle behaves when a key cannot be resolved.
type MissingKeyStrategy int

const (
	// ReturnKeyOnMissing returns the original key. This mirrors the common i18n fallback pattern.
	ReturnKeyOnMissing MissingKeyStrategy = iota
	// ReturnEmptyOnMissing returns an empty string.
	ReturnEmptyOnMissing
	// ErrorOnMissing returns ErrMissingKey.
	ErrorOnMissing
)

// DuplicateKeyStrategy defines how the bundle behaves when the same key is loaded twice.
type DuplicateKeyStrategy int

const (
	// OverwriteOnDuplicate keeps the most recently loaded value.
	OverwriteOnDuplicate DuplicateKeyStrategy = iota
	// ErrorOnDuplicate aborts loading with ErrDuplicateKey.
	ErrorOnDuplicate
)

// Option configures a Bundle.
type Option func(*config)

type config struct {
	fallbackLocale       language.Tag
	currentLocale        language.Tag
	missingKeyStrategy   MissingKeyStrategy
	duplicateKeyStrategy DuplicateKeyStrategy
}

func defaultConfig() config {
	return config{
		fallbackLocale:       language.English,
		currentLocale:        language.Und,
		missingKeyStrategy:   ReturnKeyOnMissing,
		duplicateKeyStrategy: ErrorOnDuplicate,
	}
}

// WithFallbackLocale sets the bundle fallback locale.
func WithFallbackLocale(locale language.Tag) Option {
	return func(cfg *config) {
		cfg.fallbackLocale = locale
	}
}

// WithLocale sets the initial current locale.
func WithLocale(locale language.Tag) Option {
	return func(cfg *config) {
		cfg.currentLocale = locale
	}
}

// WithMissingKeyStrategy configures the behavior when a key is not found.
func WithMissingKeyStrategy(strategy MissingKeyStrategy) Option {
	return func(cfg *config) {
		cfg.missingKeyStrategy = strategy
	}
}

// WithDuplicateKeyStrategy configures the behavior when duplicate keys are found while loading.
func WithDuplicateKeyStrategy(strategy DuplicateKeyStrategy) Option {
	return func(cfg *config) {
		cfg.duplicateKeyStrategy = strategy
	}
}

func normalizeExt(ext string) string {
	ext = strings.TrimSpace(strings.ToLower(ext))
	if ext == "" {
		return ext
	}
	if strings.HasPrefix(ext, ".") {
		return ext
	}
	return "." + ext
}
