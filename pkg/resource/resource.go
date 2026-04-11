package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

type bundleEngine int

const (
	jsonEngine bundleEngine = iota
	tomlEngine
	yamlEngine
)

var (
	ErrUnknownBundleEngine     = errors.New("resource: unknown bundle backend engine")
	ErrUnsupportedFileExt      = errors.New("resource: unsupported file extension")
	ErrInvalidResourceFileName = errors.New("resource: invalid resource file name")
)

type Bundle struct {
	resourcesFolder string
	messages        map[language.Tag]map[string]string
	defaultLocale   language.Tag
	currentLocale   language.Tag
	lastErr         error
}

func NewBundle(resourcesFolder string, defaultLocale language.Tag) *Bundle {
	return &Bundle{
		resourcesFolder: resourcesFolder,
		messages:        make(map[language.Tag]map[string]string),
		defaultLocale:   defaultLocale,
		currentLocale:   defaultLocale,
	}
}

func (b *Bundle) loadResources(store map[language.Tag]map[string]string, file string, lang language.Tag, engine bundleEngine) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("resource: read %q: %w", file, err)
	}

	messages := make(map[string]string)

	switch engine {
	case jsonEngine:
		err = json.Unmarshal(content, &messages)
	case tomlEngine:
		err = toml.Unmarshal(content, &messages)
	case yamlEngine:
		err = yaml.Unmarshal(content, &messages)
	default:
		return fmt.Errorf("%w: %d", ErrUnknownBundleEngine, engine)
	}

	if err != nil {
		return fmt.Errorf("resource: decode %q: %w", file, err)
	}

	if store[lang] == nil {
		store[lang] = messages
		return nil
	}

	maps.Copy(store[lang], messages)
	return nil
}

func (b *Bundle) loadFolder(store map[language.Tag]map[string]string, dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("resource: read dir %q: %w", dir, err)
	}

	for _, file := range files {
		fullPath := path.Join(dir, file.Name())

		if file.IsDir() {
			if err := b.loadFolder(store, fullPath); err != nil {
				return err
			}
			continue
		}

		parts := strings.Split(file.Name(), ".")
		if len(parts) < 2 || parts[0] == "" {
			return fmt.Errorf("%w: %s", ErrInvalidResourceFileName, file.Name())
		}

		lang, err := language.Parse(parts[0])
		if err != nil {
			return fmt.Errorf("resource: parse locale from %q: %w", file.Name(), err)
		}

		switch ext := path.Ext(file.Name()); ext {
		case ".yaml", ".yml":
			if err := b.loadResources(store, fullPath, lang, yamlEngine); err != nil {
				return err
			}
		case ".json":
			if err := b.loadResources(store, fullPath, lang, jsonEngine); err != nil {
				return err
			}
		case ".toml":
			if err := b.loadResources(store, fullPath, lang, tomlEngine); err != nil {
				return err
			}
		default:
			return fmt.Errorf("%w: %s", ErrUnsupportedFileExt, file.Name())
		}
	}

	return nil
}

// Load keeps source compatibility with the original v1 API.
//
// In v1.0.1 this method no longer panics. Any loading error is stored and can
// be inspected with Err(). New code should prefer LoadWithError().
func (b *Bundle) Load() {
	_ = b.LoadWithError()
}

// LoadWithError loads the configured resource tree and returns any error found
// during directory walking, locale parsing or file decoding.
func (b *Bundle) LoadWithError() error {
	loaded := make(map[language.Tag]map[string]string)

	if err := b.loadFolder(loaded, b.resourcesFolder); err != nil {
		b.lastErr = err
		return err
	}

	b.messages = loaded
	b.lastErr = nil
	return nil
}

// Err returns the last loading error captured by Load or LoadWithError.
func (b *Bundle) Err() error {
	return b.lastErr
}

func formatMessage(message string, replacers ...any) string {
	if len(replacers) > 0 {
		return fmt.Sprintf(message, replacers...)
	}

	return fmt.Sprint(message)
}

func (b *Bundle) lookup(locale language.Tag, id string) (string, bool) {
	messages, ok := b.messages[locale]
	if !ok {
		return "", false
	}

	message, ok := messages[id]
	if !ok {
		return "", false
	}

	return message, true
}

func (b *Bundle) GetWithLocale(locale language.Tag, id string, replacers ...any) string {
	message, ok := b.lookup(locale, id)
	if !ok {
		return id
	}

	return formatMessage(message, replacers...)
}

func (b *Bundle) Get(id string, replacers ...any) string {
	if message, ok := b.lookup(b.currentLocale, id); ok {
		return formatMessage(message, replacers...)
	}

	if message, ok := b.lookup(b.defaultLocale, id); ok {
		return formatMessage(message, replacers...)
	}

	return id
}

func (b *Bundle) SetLocale(locale language.Tag) {
	b.currentLocale = locale
}
