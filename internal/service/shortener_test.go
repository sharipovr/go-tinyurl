package service

import (
	"testing"

	"github.com/sharipovr/go-tinyurl/internal/infrastructure/memory"
)

func TestShortener_CreateAndResolve(t *testing.T) {
	repo := memory.NewURLRepo()
	gen := memory.RandIDGen{}
	s := NewShortener(repo, gen)

	orig := "https://example.com/long"
	u, err := s.Create(orig)
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if u.Original != orig {
		t.Errorf("want %q, got %q", orig, u.Original)
	}

	got, err := s.Resolve(u.ID)
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if got.ID != u.ID {
		t.Errorf("resolve returned wrong URL")
	}
}
