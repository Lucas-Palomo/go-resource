# v1 Reference

[English](./v1-reference.md) | [Português (Brasil)](./v1-reference.pt_br.md)

## Package

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

## Purpose

The v1 line is the compatibility-focused API for consumers already using the original package. In `v1.0.1`, its main improvement is not a redesign. It is a contract correction: loading failures no longer crash the process by default.

## Main type

### `type Bundle`

Central structure of the library. It stores:

- the configured resource root
- locale catalogs loaded in memory
- the current locale
- the default fallback locale
- the last loading error captured by compatibility APIs

## Construction

### `func NewBundle(resourcesFolder string, defaultLocale language.Tag) *Bundle`

Creates a new bundle for the given resource root and default locale.

Behavior in `v1.0.1`:

- initializes `currentLocale` with the default locale
- starts with an empty in-memory catalog
- preserves compatibility with existing construction code

## Loading

### `func (b *Bundle) Load()`

Legacy-compatible loading method.

In `v1.0.1`, this method no longer uses `panic`. Instead, it records the last load error internally. Call `Err()` immediately after `Load()` if you need to inspect the failure.

```go
bundle.Load()
if err := bundle.Err(); err != nil {
	return err
}
```

### `func (b *Bundle) LoadWithError() error`

Preferred loading method for new code.

It returns filesystem, locale parsing, decoder, or unsupported-format errors directly:

```go
if err := bundle.LoadWithError(); err != nil {
	return err
}
```

### `func (b *Bundle) Err() error`

Returns the last loading error captured by `Load()` or `LoadWithError()`.

## Lookup

### `func (b *Bundle) Get(id string, replacers ...any) string`

Looks up a key using the current locale.

Resolution order in `v1.0.1`:

1. current locale
2. default locale
3. return the key itself if still unresolved

### `func (b *Bundle) GetWithLocale(locale language.Tag, id string, replacers ...any) string`

Looks up the key for a specific locale without changing bundle state.

If the key does not exist, the key itself is returned.

## Runtime configuration

### `func (b *Bundle) SetLocale(locale language.Tag)`

Changes the current locale used by `Get()`.

## Supported formats

- `.json`
- `.yaml`
- `.yml`
- `.toml`

## File naming rule

v1 expects each resource file to be named with the locale in the first filename segment.

Examples:

- `en.json`
- `pt_BR.yaml`
- `es.toml`

Nested directories are allowed, but the locale still comes from the file name.

## Error model in `v1.0.1`

Typical loading failure cases:

- unreadable directory or file
- invalid locale in filename
- unsupported file extension
- invalid JSON / YAML / TOML payload

The important point is that these failures now come back as `error` values instead of terminating the application through `panic`.

## Recommendation

Keep using v1 when compatibility matters more than architecture changes.

For new systems or deeper evolution, migrate toward `/v2`.
