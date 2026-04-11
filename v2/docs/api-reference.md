# v2 API Reference

[English](./api-reference.md) | [Português (Brasil)](./api-reference.pt_br.md)

## Package

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Main types

### `type Bundle`

Core structure of the library. It stores loaded catalogs by locale, registered decoders, current locale state, fallback configuration, and lookup strategies.

### `type Decoder`

Contract used to support new file formats.

```go
type Decoder interface {
	Decode(data []byte) (map[string]any, error)
}
```

### `type DecoderFunc`

Adapter that allows a decoder to be registered from a function.

### `type MissingKeyStrategy`

Controls what happens when a key cannot be resolved.

Available values:

- `ReturnKeyOnMissing`
- `ReturnEmptyOnMissing`
- `ErrorOnMissing`

### `type DuplicateKeyStrategy`

Controls what happens when the same logical key is declared more than once.

Available values:

- `OverwriteOnDuplicate`
- `ErrorOnDuplicate`

## Construction

### `func New(opts ...Option) *Bundle`

Creates a new bundle with built-in decoders for:

- JSON
- YAML / YML
- TOML
- Java-style `.properties`

Typical initialization:

```go
bundle := resource.New(
	resource.WithFallbackLocale(language.English),
	resource.WithLocale(language.MustParse("pt-BR")),
)
```

## Options

### `func WithFallbackLocale(locale language.Tag) Option`

Sets the fallback locale consulted by lookup operations.

### `func WithLocale(locale language.Tag) Option`

Sets the initial locale used by `Get`, `Lookup`, `Has`, and related helpers.

### `func WithMissingKeyStrategy(strategy MissingKeyStrategy) Option`

Defines the missing-key policy.

### `func WithDuplicateKeyStrategy(strategy DuplicateKeyStrategy) Option`

Defines the duplicate-key policy.

## Loading

### `func (b *Bundle) LoadDir(root string) error`

Loads resources from a directory in the operating system filesystem.

```go
if err := bundle.LoadDir("./resources"); err != nil {
	return err
}
```

### `func (b *Bundle) LoadFS(fsys fs.FS, root string) error`

Loads resources from any `fs.FS`, including `embed.FS`.

## Lookup

### `func (b *Bundle) Lookup(key string, args ...any) (string, error)`

Preferred API for serious usage. Resolves a key using the current locale and the fallback chain, returning an explicit `error` when configured to do so.

### `func (b *Bundle) LookupFor(locale language.Tag, key string, args ...any) (string, error)`

Resolves a key for a specific locale without mutating bundle state.

### `func (b *Bundle) Get(key string, args ...any) string`

Ergonomic shortcut for `Lookup`. It exists for simpler usage and continuity with v1.

### `func (b *Bundle) GetFor(locale language.Tag, key string, args ...any) string`

Ergonomic shortcut for `LookupFor`.

### `func (b *Bundle) Has(key string) bool`

Reports whether a key can be resolved through the current locale and fallback chain.

### `func (b *Bundle) HasFor(locale language.Tag, key string) bool`

Reports whether a key can be resolved for a specific locale and its fallback chain.

## Runtime configuration

### `func (b *Bundle) SetLocale(locale language.Tag)`

Changes the current locale of the bundle.

### `func (b *Bundle) RegisterDecoder(ext string, decoder Decoder)`

Registers a new decoder for a file extension.

```go
bundle.RegisterDecoder(".ini", myDecoder)
```

### `func (b *Bundle) Reset()`

Clears loaded catalogs while preserving runtime configuration and registered decoders.

## Inspection helpers

### `func (b *Bundle) Locale() language.Tag`

Returns the current locale.

### `func (b *Bundle) FallbackLocale() language.Tag`

Returns the configured fallback locale.

### `func (b *Bundle) Locales() []language.Tag`

Lists loaded locales.

### `func (b *Bundle) Catalog(locale language.Tag) Catalog`

Returns a defensive copy of the flattened catalog for the given locale.

### `func (b *Bundle) Loaded() bool`

Reports whether at least one load operation has completed successfully.

## Built-in decoders

- `type JSONDecoder`
- `type YAMLDecoder`
- `type TOMLDecoder`
- `type PropertiesDecoder`

## Relevant errors

- `ErrBundleNotLoaded`
- `ErrUnsupportedFormat`
- `ErrInvalidResourceFileName`
- `ErrDuplicateKey`
- `ErrMissingKey`
- `ErrInvalidResourceValue`

## Structured errors

### `type KeyNotFoundError`

Returned when `ErrorOnMissing` is enabled and the key cannot be resolved.

### `type DuplicateKeyError`

Returned when `ErrorOnDuplicate` is enabled and a key is declared more than once.

### `type ResourceFileNameError`

Returned when a resource file name does not follow the expected contract.
