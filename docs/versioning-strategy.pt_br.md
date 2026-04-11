# Estratégia de Versionamento e Release

[English](./versioning-strategy.md) | [Português (Brasil)](./versioning-strategy.pt_br.md)

Este repositório segue a estratégia padrão de major version do Go.

## Linha v1

- módulo: `github.com/Lucas-Palomo/go-resource`
- pacote: `github.com/Lucas-Palomo/go-resource/pkg/resource`
- status: manutenção / compatibilidade

## Por que `v1.0.1`

A primeira release pública expunha falhas de carregamento por meio de `panic`.
Esse não é um bom contrato padrão para uma biblioteca reutilizável.

A `v1.0.1` existe para corrigir a linha raiz sem forçar uma reescrita disruptiva.

## Política de retract

O `go.mod` da raiz retrai `v1.0.0` para sinalizar que a release inicial não deve ser usada como versão recomendada.

## Fluxo sugerido de release

1. publicar a correção da linha raiz como `v1.0.1`
2. manter mudanças futuras na raiz pequenas e focadas em compatibilidade
3. publicar a nova major separadamente como `v2.0.0` em `/v2`

## Consequências práticas

- usuários existentes da v1 continuam importando `github.com/Lucas-Palomo/go-resource/pkg/resource`
- a v1 continua adequada para manutenção de baixo risco
- a evolução arquitetural deve acontecer na v2
