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

Creates a new bundle with the built-in decoders and the default runtime strategies.

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

### Missing key strategies

- `ReturnKeyOnMissing`: returns the original key
- `ReturnEmptyOnMissing`: returns `""`
- `ErrorOnMissing`: returns `ErrMissingKey` through `Lookup` and `LookupFor`

### Duplicate key strategies

- `OverwriteOnDuplicate`: keeps the last loaded value
- `ErrorOnDuplicate`: aborts loading with `ErrDuplicateKey`

## Loading API

- `LoadDir(root string) error`
- `LoadFS(fsys fs.FS, root string) error`
- `RegisterDecoder(ext string, decoder Decoder)`

### Loading rules

- the first filename segment is the locale
- directory names become namespace segments
- remaining filename segments become namespace segments
- nested objects are flattened into dot notation
- the runtime catalog is always `map[string]string`
- repeated successful loads merge into the catalogs already in memory
- use `Reset()` when you want replacement instead of merge
- files with unsupported extensions abort loading with `ErrUnsupportedFormat`
- `RegisterDecoder` normalizes extensions such as `json` into `.json`
- `RegisterDecoder` ignores empty extensions and nil decoders

Examples:

```text
resources/errors/en.json            -> namespace prefix: errors
resources/en.messages.checkout.toml -> namespace prefix: messages.checkout
```

## Lookup API

- `Lookup(key string, args ...any) (string, error)`
- `LookupFor(locale language.Tag, key string, args ...any) (string, error)`
- `Get(key string, args ...any) string`
- `GetFor(locale language.Tag, key string, args ...any) string`
- `Has(key string) bool`
- `HasFor(locale language.Tag, key string) bool`

### Behavior notes

- `Lookup` and `LookupFor` are the strict APIs
- `Get` and `GetFor` are lenient helpers
- `Get` and `GetFor` return the original key on lookup errors, including missing keys and `ErrBundleNotLoaded`
- when formatting arguments are provided, values are formatted with `fmt.Sprintf`
- lookups traverse the requested locale, its parents, the fallback locale, the fallback parents, and finally `language.Und`

## Runtime helpers

- `SetLocale(locale language.Tag)`
- `Locale() language.Tag`
- `FallbackLocale() language.Tag`
- `Loaded() bool`
- `Locales() []language.Tag`
- `Catalog(locale language.Tag) Catalog`
- `Reset()`

### Helper semantics

- `SetLocale` changes the bundle default locale used by `Lookup` and `Get`
- `Catalog` returns a copy of the locale catalog, not the internal map itself
- `Locales` returns the locales currently loaded in memory
- `Reset` clears loaded catalogs and marks the bundle as not loaded

## Error values

Sentinel errors:

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
