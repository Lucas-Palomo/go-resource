# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

## [v1.0.1] - 2026-04-11

### Alterado
- removido o `panic` do fluxo de carregamento da v1
- adicionado `LoadWithError()` para tratamento explícito de erro sem quebrar o call site existente de `Load()`
- adicionado `Err()` para inspecionar a última falha de carregamento ao usar o `Load()` legado
- corrigida a resolução de fallback em `Get()` para que a locale default seja realmente consultada quando a locale atual não contiver a chave
- inicializado `currentLocale` com a locale default em `NewBundle()` para um comportamento mais seguro logo de saída

### Documentação
- reescrito o README da raiz para a linha de manutenção da v1
- adicionadas referências dedicadas da v1 em inglês e português
- adicionada documentação do pacote via `pkg/resource/doc.go`

### Gestão de release
- retraída a `v1.0.0` no `go.mod` da raiz

## [v1.0.0] - 2024-08-20

### Release pública inicial
- primeira release pública da API v1
