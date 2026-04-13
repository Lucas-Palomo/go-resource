# Estratégia de Versionamento e Release

[English](./versioning-strategy.md) | [Português (Brasil)](./versioning-strategy.pt_br.md)

Este repositório segue um **layout multi-major**: a v1 original permanece na raiz do repositório, enquanto a reescrita com breaking changes vive em `/v2`.

## Tags publicadas

| Tag | Módulo | Import path | Data | Significado |
|---|---|---|---|---|
| `v1.0.0` | v1 | `github.com/Lucas-Palomo/go-resource` | 2024-08-20 | Primeira release pública, hoje retraída |
| `v1.0.1` | v1 | `github.com/Lucas-Palomo/go-resource` | 2026-04-11 | Release estável de manutenção para usuários existentes |
| `v2.0.0` | v2 | `github.com/Lucas-Palomo/go-resource/v2` | 2026-04-11 | Primeira tag estável da v2, depois retraída porque a documentação publicada estava semanticamente incorreta |
| `v2.0.1` | v2 | `github.com/Lucas-Palomo/go-resource/v2` | 2026-04-12 | Release estável de correção documental da v2 |

## Paths de módulo

- módulo raiz: `github.com/Lucas-Palomo/go-resource`
- path do pacote v1: `github.com/Lucas-Palomo/go-resource/pkg/resource`
- módulo v2: `github.com/Lucas-Palomo/go-resource/v2`

## Por que o repositório é estruturado assim

Go exige **semantic import versioning** para majors com breaking changes.

Isso significa:

- a linha pública original mantém o import path da raiz
- a reescrita com breaking changes precisa viver em `/v2`
- as duas linhas podem coexistir sem quebrar consumidores existentes

## Política atual de suporte

### Raiz / v1

- propósito: linha de manutenção com foco em compatibilidade
- versão recomendada: `v1.0.1`
- perfil esperado de mudança: correções de baixo risco, documentação e estabilidade

### `/v2`

- propósito: linha major ativa
- versão recomendada para projetos novos: `v2.0.1`
- perfil esperado de mudança: evolução principal do produto sob o path de módulo v2

## Por que `v1.0.0` foi retraída

A primeira release pública da v1 expunha falhas normais de carregamento por meio de `panic`.

A `v1.0.1` corrigiu esse contrato sem forçar uma reescrita com breaking changes:

- `Load()` não entra mais em `panic` por padrão
- `LoadWithError()` retorna erros explícitos
- `Err()` preserva compatibilidade para estilos de chamada antigos
- o comportamento de fallback em `Get()` foi corrigido

O `go.mod` da raiz retracta `v1.0.0` para que o tooling do Go sinalize que ela não é a versão recomendada.

## Por que `v2.0.0` foi retraída

A `v2.0.0` era funcionalmente válida, mas a documentação publicada descrevia o estado do módulo e a semântica de uso de forma incorreta o bastante para justificar uma release de correção.

A `v2.0.1` existe para se tornar o ponto de entrada estável da v2 com documentação e metadados de release corrigidos.

## Regras práticas

- continue na v1 quando preservar o import path existente for o mais importante
- escolha a v2 para código novo ou para times que precisam de tratamento explícito de erro e um modelo de recursos mais rico
- não publique breaking changes no path do módulo raiz
- publique futuras mudanças breaking sob o path versionado que o Go espera

## Regra de bolso para releases

- **releases v1**: manutenção e compatibilidade
- **releases v2**: evolução principal do produto
