package resource

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// PropertiesDecoder decodes Java-style .properties resource files.
type PropertiesDecoder struct{}

// Decode implements Decoder.
func (PropertiesDecoder) Decode(data []byte) (map[string]any, error) {
	return parseProperties(data)
}

func parseProperties(data []byte) (map[string]any, error) {
	out := make(map[string]any)
	for lineNumber, line := range propertiesLogicalLines(string(data)) {
		trimmed := trimPropertiesLeadingWhitespace(line)
		if trimmed == "" {
			continue
		}
		if trimmed[0] == '#' || trimmed[0] == '!' {
			continue
		}

		keyRaw, valueRaw := splitPropertyLine(trimmed)
		key, err := decodePropertiesEscapes(keyRaw)
		if err != nil {
			return nil, fmt.Errorf("decode properties key on line %d: %w", lineNumber+1, err)
		}
		value, err := decodePropertiesEscapes(valueRaw)
		if err != nil {
			return nil, fmt.Errorf("decode properties value on line %d: %w", lineNumber+1, err)
		}

		out[key] = value
	}

	return out, nil
}

func propertiesLogicalLines(input string) []string {
	physical := splitPhysicalLines(input)
	logical := make([]string, 0, len(physical))

	var current strings.Builder
	continuing := false

	for _, line := range physical {
		segment := line
		if continuing {
			segment = trimPropertiesLeadingWhitespace(segment)
		}

		if continued := propertiesContinues(segment); continued {
			current.WriteString(segment[:len(segment)-1])
			continuing = true
			continue
		}

		current.WriteString(segment)
		logical = append(logical, current.String())
		current.Reset()
		continuing = false
	}

	if continuing || current.Len() > 0 {
		logical = append(logical, current.String())
	}

	return logical
}

func splitPhysicalLines(input string) []string {
	lines := make([]string, 0, strings.Count(input, "\n")+1)
	start := 0
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '\n':
			lines = append(lines, input[start:i])
			start = i + 1
		case '\r':
			lines = append(lines, input[start:i])
			if i+1 < len(input) && input[i+1] == '\n' {
				i++
			}
			start = i + 1
		}
	}

	if start <= len(input) {
		lines = append(lines, input[start:])
	}

	return lines
}

func propertiesContinues(line string) bool {
	count := 0
	for i := len(line) - 1; i >= 0 && line[i] == '\\'; i-- {
		count++
	}
	return count%2 == 1
}

func trimPropertiesLeadingWhitespace(text string) string {
	for i, r := range text {
		if !isPropertiesWhitespace(r) {
			return text[i:]
		}
	}
	return ""
}

func isPropertiesWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\f'
}

func splitPropertyLine(line string) (string, string) {
	sep := -1
	hasExplicitSeparator := false
	escaped := false

	for i, r := range line {
		if escaped {
			escaped = false
			continue
		}
		if r == '\\' {
			escaped = true
			continue
		}
		if r == '=' || r == ':' {
			sep = i
			hasExplicitSeparator = true
			break
		}
		if isPropertiesWhitespace(r) {
			sep = i
			break
		}
	}

	if sep == -1 {
		return line, ""
	}

	key := line[:sep]
	valueStart := sep
	if hasExplicitSeparator {
		valueStart++
	} else {
		for valueStart < len(line) {
			r, size := utf8.DecodeRuneInString(line[valueStart:])
			if !isPropertiesWhitespace(r) {
				break
			}
			valueStart += size
		}
		if valueStart < len(line) && (line[valueStart] == '=' || line[valueStart] == ':') {
			valueStart++
		}
	}

	for valueStart < len(line) {
		r, size := utf8.DecodeRuneInString(line[valueStart:])
		if !isPropertiesWhitespace(r) {
			break
		}
		valueStart += size
	}

	return key, line[valueStart:]
}

func decodePropertiesEscapes(text string) (string, error) {
	var out strings.Builder
	out.Grow(len(text))

	for i := 0; i < len(text); i++ {
		if text[i] != '\\' {
			out.WriteByte(text[i])
			continue
		}

		i++
		if i >= len(text) {
			out.WriteByte('\\')
			break
		}

		switch text[i] {
		case 't':
			out.WriteByte('\t')
		case 'n':
			out.WriteByte('\n')
		case 'r':
			out.WriteByte('\r')
		case 'f':
			out.WriteByte('\f')
		case 'u':
			if i+4 >= len(text) {
				return "", fmt.Errorf("invalid unicode escape")
			}
			hex := text[i+1 : i+5]
			value, err := strconv.ParseUint(hex, 16, 16)
			if err != nil {
				return "", fmt.Errorf("invalid unicode escape %q", `\\u`+hex)
			}
			out.WriteRune(rune(value))
			i += 4
		default:
			out.WriteByte(text[i])
		}
	}

	return out.String(), nil
}
