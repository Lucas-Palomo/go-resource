# Notas de arquitetura da v2

## Objetivo de arquitetura

A v2 foi desenhada para tratar o `go-resource` como biblioteca de verdade, não apenas como um helper de leitura de arquivos.

O alvo é simples:

- separar responsabilidades
- reduzir acoplamento
- eliminar `panic` do caminho principal
- preparar terreno para futuras extensões sem deformar a API

## Separação por responsabilidades

### `bundle.go`

Camada pública de estado, lookup, fallback e operações de runtime.

### `loader.go`

Camada de carregamento, walking de diretórios/`fs.FS`, parsing de nomes de arquivos e flatten de dados.

### `decoder.go`

Contrato de decoder e implementações padrão para JSON, YAML e TOML.

### `options.go`

Configuração declarativa da biblioteca.

### `errors.go`

Erros públicos e erros estruturados.

## Decisão: `LoadDir` e `LoadFS`

A v1 estava presa ao filesystem tradicional. A v2 passa a aceitar `fs.FS`, o que abre espaço para:

- `embed.FS`
- testes mais simples
- composição melhor com bibliotecas Go modernas

## Decisão: flatten para dot notation

Em i18n real, arquivos crescem. Mapas planos rapidamente ficam difíceis de manter. Por isso a v2 converte objetos aninhados para chaves planas previsíveis.

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

## Decisão: namespacing híbrido

A v2 aceita namespace por:

- estrutura de diretórios
- segmentos extras no nome do arquivo

Isso permite duas escolas de organização sem forçar apenas uma.

## Decisão: política explícita para falhas

Biblioteca madura não deve esconder escolhas importantes.

Por isso a v2 torna explícitos dois comportamentos:

- como tratar chave ausente
- como tratar chave duplicada

## Decisão: `Get` continua, mas `Lookup` é a API mais correta

`Get` existe por ergonomia e compatibilidade mental com a v1.

Mas `Lookup` é a API que melhor serve cenários sérios, porque devolve `error` e torna o fluxo explícito.

## Decisão: thread-safe para leitura

A estrutura usa `sync.RWMutex` para proteger estado interno, permitindo leituras concorrentes com segurança.

## Evoluções futuras já preparadas

A arquitetura da v2 já abre caminho para:

- pluralização
- placeholders nomeados
- fallback mais sofisticado por matcher
- carregamento incremental
- hot reload opcional
- métricas de missing keys
- integração com backends externos
