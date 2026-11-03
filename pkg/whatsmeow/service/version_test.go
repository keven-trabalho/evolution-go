package whatsmeow_service

import (
	"testing"

	"go.mau.fi/whatsmeow/store"
)

func TestApplyWhatsAppHandshakeVersionOverridesLibraryDefault(t *testing.T) {
	original := store.GetWAVersion()
	t.Cleanup(func() {
		store.SetWAVersion(original)
	})

	applied, ok := applyWhatsAppHandshakeVersion(clientVersion{
		Major: 2,
		Minor: 3000,
		Patch: 1052855845,
	})
	if !ok {
		t.Fatal("expected version to be applied")
	}
	if applied != "2.3000.1052855845" {
		t.Fatalf("applied version = %s", applied)
	}
	if got := store.GetWAVersion().String(); got != "2.3000.1052855845" {
		t.Fatalf("handshake version = %s", got)
	}
}

func TestApplyWhatsAppHandshakeVersionIgnoresZero(t *testing.T) {
	original := store.GetWAVersion()
	t.Cleanup(func() {
		store.SetWAVersion(original)
	})

	applied, ok := applyWhatsAppHandshakeVersion(clientVersion{})
	if ok || applied != "" {
		t.Fatalf("zero version applied = %q ok=%v", applied, ok)
	}
	if got := store.GetWAVersion(); got != original {
		t.Fatalf("library default changed from %s to %s", original, got)
	}
}
