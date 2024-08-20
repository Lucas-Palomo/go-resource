# Go Resource
This project is a Go package designed to provide internationalization resources. 
It fetches resource files in different formats (JSON, YAML or TOML) and uses this data for internationalization.

## How to use

The resource folder needs to be organized semantically, as in the following example

```text
/resources              # Resources root folder
|-- errors              # folder for internationalizing errors
    |-- en.json
    |-- pt_BR.json
|-- messages            # folder for internationalizing messages
    |-- checkout        # message category (example with subfolder)
        |-- en.json
        |-- pt_BR.json
```

Another valid structure is

```text
/resources              # Resources root folder
|-- en.errors.json
|-- en.messages.json
|-- pt_BR.errors.json
|-- pt_BR.messages.json
```
 
### File content

Example of Resources

**Json**
```json
{
  "title": "My title",
  "hello": "Hello, %s"
}
```

**Toml**
```toml
title="My title"
hello="Hello, %s"
```

**Yaml**
```yaml
title: "My title"
hello: "Hello, %s"
```

### Working with the Bundle

```go
package main

import (
	"github.com/Lucas-Palomo/go-resource/pkg/resource"
	"golang.org/x/text/language"
)

func main() {
	bundle := resource.NewBundle("./resource", language.English) // Resources root folder and a fallback language
	bundle.Load() // Load all resources files
	
	println(bundle.Get("title")) 
	// or, to pass arguments
	println(bundle.Get("hello", "Lucas")) // output is "Hello, Lucas"
}
```

### Working with Multi-language

**pt_BR.json**
```json
{
  "title": "Meu título",
  "hello": "Olá, %s"
}
```

**en.json**
```json
{
  "title": "My title",
  "hello": "Hello, %s",
  "monday": "Today is monday"
}
```


```go
package main

import (
	"github.com/Lucas-Palomo/go-resource/pkg/resource"
	"golang.org/x/text/language"
)

func main() {
	bundle := resource.NewBundle("./resource", language.English) // Resources root folder and a fallback language
	bundle.Load()
	
	bundle.SetLocale(language.BrazilianPortuguese) // Now this is the current bundle language
	
	println(bundle.Get("title")) // output is "Meu título"
	// or, to pass arguments
	println(bundle.Get("hello", "Lucas")) // output is "Olá, Lucas"
	
	// Fallback case
	println(bundle.Get("monday")) // output is "Today is monday"
	
	// Unregistered resource key
	println(bundle.Get("friday")) // output is the same key "friday"
	
	// Force locale
	println(bundle.GetWithLocale(language.English, "title")) // output is "My title"
}
```
