package resource

import (
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path"
	"strings"

	"golang.org/x/text/language"
)

// LoadDir loads all resources from a directory in the operating system filesystem.
func (b *Bundle) LoadDir(root string) error {
	return b.LoadFS(os.DirFS(root), ".")
}

// LoadFS loads all resources from any fs.FS, including embed.FS.
//
// Repeated successful calls merge into the existing in-memory catalogs.
// Call Reset() first when replacement semantics are required.
func (b *Bundle) LoadFS(fsys fs.FS, root string) error {
	decoders := b.snapshotDecoders()
	duplicateKeyStrategy := b.snapshotDuplicateKeyStrategy()

	pending := make(map[language.Tag]Catalog)
	origins := make(map[language.Tag]map[string]string)

	err := fs.WalkDir(fsys, root, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}

		ext := strings.ToLower(path.Ext(entry.Name()))
		decoder, ok := decoders[ext]
		if !ok || decoder == nil {
			return fmt.Errorf("%w: %s", ErrUnsupportedFormat, filePath)
		}

		locale, namespace, err := parseResourcePath(root, filePath)
		if err != nil {
			return err
		}

		data, err := fs.ReadFile(fsys, filePath)
		if err != nil {
			return err
		}

		decoded, err := decoder.Decode(data)
		if err != nil {
			return fmt.Errorf("decode %s: %w", filePath, err)
		}

		flat := make(Catalog)
		if err := flattenMap(namespace, decoded, flat); err != nil {
			return fmt.Errorf("flatten %s: %w", filePath, err)
		}

		if pending[locale] == nil {
			pending[locale] = make(Catalog)
		}
		if origins[locale] == nil {
			origins[locale] = make(map[string]string)
		}

		for key, value := range flat {
			if _, exists := pending[locale][key]; exists && duplicateKeyStrategy == ErrorOnDuplicate {
				return DuplicateKeyError{Locale: locale, Key: key, File: filePath}
			}
			pending[locale][key] = value
			origins[locale][key] = filePath
		}

		return nil
	})
	if err != nil {
		return err
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	if duplicateKeyStrategy == ErrorOnDuplicate {
		for locale, catalog := range pending {
			existing := b.catalogs[locale]
			for key := range catalog {
				if _, exists := existing[key]; exists {
					file := root
					if origin, ok := origins[locale][key]; ok {
						file = origin
					}
					return DuplicateKeyError{Locale: locale, Key: key, File: file}
				}
			}
		}
	}

	for locale, catalog := range pending {
		if b.catalogs[locale] == nil {
			b.catalogs[locale] = make(Catalog)
		}
		maps.Copy(b.catalogs[locale], catalog)
	}

	b.loaded = len(b.catalogs) > 0
	if b.currentLocale == language.Und {
		b.currentLocale = b.fallbackLocale
	}

	return nil
}

func (b *Bundle) snapshotDecoders() map[string]Decoder {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return maps.Clone(b.decoders)
}

func (b *Bundle) snapshotDuplicateKeyStrategy() DuplicateKeyStrategy {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.duplicateKeyStrategy
}

func parseResourcePath(root string, filePath string) (language.Tag, string, error) {
	relPath := strings.TrimPrefix(filePath, strings.TrimSuffix(root, "/")+"/")
	if root == "." || root == "" {
		relPath = filePath
	}

	dir := path.Dir(relPath)
	if dir == "." {
		dir = ""
	}

	base := path.Base(relPath)
	ext := path.Ext(base)
	name := strings.TrimSuffix(base, ext)
	parts := strings.Split(name, ".")
	if len(parts) == 0 || strings.TrimSpace(parts[0]) == "" {
		return language.Und, "", ResourceFileNameError{
			Path:   filePath,
			Reason: "missing locale segment",
		}
	}

	localeText := strings.ReplaceAll(parts[0], "_", "-")
	locale, err := language.Parse(localeText)
	if err != nil {
		return language.Und, "", ResourceFileNameError{
			Path:   filePath,
			Reason: err.Error(),
		}
	}

	namespaceParts := make([]string, 0, 8)
	if dir != "" {
		namespaceParts = append(namespaceParts, strings.Split(dir, "/")...)
	}
	if len(parts) > 1 {
		namespaceParts = append(namespaceParts, parts[1:]...)
	}

	return locale, joinKey(namespaceParts...), nil
}

func flattenMap(prefix string, input map[string]any, out Catalog) error {
	for key, value := range input {
		currentKey := joinKey(prefix, key)
		if err := flattenValue(currentKey, value, out); err != nil {
			return err
		}
	}
	return nil
}

func flattenValue(fullKey string, value any, out Catalog) error {
	switch typed := value.(type) {
	case map[string]any:
		return flattenMap(fullKey, typed, out)
	case map[any]any:
		converted := make(map[string]any, len(typed))
		for rawKey, rawValue := range typed {
			textKey, ok := rawKey.(string)
			if !ok {
				return fmt.Errorf("%w: key=%q type=%T", ErrInvalidResourceValue, fullKey, rawKey)
			}
			converted[textKey] = rawValue
		}
		return flattenMap(fullKey, converted, out)
	case string:
		out[fullKey] = typed
		return nil
	case fmt.Stringer:
		out[fullKey] = typed.String()
		return nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		bool:
		out[fullKey] = fmt.Sprint(typed)
		return nil
	case nil:
		return fmt.Errorf("%w: key=%q nil values are not supported", ErrInvalidResourceValue, fullKey)
	default:
		return fmt.Errorf("%w: key=%q type=%T", ErrInvalidResourceValue, fullKey, value)
	}
}
