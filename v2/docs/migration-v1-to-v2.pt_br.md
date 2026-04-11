# Migração da v1 para a v2

[English](./migration-v1-to-v2.md) | [Português (Brasil)](./migration-v1-to-v2.pt_br.md)

## Resumo executivo

A v2 não é uma reescrita cosmética. Ela corrige o contrato da biblioteca e dá ao projeto uma base mais forte para evolução de longo prazo.

A regra prática é simples:

- fique na v1.0.1 quando compatibilidade for o objetivo principal
- migre para a v2 quando precisar de uma API mais forte e espaço para crescer

## Problemas centrais da v1

### 1. `panic` como fluxo de erro

Historicamente, a v1 expunha falhas de carregamento por `panic`. Esse não é um bom default para uma biblioteca reutilizável.

### 2. Bug de fallback

Na linha antiga, `Get()` tentava fazer fallback para a locale default, mas a interação com `GetWithLocale()` podia impedir o fallback correto quando a locale atual não existia.

### 3. Responsabilidades acopladas

Walking de diretório, parsing, merge de mensagens, escolha de locale e formatação viviam próximos demais.

### 4. Ausência de `fs.FS`

A linha antiga era presa ao filesystem tradicional, o que limitava o uso com `embed.FS`.

### 5. Ausência de política explícita para chave ausente e duplicada

Havia comportamento, mas ele não era uma decisão claramente configurável da API.

### 6. Ausência de suporte de primeira classe para objetos aninhados

Catálogos maiores ficam mais difíceis de organizar quando tudo precisa se comportar como um `map[string]string` plano.

## Breaking changes

### 1. Path do módulo

A v2 segue semantic import versioning do Go:

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

### 2. Inicialização

#### v1

```go
bundle := resource.NewBundle("./resources", language.English)
bundle.Load()
```

#### v2

```go
bundle := resource.New(
	resource.WithFallbackLocale(language.English),
)

if err := bundle.LoadDir("./resources"); err != nil {
	return err
}
```

### 3. API de carregamento

A v2 separa o carregamento por origem:

- `LoadDir(root string)`
- `LoadFS(fsys fs.FS, root string)`

### 4. API de lookup

#### v1

```go
bundle.Get("hello", "Lucas")
bundle.GetWithLocale(language.English, "hello", "Lucas")
```

#### v2

```go
bundle.Get("hello", "Lucas")
bundle.GetFor(language.English, "hello", "Lucas")

value, err := bundle.Lookup("hello", "Lucas")
value, err := bundle.LookupFor(language.English, "hello", "Lucas")
```

Regra prática:

- use `Get` / `GetFor` por conveniência
- use `Lookup` / `LookupFor` quando tratamento explícito de erro for importante

### 5. Estruturas aninhadas

#### v1

```json
{
  "hello": "Hello"
}
```

#### v2

A v2 continua aceitando arquivos planos, mas também aceita objetos aninhados:

```json
{
  "checkout": {
    "title": "Checkout",
    "button": {
      "confirm": "Confirm"
    }
  }
}
```

que viram:

- `checkout.title`
- `checkout.button.confirm`

### 6. Namespace por diretório e nome de arquivo

Exemplo:

```text
resources/errors/en.json
```

com:

```json
{
  "validation": {
    "required": "Required"
  }
}
```

vira:

```text
errors.validation.required
```

### 7. Suporte a `.properties`

A v2 adiciona suporte nativo a arquivos `.properties` no estilo Java. Isso importa quando seus assets de tradução já usam esse formato.

## Novas capacidades da v2

### Estratégia para chave ausente

```go
resource.WithMissingKeyStrategy(resource.ReturnKeyOnMissing)
resource.WithMissingKeyStrategy(resource.ReturnEmptyOnMissing)
resource.WithMissingKeyStrategy(resource.ErrorOnMissing)
```

### Estratégia para chave duplicada

```go
resource.WithDuplicateKeyStrategy(resource.OverwriteOnDuplicate)
resource.WithDuplicateKeyStrategy(resource.ErrorOnDuplicate)
```

### Registro de decoder customizado

```go
bundle.RegisterDecoder(".ini", myDecoder)
```

### Helpers de inspeção e reset

```go
bundle.Has("checkout.title")
bundle.HasFor(language.English, "checkout.title")
bundle.Reset()
```

## Estratégia recomendada de migração

### Migração conservadora

Use esse caminho quando você quer a menor mudança operacional possível.

- troque o import para `/v2`
- inicialize com `resource.New(...)`
- substitua `Load()` por `LoadDir()`
- mantenha o layout atual de arquivos no início
- continue usando `Get()` no primeiro momento
- migre fluxos específicos para `Lookup()` quando controle explícito fizer diferença

### Migração estrutural

Use esse caminho quando você quer extrair mais valor da v2.

- reorganize recursos por namespace
- converta arquivos planos em objetos aninhados quando isso melhorar a legibilidade
- habilite `ErrorOnDuplicate`
- habilite `ErrorOnMissing` em testes ou pipelines de validação
- adote `.properties` quando compatibilidade com Java for importante

## Exemplo completo

```go
bundle := resource.New(
	resource.WithFallbackLocale(language.English),
	resource.WithLocale(language.MustParse("pt-BR")),
	resource.WithMissingKeyStrategy(resource.ErrorOnMissing),
	resource.WithDuplicateKeyStrategy(resource.ErrorOnDuplicate),
)

if err := bundle.LoadDir("./resources"); err != nil {
	return err
}

msg, err := bundle.Lookup("checkout.hello", "Lucas")
if err != nil {
	return err
}

fmt.Println(msg)
```

## Recomendação final

Mantenha a v1 para compatibilidade.

Escolha a v2 como linha principal de evolução.
