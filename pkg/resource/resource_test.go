package resource

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/text/language"
)

func writeFile(t *testing.T, root, name, content string) {
	t.Helper()

	fullPath := filepath.Join(root, name)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}

	if err := os.WriteFile(fullPath, []byte(content), 0o644); err != nil {
		t.Fatalf("write failed: %v", err)
	}
}

func TestLoadWithErrorDoesNotPanicOnMissingDirectory(t *testing.T) {
	bundle := NewBundle(filepath.Join(t.TempDir(), "missing"), language.English)

	if err := bundle.LoadWithError(); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestLoadDoesNotPanicAndStoresError(t *testing.T) {
	bundle := NewBundle(filepath.Join(t.TempDir(), "missing"), language.English)

	bundle.Load()

	if bundle.Err() == nil {
		t.Fatal("expected stored error, got nil")
	}
}

func TestGetFallsBackToDefaultLocale(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "en.json", `{"title":"Hello"}`)
	writeFile(t, root, "pt_BR.json", `{"other":"Outro"}`)

	bundle := NewBundle(root, language.English)
	if err := bundle.LoadWithError(); err != nil {
		t.Fatalf("load failed: %v", err)
	}

	bundle.SetLocale(language.BrazilianPortuguese)

	got := bundle.Get("title")
	if got != "Hello" {
		t.Fatalf("expected fallback to default locale, got %q", got)
	}
}

func TestUnsupportedExtensionReturnsError(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "en.txt", `title=hello`)

	bundle := NewBundle(root, language.English)
	if err := bundle.LoadWithError(); err == nil {
		t.Fatal("expected unsupported extension error, got nil")
	}
}
