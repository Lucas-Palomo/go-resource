# v1 Reference

[English](./v1-reference.md) | [Português (Brasil)](./v1-reference.pt_br.md)

## Package

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

## Main type

### `type Bundle`

Central structure of the library. It stores the configured resource root, the locale catalogs and the current/default locale state.

## Construction

### `func NewBundle(resourcesFolder string, defaultLocale language.Tag) *Bundle`

Creates a new bundle for the given directory and default locale.

Behavior in `v1.0.1`:
- initializes `currentLocale` with the default locale
- starts with an empty in-memory catalog
- keeps compatibility with existing construction code

## Loading

### `func (b *Bundle) Load()`

Legacy-compatible loading method.

In `v1.0.1`, this method no longer panics. It records the last loading error internally.
Use `Err()` right after calling it if you need to inspect the failure.

### `func (b *Bundle) LoadWithError() error`

Preferred loading method for new code. Returns any filesystem, locale parsing, decoder or unsupported-format error directly.

### `func (b *Bundle) Err() error`

Returns the last loading error captured by `Load()` or `LoadWithError()`.

## Lookup

### `func (b *Bundle) Get(id string, replacers ...any) string`

Looks up the key using the current locale. If the current locale does not contain the key, the default locale is consulted. If the key still does not exist, the key itself is returned.

### `func (b *Bundle) GetWithLocale(locale language.Tag, id string, replacers ...any) string`

Looks up the key for a specific locale without changing bundle state. If the key is missing, the key itself is returned.

## Runtime configuration

### `func (b *Bundle) SetLocale(locale language.Tag)`

Changes the current locale used by `Get()`.

## Supported formats

- `.json`
- `.yaml`
- `.yml`
- `.toml`

## File naming rule

v1 expects each resource file to be named as:

```text
<locale>.<extension>
```

Examples:
- `en.json`
- `pt_BR.yaml`
- `es.toml`

The first filename segment is parsed as the locale.

## Error model in `v1.0.1`

Loading errors now come back as `error` values instead of crashing the process through `panic`.
Typical failure cases:
- unreadable directory or file
- invalid locale in filename
- unsupported file extension
- invalid JSON / YAML / TOML payload

## Compatibility note

`Load()` was preserved to avoid forcing a breaking API change in a patch release.
For all new code, `LoadWithError()` is the recommended entrypoint.
