# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

Este changelog acompanha o **módulo raiz** (`github.com/Lucas-Palomo/go-resource`), que corresponde à linha v1.

## [v1.0.1] - 2026-04-11

### Alterado

- removido o `panic` do fluxo padrão de carregamento da v1
- adicionado `LoadWithError()` para tratamento explícito de erro sem quebrar a API legada de `Load()`
- adicionado `Err()` para que chamadas que ainda usam `Load()` possam inspecionar a última falha de carregamento
- corrigida a resolução de fallback em `Get()` para que a locale default seja consultada quando a locale atual não contiver a chave
- inicializado `currentLocale` com a locale default em `NewBundle()` para um comportamento padrão mais seguro

### Documentação

- reescrito o README da raiz com base no layout real do repositório
- documentada a diferença entre o modelo de recursos da v1 e o da v2
- adicionados documentos dedicados de referência da v1 e versionamento em inglês e pt-BR

### Gestão de release

- retraída a `v1.0.0` no `go.mod` da raiz

## [v1.0.0] - 2024-08-20

### Adicionado

- primeira release pública da API original da v1
