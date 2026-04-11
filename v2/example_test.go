package resource_test

import (
	"fmt"
	"testing/fstest"

	resource "github.com/Lucas-Palomo/go-resource/v2"
	"golang.org/x/text/language"
)

func ExampleBundle() {
	fsys := fstest.MapFS{
		"resources/en.json": {
			Data: []byte(`{
				"title": "Checkout",
				"checkout": {
					"hello": "Hello, %s"
				}
			}`),
		},
		"resources/errors/en.json": {
			Data: []byte(`{
				"validation": {
					"required": "Required field"
				}
			}`),
		},
		"resources/pt_BR.json": {
			Data: []byte(`{
				"title": "Finalização",
				"checkout": {
					"hello": "Olá, %s"
				}
			}`),
		},
	}

	bundle := resource.New(
		resource.WithFallbackLocale(language.English),
		resource.WithLocale(language.BrazilianPortuguese),
	)

	if err := bundle.LoadFS(fsys, "resources"); err != nil {
		panic(err)
	}

	fmt.Println(bundle.Get("title"))
	fmt.Println(bundle.Get("checkout.hello", "Lucas"))
	fmt.Println(bundle.Get("errors.validation.required"))

	// Output:
	// Finalização
	// Olá, Lucas
	// Required field
}

func ExampleBundle_properties() {
	fsys := fstest.MapFS{
		"resources/en.properties": {
			Data: []byte("title = Checkout\ncheckout.hello = Hello, %s\nerrors.validation.required: Required field\n"),
		},
		"resources/pt_BR.properties": {
			Data: []byte("title = Finalização\ncheckout.hello = Olá, %s\n"),
		},
	}

	bundle := resource.New(
		resource.WithFallbackLocale(language.English),
		resource.WithLocale(language.BrazilianPortuguese),
	)

	if err := bundle.LoadFS(fsys, "resources"); err != nil {
		panic(err)
	}

	fmt.Println(bundle.Get("title"))
	fmt.Println(bundle.Get("checkout.hello", "Lucas"))
	fmt.Println(bundle.Get("errors.validation.required"))

	// Output:
	// Finalização
	// Olá, Lucas
	// Required field
}
