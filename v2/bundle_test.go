package resource

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"golang.org/x/text/language"
)

func TestBundleFallbackChain(t *testing.T) {
	root := t.TempDir()

	mustWriteFile(t, filepath.Join(root, "en.json"), `{"hello":"Hello, %s", "checkout": {"title": "Checkout"}}`)
	mustWriteFile(t, filepath.Join(root, "pt-BR.json"), `{"hello":"Olá, %s"}`)

	bundle := New(
		WithFallbackLocale(language.English),
		WithLocale(language.MustParse("pt-BR")),
	)

	if err := bundle.LoadDir(root); err != nil {
		t.Fatalf("LoadDir() error = %v", err)
	}

	got, err := bundle.Lookup("checkout.title")
	if err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}
	if got != "Checkout" {
		t.Fatalf("Lookup() = %q, want %q", got, "Checkout")
	}

	got, err = bundle.Lookup("hello", "Lucas")
	if err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}
	if got != "Olá, Lucas" {
		t.Fatalf("Lookup() = %q, want %q", got, "Olá, Lucas")
	}
}

func TestBundleNamespacedFolders(t *testing.T) {
	root := t.TempDir()
	mustWriteFile(t, filepath.Join(root, "errors", "en.json"), `{"validation": {"required": "required"}}`)

	bundle := New(WithFallbackLocale(language.English))
	if err := bundle.LoadDir(root); err != nil {
		t.Fatalf("LoadDir() error = %v", err)
	}

	got, err := bundle.LookupFor(language.English, "errors.validation.required")
	if err != nil {
		t.Fatalf("LookupFor() error = %v", err)
	}
	if got != "required" {
		t.Fatalf("LookupFor() = %q, want %q", got, "required")
	}
}

func TestBundleLoadFS(t *testing.T) {
	fsys := fstest.MapFS{
		"resources/en.messages.checkout.json": {
			Data: []byte(`{"hello":"Hello"}`),
		},
		"resources/pt_BR.messages.checkout.json": {
			Data: []byte(`{"hello":"Olá"}`),
		},
	}

	bundle := New(
		WithFallbackLocale(language.English),
		WithLocale(language.BrazilianPortuguese),
	)

	if err := bundle.LoadFS(fsys, "resources"); err != nil {
		t.Fatalf("LoadFS() error = %v", err)
	}

	got, err := bundle.Lookup("messages.checkout.hello")
	if err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}
	if got != "Olá" {
		t.Fatalf("Lookup() = %q, want %q", got, "Olá")
	}
}

func TestBundleDuplicateKeysOnReload(t *testing.T) {
	root := t.TempDir()

	mustWriteFile(t, filepath.Join(root, "en.json"), `{"hello": "Hello"}`)

	bundle := New(
		WithFallbackLocale(language.English),
		WithDuplicateKeyStrategy(ErrorOnDuplicate),
	)

	if err := bundle.LoadDir(root); err != nil {
		t.Fatalf("LoadDir() unexpected error = %v", err)
	}

	if err := bundle.LoadDir(root); err == nil {
		t.Fatal("LoadDir() error = nil, want duplicate key error")
	}
}

func TestBundleHasAndReset(t *testing.T) {
	root := t.TempDir()
	mustWriteFile(t, filepath.Join(root, "en.json"), `{"title":"Checkout"}`)

	bundle := New(WithFallbackLocale(language.English))
	if err := bundle.LoadDir(root); err != nil {
		t.Fatalf("LoadDir() error = %v", err)
	}

	if !bundle.Has("title") {
		t.Fatal("Has(title) = false, want true")
	}
	if bundle.Has("missing") {
		t.Fatal("Has(missing) = true, want false")
	}

	bundle.Reset()

	if bundle.Loaded() {
		t.Fatal("Loaded() = true after Reset(), want false")
	}
	if bundle.Has("title") {
		t.Fatal("Has(title) = true after Reset(), want false")
	}
}

func TestBundleMissingKeyError(t *testing.T) {
	root := t.TempDir()
	mustWriteFile(t, filepath.Join(root, "en.json"), `{"title":"Checkout"}`)

	bundle := New(
		WithFallbackLocale(language.English),
		WithMissingKeyStrategy(ErrorOnMissing),
	)

	if err := bundle.LoadDir(root); err != nil {
		t.Fatalf("LoadDir() error = %v", err)
	}

	_, err := bundle.Lookup("missing")
	if err == nil {
		t.Fatal("Lookup() error = nil, want missing key error")
	}
	if !errors.Is(err, ErrMissingKey) {
		t.Fatalf("Lookup() error = %v, want ErrMissingKey", err)
	}
}

func TestBundleLoadProperties(t *testing.T) {
	root := t.TempDir()

	mustWriteFile(t, filepath.Join(root, "en.properties"), strings.Join([]string{
		"# comments are ignored",
		"title = Checkout",
		"checkout.hello = Hello, %s",
		"errors.validation.required: Required",
		"checkout.description = Confirm your order \\",
		"  before leaving",
		"welcome = Ol\\u00E1",
	}, "\n"))
	mustWriteFile(t, filepath.Join(root, "pt_BR.properties"), strings.Join([]string{
		"title = Finalizar compra",
		"checkout.hello = Olá, %s",
	}, "\n"))

	bundle := New(
		WithFallbackLocale(language.English),
		WithLocale(language.BrazilianPortuguese),
	)

	if err := bundle.LoadDir(root); err != nil {
		t.Fatalf("LoadDir() error = %v", err)
	}

	got, err := bundle.Lookup("title")
	if err != nil {
		t.Fatalf("Lookup(title) error = %v", err)
	}
	if got != "Finalizar compra" {
		t.Fatalf("Lookup(title) = %q, want %q", got, "Finalizar compra")
	}

	got, err = bundle.Lookup("errors.validation.required")
	if err != nil {
		t.Fatalf("Lookup(errors.validation.required) error = %v", err)
	}
	if got != "Required" {
		t.Fatalf("Lookup(errors.validation.required) = %q, want %q", got, "Required")
	}

	got, err = bundle.Lookup("checkout.description")
	if err != nil {
		t.Fatalf("Lookup(checkout.description) error = %v", err)
	}
	if got != "Confirm your order before leaving" {
		t.Fatalf("Lookup(checkout.description) = %q, want %q", got, "Confirm your order before leaving")
	}

	got, err = bundle.LookupFor(language.English, "welcome")
	if err != nil {
		t.Fatalf("LookupFor(welcome) error = %v", err)
	}
	if got != "Olá" {
		t.Fatalf("LookupFor(welcome) = %q, want %q", got, "Olá")
	}
}

func TestBundleLoadPropertiesNamespaced(t *testing.T) {
	root := t.TempDir()
	mustWriteFile(t, filepath.Join(root, "messages", "en.properties"), "checkout.title = Checkout")

	bundle := New(WithFallbackLocale(language.English))
	if err := bundle.LoadDir(root); err != nil {
		t.Fatalf("LoadDir() error = %v", err)
	}

	got, err := bundle.LookupFor(language.English, "messages.checkout.title")
	if err != nil {
		t.Fatalf("LookupFor() error = %v", err)
	}
	if got != "Checkout" {
		t.Fatalf("LookupFor() = %q, want %q", got, "Checkout")
	}
}

func mustWriteFile(t *testing.T, filename string, content string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(filename), 0o755); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}
	if err := os.WriteFile(filename, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
}
