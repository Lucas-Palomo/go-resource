# Referência da API v2

## Package

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Tipos principais

### `type Bundle`

Estrutura central da biblioteca. Armazena catálogos por locale, decoders registrados e estratégias de lookup.

### `type Decoder`

Contrato para suportar novos formatos de arquivo.

```go
type Decoder interface {
    Decode(data []byte) (map[string]any, error)
}
```

### `type DecoderFunc`

Adapter para registrar decoder a partir de função.

### `type MissingKeyStrategy`

Controla o comportamento quando uma key não é encontrada.

- `ReturnKeyOnMissing`
- `ReturnEmptyOnMissing`
- `ErrorOnMissing`

### `type DuplicateKeyStrategy`

Controla o comportamento quando uma key é declarada mais de uma vez.

- `OverwriteOnDuplicate`
- `ErrorOnDuplicate`

## Construção e opções

### `func New(opts ...Option) *Bundle`

Cria um bundle com decoders padrão para JSON, YAML, YML e TOML.

Exemplo:

```go
bundle := resource.New(
    resource.WithFallbackLocale(language.English),
    resource.WithLocale(language.MustParse("pt-BR")),
)
```

### `func WithFallbackLocale(locale language.Tag) Option`

Define a locale default do bundle.

### `func WithLocale(locale language.Tag) Option`

Define a locale inicial usada por `Get` e `Lookup`.

### `func WithMissingKeyStrategy(strategy MissingKeyStrategy) Option`

Define a política para chave ausente.

### `func WithDuplicateKeyStrategy(strategy DuplicateKeyStrategy) Option`

Define a política para chave duplicada.

## Carregamento

### `func (b *Bundle) LoadDir(root string) error`

Carrega recursos a partir de um diretório do sistema operacional.

Exemplo:

```go
if err := bundle.LoadDir("./resources"); err != nil {
    return err
}
```

### `func (b *Bundle) LoadFS(fsys fs.FS, root string) error`

Carrega recursos a partir de qualquer `fs.FS`, incluindo `embed.FS`.

## Lookup

### `func (b *Bundle) Lookup(key string, args ...any) (string, error)`

Resolve uma chave usando a locale atual e a cadeia de fallback.

### `func (b *Bundle) LookupFor(locale language.Tag, key string, args ...any) (string, error)`

Resolve uma chave forçando uma locale específica, sem alterar o estado atual do bundle.

### `func (b *Bundle) Get(key string, args ...any) string`

Atalho ergonômico para `Lookup`. Útil para casos simples e compatibilidade de estilo com a v1.

### `func (b *Bundle) GetFor(locale language.Tag, key string, args ...any) string`

Atalho ergonômico para `LookupFor`.

## Configuração em runtime

### `func (b *Bundle) SetLocale(locale language.Tag)`

Troca a locale atual do bundle.

### `func (b *Bundle) RegisterDecoder(ext string, decoder Decoder)`

Registra um novo decoder para uma extensão. Exemplo:

```go
bundle.RegisterDecoder(".ini", myDecoder)
```

## Inspeção

### `func (b *Bundle) Locale() language.Tag`

Retorna a locale atual.

### `func (b *Bundle) FallbackLocale() language.Tag`

Retorna a locale default configurada.

### `func (b *Bundle) Locales() []language.Tag`

Lista as locales carregadas.

### `func (b *Bundle) Catalog(locale language.Tag) Catalog`

Retorna uma cópia defensiva do catálogo achatado daquela locale.

### `func (b *Bundle) Loaded() bool`

Informa se ao menos um carregamento foi concluído com sucesso.

## Decoders built-in

### `type JSONDecoder`

Decoder padrão para `.json`.

### `type YAMLDecoder`

Decoder padrão para `.yaml` e `.yml`.

### `type TOMLDecoder`

Decoder padrão para `.toml`.

## Erros relevantes

- `ErrBundleNotLoaded`
- `ErrUnsupportedFormat`
- `ErrInvalidResourceFileName`
- `ErrDuplicateKey`
- `ErrMissingKey`
- `ErrInvalidResourceValue`

## Erros estruturados

### `type KeyNotFoundError`

Erro retornado quando `ErrorOnMissing` está habilitado e a chave não é encontrada.

### `type DuplicateKeyError`

Erro retornado quando `ErrorOnDuplicate` está habilitado e uma key se repete.

### `type ResourceFileNameError`

Erro retornado quando o nome do arquivo não segue o contrato esperado.
