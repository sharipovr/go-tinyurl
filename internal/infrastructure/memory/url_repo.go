package memory

import (
	"errors"
	"sync"

	"github.com/sharipovr/go-tinyurl/internal/domain/url"
)

var ErrNotFound = errors.New("url not found")

type URLRepo struct {
	mu   sync.RWMutex
	data map[string]url.URL
}

func NewURLRepo() *URLRepo {
	return &URLRepo{data: make(map[string]url.URL)}
}

func (r *URLRepo) Save(u url.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[u.ID] = u
	return nil
}

func (r *URLRepo) Find(id string) (url.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.data[id]
	if !ok {
		return url.URL{}, ErrNotFound
	}
	return u, nil
}
