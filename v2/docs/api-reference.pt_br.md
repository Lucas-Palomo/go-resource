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

## Estratégias

### Chave ausente

- `ReturnKeyOnMissing`: retorna a chave original
- `ReturnEmptyOnMissing`: retorna `""`
- `ErrorOnMissing`: retorna `ErrMissingKey` por meio de `Lookup` e `LookupFor`

### Chave duplicada

- `OverwriteOnDuplicate`: mantém o último valor carregado
- `ErrorOnDuplicate`: aborta o carregamento com `ErrDuplicateKey`

## Carregamento

- `LoadDir(root string) error`
- `LoadFS(fsys fs.FS, root string) error`
- `RegisterDecoder(ext string, decoder Decoder)`

Regras de carregamento:

- o primeiro segmento do nome do arquivo é a locale
- nomes de diretório viram segmentos de namespace
- segmentos restantes do nome do arquivo viram segmentos de namespace
- objetos aninhados são achatados em notação por ponto
- o catálogo em runtime é sempre `map[string]string`

Exemplos:

```text
resources/errors/en.json            -> prefixo de namespace: errors
resources/en.messages.checkout.toml -> prefixo de namespace: messages.checkout
```

## Lookup

- `Lookup(key string, args ...any) (string, error)`
- `LookupFor(locale language.Tag, key string, args ...any) (string, error)`
- `Get(key string, args ...any) string`
- `GetFor(locale language.Tag, key string, args ...any) string`
- `Has(key string) bool`
- `HasFor(locale language.Tag, key string) bool`

Notas de comportamento:

- `Lookup` e `LookupFor` são as APIs explícitas
- `Get` e `GetFor` são wrappers de conveniência
- quando argumentos de formatação são fornecidos, os valores são formatados com `fmt.Sprintf`
- a resolução usa a locale solicitada, seus pais, a locale de fallback, os pais do fallback e por fim `language.Und`

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
