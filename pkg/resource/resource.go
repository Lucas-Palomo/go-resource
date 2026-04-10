package resource

import (
	"encoding/json"
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

type Bundle struct {
	resourcesFolder string
	messages        map[language.Tag]map[string]string
	defaultLocale   language.Tag
	currentLocale   language.Tag
}

func NewBundle(resourcesFolder string, defaultLocale language.Tag) *Bundle {
	return &Bundle{
		resourcesFolder: resourcesFolder,
		messages:        make(map[language.Tag]map[string]string),
		defaultLocale:   defaultLocale,
	}
}

func (b *Bundle) loadResources(file string, lang language.Tag, engine bundleEngine) {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
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
		panic(fmt.Sprintf("Unknown bundle backend engine: %d", engine))
	}

	if err != nil {
		panic(err)
	}

	if b.messages[lang] == nil {
		b.messages[lang] = messages
	} else {
		maps.Copy(b.messages[lang], messages)
	}
}

func (b *Bundle) loadFolder(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			b.loadFolder(path.Join(dir, file.Name()))
			continue
		}

		lang, err := language.Parse(strings.Split(file.Name(), ".")[0])
		if err != nil {
			panic(err)
		}

		switch path.Ext(file.Name()) {
		case ".yaml", ".yml":
			b.loadResources(path.Join(dir, file.Name()), lang, yamlEngine)
		case ".json":
			b.loadResources(path.Join(dir, file.Name()), lang, jsonEngine)
		case ".toml":
			b.loadResources(path.Join(dir, file.Name()), lang, tomlEngine)
		default:
			panic(fmt.Sprintf("unsupported file extension: %s", file.Name()))
		}
	}
}

func (b *Bundle) Load() {
	b.loadFolder(b.resourcesFolder)
}

func (b *Bundle) GetWithLocale(locale language.Tag, id string, replacers ...any) string {
	if message, ok := b.messages[locale]; ok {
		if len(replacers) > 0 {
			return fmt.Sprintf(message[id], replacers...)
		}
		return fmt.Sprint(message[id])
	}

	return id
}

func (b *Bundle) Get(id string, replacers ...any) string {
	message := b.GetWithLocale(b.currentLocale, id, replacers...)
	if message == "" {
		message = b.GetWithLocale(b.defaultLocale, id, replacers...)
	}
	return message
}

func (b *Bundle) SetLocale(language language.Tag) {
	b.currentLocale = language
}
