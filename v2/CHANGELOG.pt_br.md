# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

Este changelog acompanha o módulo v2 (`github.com/Lucas-Palomo/go-resource/v2`).

## [v2.0.1] - 2026-04-12

### Documentação

- reescreveu o README da v2 para refletir o estado realmente publicado do módulo
- esclareceu instalação, convenções de recurso suportadas e semântica de lookup
- alinhou a referência da API, as notas de arquitetura e o guia de migração com o modelo real de runtime da v2
- sincronizou a documentação em inglês e pt-BR

### Gestão de release

- retractou `v2.0.0` porque sua documentação estava semanticamente incorreta
- estabeleceu `v2.0.1` como ponto de entrada estável da v2

### Runtime

- nenhuma mudança intencional de API ou de comportamento em relação à `v2.0.0`

## [v2.0.0] - 2026-04-11

### Adicionado

- nova API centrada em `resource.New(...Option)`
- `LoadDir` e `LoadFS`
- suporte a `fs.FS` e `embed.FS`
- decoders nativos para JSON, YAML, TOML e `.properties` no estilo Java
- flattening de objetos aninhados em notação por ponto
- composição de namespace por pastas e segmentos do nome do arquivo
- `Lookup`, `LookupFor`, `Get` e `GetFor`
- estratégias configuráveis para chave ausente e chave duplicada
- helpers de runtime como `Has`, `HasFor`, `Locales`, `Catalog` e `Reset`
