package resource

import (
	"os"
	"path/filepath"
	"testing"

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

func TestBundleDuplicateKeys(t *testing.T) {
	root := t.TempDir()

	mustWriteFile(t, filepath.Join(root, "en.json"), `{"hello": "Hello"}`)
	mustWriteFile(t, filepath.Join(root, "en.messages.json"), `{"hello": "Hi"}`)

	bundle := New(
		WithFallbackLocale(language.English),
		WithDuplicateKeyStrategy(ErrorOnDuplicate),
	)

	if err := bundle.LoadDir(root); err != nil {
		t.Fatalf("LoadDir() unexpected error = %v", err)
	}

	mustWriteFile(t, filepath.Join(root, "more", "en.json"), `{"hello": "Hello again"}`)
	if err := bundle.LoadDir(root); err == nil {
		t.Fatal("LoadDir() error = nil, want duplicate key error")
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
