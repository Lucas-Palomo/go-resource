# Referência da API v2

[English](./api-reference.md) | [Português (Brasil)](./api-reference.pt_br.md)

## Pacote

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Tipos públicos principais

- `Bundle`
- `Catalog`
- `Decoder`
- `DecoderFunc`
- `MissingKeyStrategy`
- `DuplicateKeyStrategy`

## Construção

### `func New(opts ...Option) *Bundle`

Defaults:

- locale de fallback: `language.English`
- locale atual: locale de fallback quando não for definida explicitamente
- estratégia para chave ausente: `ReturnKeyOnMissing`
- estratégia para chave duplicada: `ErrorOnDuplicate`

Decoders nativos:

- `.json`
- `.yaml`
- `.yml`
- `.toml`
- `.properties`

## Opções

- `WithFallbackLocale(locale language.Tag)`
- `WithLocale(locale language.Tag)`
- `WithMissingKeyStrategy(strategy MissingKeyStrategy)`
- `WithDuplicateKeyStrategy(strategy DuplicateKeyStrategy)`

## Carregamento

- `LoadDir(root string) error`
- `LoadFS(fsys fs.FS, root string) error`
- `RegisterDecoder(ext string, decoder Decoder)`

Regras de carregamento:

- o primeiro segmento do nome do arquivo é a locale
- nomes de diretório viram segmentos de namespace
- segmentos restantes do nome do arquivo viram segmentos de namespace
- objetos aninhados são achatados em notação por ponto

## Lookup

- `Lookup(key string, args ...any) (string, error)`
- `LookupFor(locale language.Tag, key string, args ...any) (string, error)`
- `Get(key string, args ...any) string`
- `GetFor(locale language.Tag, key string, args ...any) string`
- `Has(key string) bool`
- `HasFor(locale language.Tag, key string) bool`

## Helpers de runtime

- `SetLocale(locale language.Tag)`
- `Locale() language.Tag`
- `FallbackLocale() language.Tag`
- `Loaded() bool`
- `Locales() []language.Tag`
- `Catalog(locale language.Tag) Catalog`
- `Reset()`

## Valores de erro

- `ErrBundleNotLoaded`
- `ErrUnsupportedFormat`
- `ErrInvalidResourceFileName`
- `ErrDuplicateKey`
- `ErrMissingKey`
- `ErrInvalidResourceValue`

Erros estruturados:

- `KeyNotFoundError`
- `DuplicateKeyError`
- `ResourceFileNameError`
