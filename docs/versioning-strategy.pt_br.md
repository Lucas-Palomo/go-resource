# Estratégia de Versionamento e Release

[English](./versioning-strategy.md) | [Português (Brasil)](./versioning-strategy.pt_br.md)

Este repositório usa um **layout multi-major**:

- módulo raiz: `github.com/Lucas-Palomo/go-resource`
- módulo v2: `github.com/Lucas-Palomo/go-resource/v2`

## Por quê

O projeto já havia sido publicado publicamente como módulo v1. Para não quebrar consumidores existentes, a nova major vive em `/v2`.

## Política de retract

O `go.mod` da raiz retrai `v1.0.0` para sinalizar que a primeira release pública não deve ser usada como versão recomendada.

Fluxo sugerido de release:

1. manter a raiz como linha legada da v1
2. publicar uma tag corrigida na raiz, como `v1.0.1`
3. publicar a nova major em `/v2` como `v2.0.0`

## Consequências práticas

- usuários existentes da v1 continuam importando `github.com/Lucas-Palomo/go-resource/pkg/resource`
- novos projetos devem preferir `github.com/Lucas-Palomo/go-resource/v2`
- a v1 pode receber manutenção mínima de compatibilidade enquanto a v2 evolui de forma independente
