# v1 Reference

[English](./v1-reference.md) | [Português (Brasil)](./v1-reference.pt_br.md)

## Package

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

## What v1 is

v1 is the compatibility-focused line of `go-resource`.

It is intentionally small:

- filesystem-based loading
- flat string catalogs
- simple locale switching
- formatting through `fmt.Sprintf`

As of `v1.0.1`, the line is stable and no longer uses `panic` as the default loading behavior.

## Supported formats

- JSON
- YAML / YML
- TOML

v1 does **not** support:

- Java-style `.properties`
- nested resource objects
- namespace composition from directories or dotted filename segments

## Resource model

The v1 decoder path loads files into `map[string]string`.

Practical consequences:

- every value should be a string
- keys stay exactly as declared inside the file
- folder structure only organizes files on disk
- a name such as `en.errors.json` is accepted because the first segment is the locale, but `errors` does not become part of the runtime key

## Construction and loading

### `func NewBundle(resourcesFolder string, defaultLocale language.Tag) *Bundle`

Creates a bundle for a root folder and a default locale.

### `func (b *Bundle) Load()`

Compatibility method. In `v1.0.1`, it no longer panics by default. It stores the last loading error internally.

### `func (b *Bundle) LoadWithError() error`

Preferred loading method for new v1 code. Returns load errors directly.

### `func (b *Bundle) Err() error`

Returns the last loading error captured by `Load()` or `LoadWithError()`.

## Lookup API

### `func (b *Bundle) SetLocale(locale language.Tag)`

Changes the current locale.

### `func (b *Bundle) Get(id string, replacers ...any) string`

Resolution order in `v1.0.1`:

1. current locale
2. default locale
3. return the key itself if still unresolved

### `func (b *Bundle) GetWithLocale(locale language.Tag, id string, replacers ...any) string`

Resolves a key for a specific locale without changing bundle state.

## Error values

- `ErrUnknownBundleEngine`
- `ErrUnsupportedFileExt`
- `ErrInvalidResourceFileName`

## Behavioral notes

- duplicate keys are not exposed as a configurable public policy
- files loaded later can overwrite earlier keys in the same locale
- the package is best suited to smaller, simpler setups
- move to v2 when you need nested objects, `.properties`, `fs.FS`, or explicit policies
