# Estratégia de Versionamento e Release

[English](./versioning-strategy.md) | [Português (Brasil)](./versioning-strategy.pt_br.md)

Este repositório usa um **layout multi-major** para que a v1 pública original continue compatível enquanto a próxima major evolui de forma independente.

## Módulos

- módulo raiz: `github.com/Lucas-Palomo/go-resource`
- pacote raiz: `github.com/Lucas-Palomo/go-resource/pkg/resource`
- módulo v2: `github.com/Lucas-Palomo/go-resource/v2`

## Status atual

### Raiz / v1

- status: **linha estável de manutenção**
- release recomendada: **`v1.0.1+`**
- objetivo: preservar compatibilidade para consumidores existentes

### `/v2`

- status: **próxima major**
- objetivo: carregar a reescrita arquitetural e a evolução futura
- modelo de release: publicar separadamente como **`v2.0.0`** no path de módulo `/v2`

## Por que esse layout existe

O projeto já havia sido publicado publicamente como módulo v1. Como o Go exige semantic import versioning para novas majors, o path correto para a próxima linha com breaking changes é `/v2`.

Isso mantém consumidores existentes no import path da raiz e permite que a nova linha evolua sem quebrar usuários da v1.

## Por que a `v1.0.1` existe

A release pública inicial expunha falhas de carregamento por meio de `panic`, o que não é um contrato sólido para uma biblioteca reutilizável.

A `v1.0.1` existe para estabilizar a linha raiz sem forçar uma reescrita disruptiva:

- manter o import path existente
- parar de derrubar o processo em falhas normais de carregamento
- corrigir o comportamento de fallback
- documentar com clareza o contrato suportado da v1

## Política de retract

O `go.mod` da raiz retrai `v1.0.0` para que o tooling do Go sinalize que a primeira release pública não deve ser considerada a versão recomendada.

## Consequências práticas

- usuários existentes da v1 continuam importando `github.com/Lucas-Palomo/go-resource/pkg/resource`
- a linha da raiz deve receber apenas mudanças de baixo risco e foco em compatibilidade
- a evolução arquitetural deve acontecer em `/v2`
- quando `v2.0.0` for publicada, projetos novos devem preferir `github.com/Lucas-Palomo/go-resource/v2`
