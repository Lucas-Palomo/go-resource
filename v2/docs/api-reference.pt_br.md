# Referência da API v2

[English](./api-reference.md) | [Português (Brasil)](./api-reference.pt_br.md)

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

Controla o comportamento quando uma chave não é encontrada.

- `ReturnKeyOnMissing`
- `ReturnEmptyOnMissing`
- `ErrorOnMissing`

### `type DuplicateKeyStrategy`

Controla o comportamento quando uma chave é declarada mais de uma vez.

- `OverwriteOnDuplicate`
- `ErrorOnDuplicate`

## Construção e opções

### `func New(opts ...Option) *Bundle`

Cria um bundle com decoders padrão para JSON, YAML, YML, TOML e arquivos `.properties` no estilo Java.

Exemplo:

```go
bundle := resource.New(
	resource.WithFallbackLocale(language.English),
	resource.WithLocale(language.MustParse("pt-BR")),
)
```

### `func WithFallbackLocale(locale language.Tag) Option`

Define a locale fallback do bundle.

### `func WithLocale(locale language.Tag) Option`

Define a locale inicial usada por `Get` e `Lookup`.

### `func WithMissingKeyStrategy(strategy MissingKeyStrategy) Option`

Define a política para chave ausente.

### `func WithDuplicateKeyStrategy(strategy DuplicateKeyStrategy) Option`

Define a política para chave duplicada.

## Carregamento

### `func (b *Bundle) LoadDir(root string) error`

Carrega recursos a partir de um diretório do sistema operacional.

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

Atalho ergonômico para `Lookup`. Útil para casos simples e continuidade mental com a v1.

### `func (b *Bundle) GetFor(locale language.Tag, key string, args ...any) string`

Atalho ergonômico para `LookupFor`.

### `func (b *Bundle) Has(key string) bool`

Informa se uma chave pode ser resolvida usando a locale atual e a cadeia de fallback.

### `func (b *Bundle) HasFor(locale language.Tag, key string) bool`

Informa se uma chave pode ser resolvida para uma locale específica e sua cadeia de fallback.

## Configuração em runtime

### `func (b *Bundle) SetLocale(locale language.Tag)`

Troca a locale atual do bundle.

### `func (b *Bundle) RegisterDecoder(ext string, decoder Decoder)`

Registra um novo decoder para uma extensão.

```go
bundle.RegisterDecoder(".ini", myDecoder)
```

### `func (b *Bundle) Reset()`

Limpa os catálogos carregados preservando a configuração em runtime e os decoders registrados.

## Inspeção

### `func (b *Bundle) Locale() language.Tag`

Retorna a locale atual.

### `func (b *Bundle) FallbackLocale() language.Tag`

Retorna a locale fallback configurada.

### `func (b *Bundle) Locales() []language.Tag`

Lista as locales carregadas.

### `func (b *Bundle) Catalog(locale language.Tag) Catalog`

Retorna uma cópia defensiva do catálogo achatado daquela locale.

### `func (b *Bundle) Loaded() bool`

Informa se ao menos uma operação de carregamento foi concluída com sucesso.

## Decoders built-in

- `type JSONDecoder`
- `type YAMLDecoder`
- `type TOMLDecoder`
- `type PropertiesDecoder`

## Erros relevantes

- `ErrBundleNotLoaded`
- `ErrUnsupportedFormat`
- `ErrInvalidResourceFileName`
- `ErrDuplicateKey`
- `ErrMissingKey`
- `ErrInvalidResourceValue`

## Erros estruturados

### `type KeyNotFoundError`

Retornado quando `ErrorOnMissing` está habilitado e a chave não pode ser resolvida.

### `type DuplicateKeyError`

Retornado quando `ErrorOnDuplicate` está habilitado e uma chave é declarada mais de uma vez.

### `type ResourceFileNameError`

Retornado quando o nome do arquivo não segue o contrato esperado.
