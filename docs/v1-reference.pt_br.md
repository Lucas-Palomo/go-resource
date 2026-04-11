# Referência da v1

[English](./v1-reference.md) | [Português (Brasil)](./v1-reference.pt_br.md)

## Pacote

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

## Tipo principal

### `type Bundle`

Estrutura central da biblioteca. Armazena a raiz configurada de recursos, os catálogos por locale e o estado da locale atual/default.

## Construção

### `func NewBundle(resourcesFolder string, defaultLocale language.Tag) *Bundle`

Cria um novo bundle para o diretório informado e para a locale default.

Comportamento na `v1.0.1`:
- inicializa `currentLocale` com a locale default
- começa com catálogo em memória vazio
- mantém compatibilidade com o código de construção já existente

## Carregamento

### `func (b *Bundle) Load()`

Método legado de carregamento compatível com a API anterior.

Na `v1.0.1`, esse método não usa mais `panic`. Em vez disso, ele registra internamente o último erro de carregamento.
Use `Err()` logo após a chamada quando precisar inspecionar a falha.

### `func (b *Bundle) LoadWithError() error`

Método preferível para código novo. Retorna diretamente qualquer erro de filesystem, parsing de locale, decoder ou formato não suportado.

### `func (b *Bundle) Err() error`

Retorna o último erro de carregamento capturado por `Load()` ou `LoadWithError()`.

## Lookup

### `func (b *Bundle) Get(id string, replacers ...any) string`

Resolve a chave usando a locale atual. Se a locale atual não contiver a chave, a locale default é consultada. Se a chave ainda não existir, a própria chave é retornada.

### `func (b *Bundle) GetWithLocale(locale language.Tag, id string, replacers ...any) string`

Resolve a chave para uma locale específica sem alterar o estado do bundle. Se a chave não existir, a própria chave é retornada.

## Configuração em runtime

### `func (b *Bundle) SetLocale(locale language.Tag)`

Troca a locale atual usada por `Get()`.

## Formatos suportados

- `.json`
- `.yaml`
- `.yml`
- `.toml`

## Regra de nome de arquivo

A v1 espera que cada arquivo de recurso tenha o formato:

```text
<locale>.<extensão>
```

Exemplos:
- `en.json`
- `pt_BR.yaml`
- `es.toml`

O primeiro segmento do nome do arquivo é interpretado como locale.

## Modelo de erro na `v1.0.1`

Erros de carregamento agora voltam como `error` em vez de derrubar o processo com `panic`.
Casos típicos de falha:
- diretório ou arquivo ilegível
- locale inválida no nome do arquivo
- extensão não suportada
- payload JSON / YAML / TOML inválido

## Nota de compatibilidade

`Load()` foi preservado para evitar uma quebra de API numa patch release.
Para qualquer código novo, `LoadWithError()` é a entrada recomendada.
