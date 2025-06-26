package service

import (
	"time"

	"github.com/sharipovr/go-tinyurl/internal/domain/url"
)

type Shortener struct {
	repo url.Repository
	gen  url.IDGenerator
}

func NewShortener(r url.Repository, g url.IDGenerator) *Shortener {
	return &Shortener{repo: r, gen: g}
}

// Create generates ID and saves new record.
func (s *Shortener) Create(original string) (url.URL, error) {
	id, err := s.gen.Generate()
	if err != nil {
		return url.URL{}, err
	}

	u := url.URL{
		ID:        id,
		Original:  original,
		CreatedAt: time.Now().UTC(),
	}
	return u, s.repo.Save(u)
}

// Resolve searches original URL by code.
func (s *Shortener) Resolve(id string) (url.URL, error) {
	return s.repo.Find(id)
}
