# Changelog

## [v2.0.1] - 2026-04-13

### Documentação

- reescreveu o README da v2 com uma narrativa de release mais clara e um guia de uso mais amigável
- esclareceu instalação, formatos de arquivo suportados, composição de namespace e comportamento de lookup
- alinhou a referência da API, notas de arquitetura e guia de migração ao contrato real de runtime da v2
- sincronizou a documentação em inglês e pt-BR

### Endurecimento de runtime

- passou a ignorar registros inválidos de decoder customizado quando a extensão é vazia ou o decoder é nil
- fez a validação de chave duplicada contra catálogos já carregados acontecer antes de qualquer mutação durante `LoadFS`
- documentou a semântica de merge incremental em carregamentos repetidos

## [v2.0.0] - 2026-04-11

### Adicionado

- nova API centrada em `resource.New(...Option)`
- `LoadDir` e `LoadFS`
- suporte a `fs.FS` e `embed.FS`
- decoders nativos para JSON, YAML, TOML e `.properties` no estilo Java
- achatamento de objetos aninhados em notação por ponto
- composição de namespace a partir de pastas e segmentos do nome do arquivo
- `Lookup`, `LookupFor`, `Get` e `GetFor`
- estratégias configuráveis para chave ausente e chave duplicada
- helpers de runtime como `Has`, `HasFor`, `Locales`, `Catalog` e `Reset`
