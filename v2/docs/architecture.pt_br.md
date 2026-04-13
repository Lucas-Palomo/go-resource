# Notas de Arquitetura da v2

[English](./architecture.md) | [Português (Brasil)](./architecture.pt_br.md)

## Objetivo de design

A arquitetura da v2 é deliberadamente contida: manter i18n baseado em arquivos simples, mas com comportamento de runtime explícito o suficiente para reutilização em aplicações reais.

A biblioteca gira em torno de uma ideia central: carregar arquivos de recurso para um catálogo normalizado em memória e então resolver chaves por uma cadeia de fallback de locale.

## Principais peças de runtime

- `Bundle` controla os catálogos, o registro de decoders, os locales e as estratégias de lookup
- decoders transformam arquivos brutos em `map[string]any`
- o loader normaliza o conteúdo decodificado em `map[string]string`
- o lookup percorre fallbacks de locale construídos com `golang.org/x/text/language`

## Pipeline de carregamento

Para cada arquivo de recurso, o loader executa estes passos:

1. seleciona o decoder pela extensão do arquivo
2. interpreta o locale a partir do nome do arquivo
3. deriva o namespace a partir de diretórios e segmentos do nome do arquivo
4. decodifica o arquivo em um grafo genérico de objetos
5. achata objetos aninhados em chaves com notação por ponto
6. normaliza valores escalares para string
7. valida chaves duplicadas segundo a estratégia configurada
8. faz merge dos catálogos pendentes na memória do bundle

## Modelo de chave

A chave final de lookup pode nascer de três camadas:

- pastas
- segmentos de namespace do nome do arquivo
- objetos aninhados no corpo do arquivo

Exemplos:

```text
resources/errors/en.json            -> errors.*
resources/en.messages.checkout.toml -> messages.checkout.*
```

Esse modelo mantém as chaves previsíveis mesmo quando o projeto mistura organização por diretório e aninhamento dentro do documento.

## Normalização de escalares

O catálogo de runtime é sempre `map[string]string`.

Isso é intencional:

- lookups permanecem determinísticos
- a formatação com `fmt.Sprintf` fica direta
- o chamador não precisa raciocinar sobre tipos escalares mistos no momento da leitura

Números e booleanos viram string. Valores `nil` e folhas compostas não suportadas são rejeitados.

## Resolução de locale

O lookup segue uma cadeia de locale, não apenas um casamento exato.

Na prática, a resolução usa:

1. o locale solicitado
2. seus locales pai
3. o locale de fallback configurado
4. os pais do fallback
5. `language.Und`

Isso permite manter um locale específico para sobrescritas sem perder a queda para catálogos mais amplos.

## Modelo de concorrência

`Bundle` usa um read-write mutex.

- caminhos de leitura como `Lookup`, `Get`, `Has`, `Catalog` e `Locales` usam read lock
- caminhos de mutação como `LoadFS`, `SetLocale`, `Reset` e `RegisterDecoder` usam write lock

A API pública é segura para leituras concorrentes após o carregamento.

## Semântica de merge

`LoadDir` e `LoadFS` são operações incrementais.

Um segundo carregamento bem-sucedido não substitui o bundle atual em memória. Ele faz merge de catálogos adicionais sobre o que já está carregado.

Esse comportamento é útil para fontes em camadas, mas muda a forma de desenhar fluxos de reload:

- chame `Reset()` antes de carregar novamente quando quiser substituição completa
- sob `ErrorOnDuplicate`, a validação de duplicidade contra catálogos já carregados termina antes de qualquer mutação no novo carregamento

## Pontos de extensão de decoder

`RegisterDecoder` permite adicionar suporte a formatos extras.

Detalhes do contrato:

- extensões são normalizadas para a forma minúscula com ponto
- extensões vazias são ignoradas
- decoders nil são ignorados
- extensões não suportadas fazem o carregamento falhar com `ErrUnsupportedFormat`

## Contrato de lookup

A API expõe de propósito dois estilos de lookup:

- `Lookup` e `LookupFor` são explícitos e conscientes de erro
- `Get` e `GetFor` são helpers lenientes que retornam a chave original em erros de lookup

Essa divisão entrega ergonomia para textos de UI sem perder um caminho estrito para serviços, validação e testes.
