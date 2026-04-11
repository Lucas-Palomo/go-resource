# Referência da v1

[English](./v1-reference.md) | [Português (Brasil)](./v1-reference.pt_br.md)

## Pacote

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

## Propósito

A linha v1 é a API com foco em compatibilidade para consumidores que já usam o pacote original. Na `v1.0.1`, a principal melhoria não é uma reescrita. É uma correção de contrato: falhas de carregamento não derrubam mais o processo por padrão.

## Tipo principal

### `type Bundle`

Estrutura central da biblioteca. Ela armazena:

- a raiz de recursos configurada
- os catálogos por locale carregados em memória
- a locale atual
- a locale fallback default
- o último erro de carregamento capturado pelas APIs de compatibilidade

## Construção

### `func NewBundle(resourcesFolder string, defaultLocale language.Tag) *Bundle`

Cria um novo bundle para a raiz de recursos informada e para a locale default.

Comportamento na `v1.0.1`:

- inicializa `currentLocale` com a locale default
- começa com catálogo em memória vazio
- preserva compatibilidade com o código de construção já existente

## Carregamento

### `func (b *Bundle) Load()`

Método legado de carregamento compatível com a API anterior.

Na `v1.0.1`, esse método não usa mais `panic`. Em vez disso, ele registra internamente o último erro de carregamento. Chame `Err()` logo após `Load()` quando precisar inspecionar a falha.

```go
bundle.Load()
if err := bundle.Err(); err != nil {
	return err
}
```

### `func (b *Bundle) LoadWithError() error`

Método preferível para código novo.

Ele retorna diretamente erros de filesystem, parsing de locale, decoder ou formato não suportado:

```go
if err := bundle.LoadWithError(); err != nil {
	return err
}
```

### `func (b *Bundle) Err() error`

Retorna o último erro de carregamento capturado por `Load()` ou `LoadWithError()`.

## Lookup

### `func (b *Bundle) Get(id string, replacers ...any) string`

Resolve uma chave usando a locale atual.

Ordem de resolução na `v1.0.1`:

1. locale atual
2. locale default
3. retorna a própria chave se ainda não houver resolução

### `func (b *Bundle) GetWithLocale(locale language.Tag, id string, replacers ...any) string`

Resolve a chave para uma locale específica sem alterar o estado do bundle.

Se a chave não existir, a própria chave é retornada.

## Configuração em runtime

### `func (b *Bundle) SetLocale(locale language.Tag)`

Troca a locale atual usada por `Get()`.

## Formatos suportados

- `.json`
- `.yaml`
- `.yml`
- `.toml`

## Regra de nome de arquivo

A v1 espera que cada arquivo de recurso use a locale no primeiro segmento do nome do arquivo.

Exemplos:

- `en.json`
- `pt_BR.yaml`
- `es.toml`

Diretórios aninhados são permitidos, mas a locale continua sendo extraída do nome do arquivo.

## Modelo de erro na `v1.0.1`

Casos típicos de falha de carregamento:

- diretório ou arquivo ilegível
- locale inválida no nome do arquivo
- extensão não suportada
- payload JSON / YAML / TOML inválido

O ponto importante é que essas falhas agora voltam como `error`, em vez de encerrar a aplicação por meio de `panic`.

## Recomendação

Continue na v1 quando compatibilidade for mais importante que mudança arquitetural.

Para sistemas novos ou evolução mais profunda, migre para `/v2`.
