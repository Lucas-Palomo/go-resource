# Migração da v1 para a v2

## Resumo executivo

A v2 muda o projeto de um pacote pequeno, útil para cenários simples, para uma base mais sólida para uso real em bibliotecas, frameworks e aplicações maiores.

A principal mudança não é cosmética. Ela corrige o contrato da API.

## Problemas centrais da v1

### 1. `panic` como fluxo de erro

Na v1, falhas de leitura de arquivo, parsing e formatos inválidos derrubam a aplicação. Para biblioteca isso é ruim. O consumidor precisa decidir o que fazer com o erro.

### 2. Bug de fallback

Na v1, `Get()` tenta fazer fallback para a locale default quando a locale atual não resolve a mensagem. O problema é que `GetWithLocale()` devolve a própria chave quando a locale não existe, então o resultado não fica vazio e o fallback não acontece como esperado.

### 3. Responsabilidades acopladas

O mesmo arquivo mistura:

- walking de diretórios
- parsing de formatos
- merge de mensagens
- escolha de locale
- formatação de string

Isso dificulta evolução, teste e manutenção.

### 4. Sem `fs.FS`

A v1 prende o consumo ao filesystem tradicional. Isso limita uso com `embed.FS`.

### 5. Sem política formal para colisão e chave ausente

Quando uma key se repete, o comportamento não é uma decisão explícita da API. Quando uma key não existe, o retorno também não é configurável.

### 6. Sem suporte a objetos aninhados

A v1 trabalha na prática com `map[string]string`. Isso limita organização de catálogos maiores.

## Quebras de compatibilidade

### 1. Path do módulo

A v2 segue o padrão semântico de major version do Go:

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## 2. Inicialização

### v1

```go
bundle := resource.NewBundle("./resources", language.English)
bundle.Load()
```

### v2

```go
bundle := resource.New(
    resource.WithFallbackLocale(language.English),
)

if err := bundle.LoadDir("./resources"); err != nil {
    panic(err)
}
```

## 3. Carregamento

A v2 separa o carregamento por fonte:

- `LoadDir(root string)`
- `LoadFS(fsys fs.FS, root string)`

## 4. Lookup

### v1

```go
bundle.Get("hello", "Lucas")
bundle.GetWithLocale(language.English, "hello", "Lucas")
```

### v2

```go
bundle.Get("hello", "Lucas")
bundle.GetFor(language.English, "hello", "Lucas")

value, err := bundle.Lookup("hello", "Lucas")
value, err := bundle.LookupFor(language.English, "hello", "Lucas")
```

`Lookup` é a API preferível quando você quer controle explícito de erro.

## 5. Estruturas aninhadas

### v1

```json
{
  "hello": "Hello"
}
```

### v2

Além do formato plano, a v2 também aceita:

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

Isso vira:

- `checkout.title`
- `checkout.button.confirm`

## 6. Namespacing por pasta e nome de arquivo

Exemplo:

```text
resources/errors/en.json
```

com conteúdo:

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

## Novas capacidades

### Estratégia de chave ausente

```go
resource.WithMissingKeyStrategy(resource.ReturnKeyOnMissing)
resource.WithMissingKeyStrategy(resource.ReturnEmptyOnMissing)
resource.WithMissingKeyStrategy(resource.ErrorOnMissing)
```

### Estratégia de chave duplicada

```go
resource.WithDuplicateKeyStrategy(resource.OverwriteOnDuplicate)
resource.WithDuplicateKeyStrategy(resource.ErrorOnDuplicate)
```

### Registro de decoder customizado

```go
bundle.RegisterDecoder(".ini", myDecoder)
```

## Estratégia recomendada para adoção

### Migração conservadora

- mantenha a estrutura atual de arquivos
- troque o import para `/v2`
- inicialize com `resource.New(...)`
- troque `Load()` por `LoadDir()`
- continue usando `Get()` no primeiro momento
- depois evolua para `Lookup()` onde precisar de controle fino

### Migração estrutural

- reorganize mensagens por namespace
- converta arquivos planos para objetos aninhados
- habilite `ErrorOnDuplicate`
- habilite `ErrorOnMissing` em ambiente de teste

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

Use a v2 como base definitiva. A v1 pode continuar existindo para compatibilidade, mas não deve ser a linha principal de evolução do projeto.
