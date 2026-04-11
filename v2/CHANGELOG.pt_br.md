# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

Este changelog acompanha o **módulo v2** (`github.com/Lucas-Palomo/go-resource/v2`).

## [v2.0.0] - 2026-04-11

### Adicionado

- nova API centrada em `resource.New(...Option)`
- `LoadDir` e `LoadFS`
- suporte a `fs.FS` e `embed.FS`
- decoders nativos para JSON, YAML, TOML e `.properties` no estilo Java
- flatten de objetos aninhados em notação por ponto
- composição de namespace por pastas e segmentos do nome do arquivo
- `Lookup`, `LookupFor`, `Get` e `GetFor`
- estratégias configuráveis para chave ausente e chave duplicada
- helpers de runtime como `Has`, `HasFor`, `Locales`, `Catalog` e `Reset`
- registro de decoder por meio de `RegisterDecoder`
- documentos de arquitetura, API e migração
- exemplos executáveis para os fluxos plano e `.properties`
