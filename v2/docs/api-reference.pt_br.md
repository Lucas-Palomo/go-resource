# Referência da API v2

[English](./api-reference.md) | [Português (Brasil)](./api-reference.pt_br.md)

## Pacote

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Tipos principais

### `type Bundle`

Estrutura central da biblioteca. Ela armazena catálogos carregados por locale, decoders registrados, estado da locale atual, configuração de fallback e estratégias de lookup.

### `type Decoder`

Contrato usado para suportar novos formatos de arquivo.

```go
type Decoder interface {
	Decode(data []byte) (map[string]any, error)
}
```

### `type DecoderFunc`

Adapter que permite registrar um decoder a partir de função.

### `type MissingKeyStrategy`

Controla o que acontece quando uma chave não pode ser resolvida.

Valores disponíveis:

- `ReturnKeyOnMissing`
- `ReturnEmptyOnMissing`
- `ErrorOnMissing`

### `type DuplicateKeyStrategy`

Controla o que acontece quando a mesma chave lógica é declarada mais de uma vez.

Valores disponíveis:

- `OverwriteOnDuplicate`
- `ErrorOnDuplicate`

## Construção

### `func New(opts ...Option) *Bundle`

Cria um novo bundle com decoders nativos para:

- JSON
- YAML / YML
- TOML
- arquivos `.properties` no estilo Java

Inicialização típica:

```go
bundle := resource.New(
	resource.WithFallbackLocale(language.English),
	resource.WithLocale(language.MustParse("pt-BR")),
)
```

## Opções

### `func WithFallbackLocale(locale language.Tag) Option`

Define a locale fallback consultada pelas operações de lookup.

### `func WithLocale(locale language.Tag) Option`

Define a locale inicial usada por `Get`, `Lookup`, `Has` e helpers relacionados.

### `func WithMissingKeyStrategy(strategy MissingKeyStrategy) Option`

Define a política para chave ausente.

### `func WithDuplicateKeyStrategy(strategy DuplicateKeyStrategy) Option`

Define a política para chave duplicada.

## Carregamento

### `func (b *Bundle) LoadDir(root string) error`

Carrega recursos de um diretório no filesystem do sistema operacional.

```go
if err := bundle.LoadDir("./resources"); err != nil {
	return err
}
```

### `func (b *Bundle) LoadFS(fsys fs.FS, root string) error`

Carrega recursos a partir de qualquer `fs.FS`, incluindo `embed.FS`.

## Lookup

### `func (b *Bundle) Lookup(key string, args ...any) (string, error)`

API preferível para uso sério. Resolve uma chave usando a locale atual e a cadeia de fallback, retornando `error` explícito quando configurado para isso.

### `func (b *Bundle) LookupFor(locale language.Tag, key string, args ...any) (string, error)`

Resolve uma chave para uma locale específica sem alterar o estado do bundle.

### `func (b *Bundle) Get(key string, args ...any) string`

Atalho ergonômico para `Lookup`. Ele existe para usos mais simples e para continuidade com a v1.

### `func (b *Bundle) GetFor(locale language.Tag, key string, args ...any) string`

Atalho ergonômico para `LookupFor`.

### `func (b *Bundle) Has(key string) bool`

Informa se uma chave pode ser resolvida pela locale atual e pela cadeia de fallback.

### `func (b *Bundle) HasFor(locale language.Tag, key string) bool`

Informa se uma chave pode ser resolvida para uma locale específica e sua cadeia de fallback.

## Configuração em runtime

### `func (b *Bundle) SetLocale(locale language.Tag)`

Troca a locale atual do bundle.

### `func (b *Bundle) RegisterDecoder(ext string, decoder Decoder)`

Registra um novo decoder para uma extensão de arquivo.

```go
bundle.RegisterDecoder(".ini", myDecoder)
```

### `func (b *Bundle) Reset()`

Limpa os catálogos carregados preservando a configuração em runtime e os decoders registrados.

## Helpers de inspeção

### `func (b *Bundle) Locale() language.Tag`

Retorna a locale atual.

### `func (b *Bundle) FallbackLocale() language.Tag`

Retorna a locale fallback configurada.

### `func (b *Bundle) Locales() []language.Tag`

Lista as locales carregadas.

### `func (b *Bundle) Catalog(locale language.Tag) Catalog`

Retorna uma cópia defensiva do catálogo achatado para a locale informada.

### `func (b *Bundle) Loaded() bool`

Informa se ao menos uma operação de carregamento terminou com sucesso.

## Decoders nativos

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

Retornado quando o nome de um arquivo de recurso não segue o contrato esperado.
