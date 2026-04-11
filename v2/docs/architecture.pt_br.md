# Notas de arquitetura da v2

[English](./architecture.md) | [Português (Brasil)](./architecture.pt_br.md)

## Objetivo arquitetural

A v2 trata o `go-resource` como uma biblioteca reutilizável de verdade, e não apenas como um helper de leitura de arquivos.

Os objetivos de design são diretos:

- separar responsabilidades
- reduzir acoplamento
- remover `panic` do caminho principal de erro
- deixar espaço para evolução futura sem deformar a API pública

## Separação de responsabilidades

### `bundle.go`

Estado público em runtime, estado de locale, comportamento de lookup, resolução de fallback e helpers de conveniência.

### `loader.go`

Walking de filesystem, suporte a `fs.FS`, descoberta de arquivos de recurso, parsing de nomes de arquivo e flatten de estruturas aninhadas.

### `decoder.go`

Contrato de decoder e implementações nativas para JSON, YAML, TOML e arquivos `.properties` no estilo Java.

### `options.go`

Configuração declarativa do comportamento do bundle.

### `errors.go`

Erros públicos e tipos de erro estruturados.

## Por que `LoadDir` e `LoadFS` coexistem

A v1 era presa ao filesystem do sistema operacional. Isso era estreito demais.

A v2 suporta explicitamente `fs.FS`, o que abre espaço para:

- `embed.FS`
- testes mais limpos
- composição mais fácil com bibliotecas Go modernas

## Por que objetos aninhados são achatados

Catálogos reais de i18n crescem. Estruturas profundas são mais fáceis de organizar do que mapas gigantes totalmente planos, mas o lookup em runtime continua se beneficiando de chaves planas previsíveis.

Exemplo:

```json
{
  "checkout": {
    "button": {
      "confirm": "Confirm"
    }
  }
}
```

vira:

```text
checkout.button.confirm
```

## Por que o namespacing é híbrido

A v2 aceita informação de namespace a partir de:

- estrutura de diretórios
- segmentos adicionais no nome do arquivo

Isso suporta os dois estilos mais comuns de organização sem impor apenas um deles.

Exemplo:

```text
resources/errors/en.json
```

com:

```json
{
  "validation": {
    "required": "Required"
  }
}
```

vira:

```text
errors.validation.required
```

## Por que as políticas de falha são explícitas

Uma biblioteca reutilizável não deve escolher silenciosamente comportamentos críticos que o usuário talvez queira controlar.

A v2 torna explícitas duas políticas:

- como tratar chaves ausentes
- como tratar chaves duplicadas

## Por que `Get` ainda existe

`Get` continua por ergonomia e continuidade com a v1.

`Lookup` é a API mais correta para cenários menos triviais porque retorna `error` e deixa o tratamento de falhas explícito.

## Modelo de concorrência

O bundle usa `sync.RWMutex` para proteger o estado interno. A intenção é permitir leituras concorrentes seguras com mutação controlada em runtime.

## Helpers extras de runtime

Os seguintes helpers foram adicionados para reduzir atrito em testes, diagnóstico e fluxos controlados de recarga:

- `Has`
- `HasFor`
- `Reset`

## Evolução futura já preparada

A estrutura atual já deixa espaço para futuras adições, como:

- pluralização
- placeholders nomeados
- matching de locale mais sofisticado
- carregamento incremental
- hot reload opcional
- métricas de missing key
- integração com backends externos
