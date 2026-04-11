# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

Este changelog acompanha o **módulo raiz** (`github.com/Lucas-Palomo/go-resource`), que corresponde à linha de manutenção da v1.

## [v1.0.1] - 2026-04-11

### Alterado

- removido o `panic` do fluxo de carregamento da v1
- adicionado `LoadWithError()` para tratamento explícito de erro sem quebrar a API legada de `Load()`
- adicionado `Err()` para que chamadas que ainda usam `Load()` consigam inspecionar a última falha de carregamento
- corrigida a resolução de fallback em `Get()` para que a locale default seja realmente consultada quando a locale atual não contiver a chave
- inicializado `currentLocale` com a locale default em `NewBundle()` para um comportamento mais seguro logo na criação

### Documentação

- reescrito o README da raiz para deixar explícito o status de manutenção da v1
- adicionadas referências dedicadas da v1 em inglês e português
- alinhada a documentação da raiz com a estratégia multi-major do repositório

### Gestão de release

- retraída a `v1.0.0` no `go.mod` da raiz

## [v1.0.0] - 2024-08-20

### Release pública inicial

- primeira release pública da API v1
