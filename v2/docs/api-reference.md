# v2 API Reference

[English](./api-reference.md) | [Português (Brasil)](./api-reference.pt_br.md)

## Package

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Main public types

- `Bundle`
- `Catalog`
- `Decoder`
- `DecoderFunc`
- `MissingKeyStrategy`
- `DuplicateKeyStrategy`

## Construction

### `func New(opts ...Option) *Bundle`

Defaults:

- fallback locale: `language.English`
- current locale: fallback locale when not explicitly set
- missing-key strategy: `ReturnKeyOnMissing`
- duplicate-key strategy: `ErrorOnDuplicate`

Built-in decoders:

- `.json`
- `.yaml`
- `.yml`
- `.toml`
- `.properties`

## Options

- `WithFallbackLocale(locale language.Tag)`
- `WithLocale(locale language.Tag)`
- `WithMissingKeyStrategy(strategy MissingKeyStrategy)`
- `WithDuplicateKeyStrategy(strategy DuplicateKeyStrategy)`

## Strategies

### Missing key

- `ReturnKeyOnMissing`: returns the original key
- `ReturnEmptyOnMissing`: returns `""`
- `ErrorOnMissing`: returns `ErrMissingKey` through `Lookup` and `LookupFor`

### Duplicate key

- `OverwriteOnDuplicate`: keeps the last loaded value
- `ErrorOnDuplicate`: aborts loading with `ErrDuplicateKey`

## Loading

- `LoadDir(root string) error`
- `LoadFS(fsys fs.FS, root string) error`
- `RegisterDecoder(ext string, decoder Decoder)`

Loading rules:

- the first filename segment is the locale
- directory names become namespace segments
- remaining filename segments become namespace segments
- nested objects are flattened into dot notation
- the runtime catalog is always `map[string]string`

Examples:

```text
resources/errors/en.json           -> namespace prefix: errors
resources/en.messages.checkout.toml -> namespace prefix: messages.checkout
```

## Lookup

- `Lookup(key string, args ...any) (string, error)`
- `LookupFor(locale language.Tag, key string, args ...any) (string, error)`
- `Get(key string, args ...any) string`
- `GetFor(locale language.Tag, key string, args ...any) string`
- `Has(key string) bool`
- `HasFor(locale language.Tag, key string) bool`

Behavior notes:

- `Lookup` and `LookupFor` are the explicit APIs
- `Get` and `GetFor` are convenience wrappers
- when formatting arguments are provided, values are formatted with `fmt.Sprintf`
- lookups use the requested locale, its parents, the fallback locale, the fallback parents, and finally `language.Und`

## Runtime helpers

- `SetLocale(locale language.Tag)`
- `Locale() language.Tag`
- `FallbackLocale() language.Tag`
- `Loaded() bool`
- `Locales() []language.Tag`
- `Catalog(locale language.Tag) Catalog`
- `Reset()`

## Error values

- `ErrBundleNotLoaded`
- `ErrUnsupportedFormat`
- `ErrInvalidResourceFileName`
- `ErrDuplicateKey`
- `ErrMissingKey`
- `ErrInvalidResourceValue`

Structured errors:

- `KeyNotFoundError`
- `DuplicateKeyError`
- `ResourceFileNameError`
