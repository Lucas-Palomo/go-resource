# Migração da v1 para a v2

[English](./migration-v1-to-v2.md) | [Português (Brasil)](./migration-v1-to-v2.pt_br.md)

## Resumo executivo

Fique na `v1.0.1` quando compatibilidade for o objetivo principal. Migre para a `v2.0.1` quando quiser um contrato mais forte e um modelo de recursos mais capaz.

Se você já adotou a `v2.0.0`, mova para a `v2.0.1` como release estável com documentação corrigida. O alvo de migração e o modelo de runtime continuam os mesmos.

## Mudanças conceituais

| Aspecto | v1 | v2 |
|---|---|---|
| Path do módulo | raiz | `/v2` |
| Construtor | `NewBundle` | `New(...Option)` |
| Carregamento | `Load()` / `LoadWithError()` | `LoadDir()` / `LoadFS()` |
| Filesystem | apenas SO | qualquer `fs.FS` |
| Modelo de recurso | apenas mapas planos de string | objetos aninhados achatados |
| `.properties` | não | sim |
| Namespace por pasta / nome de arquivo | não | sim |
| Política para chave ausente | implícita | configurável |
| Política para chave duplicada | implícita | configurável |

## Diferença comportamental central

Na v1, a estrutura de pastas e os segmentos pontuados do nome do arquivo organizam arquivos, mas **não** viram prefixos de chave em runtime.

Na v2, viram.

## Checklist de migração

- atualize os imports para `/v2`
- substitua `NewBundle` por `New(...Option)`
- substitua `GetWithLocale` por `GetFor` ou `LookupFor`
- use `LoadDir` ou `LoadFS`
- revise catálogos que dependiam apenas de chaves planas
- revise se nomes de pasta ou nomes de arquivo pontuados agora devem virar prefixos de namespace
- escolha explicitamente as estratégias para chave ausente e chave duplicada
