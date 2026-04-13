# Referência da API v2

[English](./api-reference.md) | [Português (Brasil)](./api-reference.pt_br.md)

## Pacote

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Principais tipos públicos

- `Bundle`
- `Catalog`
- `Decoder`
- `DecoderFunc`
- `MissingKeyStrategy`
- `DuplicateKeyStrategy`

## Construção

### `func New(opts ...Option) *Bundle`

Cria um novo bundle com os decoders nativos e as estratégias padrão de runtime.

Padrões:

- fallback locale: `language.English`
- locale atual: fallback locale quando não definido explicitamente
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

### Estratégias para chave ausente

- `ReturnKeyOnMissing`: retorna a chave original
- `ReturnEmptyOnMissing`: retorna `""`
- `ErrorOnMissing`: retorna `ErrMissingKey` via `Lookup` e `LookupFor`

### Estratégias para chave duplicada

- `OverwriteOnDuplicate`: mantém o valor carregado por último
- `ErrorOnDuplicate`: aborta o carregamento com `ErrDuplicateKey`

## API de carregamento

- `LoadDir(root string) error`
- `LoadFS(fsys fs.FS, root string) error`
- `RegisterDecoder(ext string, decoder Decoder)`

### Regras de carregamento

- o primeiro segmento do nome do arquivo é o locale
- nomes de diretório viram segmentos de namespace
- os demais segmentos do nome do arquivo viram segmentos de namespace
- objetos aninhados são achatados em notação por ponto
- o catálogo de runtime é sempre `map[string]string`
- carregamentos repetidos bem-sucedidos fazem merge nos catálogos já em memória
- use `Reset()` quando quiser substituição em vez de merge
- arquivos com extensão não suportada abortam o carregamento com `ErrUnsupportedFormat`
- `RegisterDecoder` normaliza extensões como `json` para `.json`
- `RegisterDecoder` ignora extensões vazias e decoders nil

Exemplos:

```text
resources/errors/en.json            -> prefixo de namespace: errors
resources/en.messages.checkout.toml -> prefixo de namespace: messages.checkout
```

## API de lookup

- `Lookup(key string, args ...any) (string, error)`
- `LookupFor(locale language.Tag, key string, args ...any) (string, error)`
- `Get(key string, args ...any) string`
- `GetFor(locale language.Tag, key string, args ...any) string`
- `Has(key string) bool`
- `HasFor(locale language.Tag, key string) bool`

### Notas de comportamento

- `Lookup` e `LookupFor` são as APIs estritas
- `Get` e `GetFor` são helpers lenientes
- `Get` e `GetFor` retornam a chave original em erros de lookup, inclusive em chave ausente e `ErrBundleNotLoaded`
- quando argumentos de formatação são fornecidos, os valores são formatados com `fmt.Sprintf`
- as buscas percorrem o locale solicitado, seus pais, o locale de fallback, os pais do fallback e por fim `language.Und`

## Helpers de runtime

- `SetLocale(locale language.Tag)`
- `Locale() language.Tag`
- `FallbackLocale() language.Tag`
- `Loaded() bool`
- `Locales() []language.Tag`
- `Catalog(locale language.Tag) Catalog`
- `Reset()`

### Semântica dos helpers

- `SetLocale` altera o locale padrão do bundle usado por `Lookup` e `Get`
- `Catalog` retorna uma cópia do catálogo do locale, não o mapa interno
- `Locales` retorna os locales atualmente carregados em memória
- `Reset` limpa os catálogos carregados e marca o bundle como não carregado

## Valores de erro

Erros sentinela:

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
