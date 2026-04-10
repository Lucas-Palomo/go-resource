package resource

import (
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

// Decoder converts the raw content of a resource file into a generic map.
//
// Implementations should return objects that can be flattened into dot-notated keys.
// Scalars are allowed only at leaf nodes.
type Decoder interface {
	Decode(data []byte) (map[string]any, error)
}

// DecoderFunc adapts a function into a Decoder.
type DecoderFunc func(data []byte) (map[string]any, error)

// Decode implements Decoder.
func (fn DecoderFunc) Decode(data []byte) (map[string]any, error) {
	return fn(data)
}

// JSONDecoder decodes JSON resource files.
type JSONDecoder struct{}

// Decode implements Decoder.
func (JSONDecoder) Decode(data []byte) (map[string]any, error) {
	var out map[string]any
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}
	return out, nil
}

// YAMLDecoder decodes YAML resource files.
type YAMLDecoder struct{}

// Decode implements Decoder.
func (YAMLDecoder) Decode(data []byte) (map[string]any, error) {
	var out map[string]any
	if err := yaml.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("decode yaml: %w", err)
	}
	return out, nil
}

// TOMLDecoder decodes TOML resource files.
type TOMLDecoder struct{}

// Decode implements Decoder.
func (TOMLDecoder) Decode(data []byte) (map[string]any, error) {
	var out map[string]any
	if err := toml.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("decode toml: %w", err)
	}
	return out, nil
}
