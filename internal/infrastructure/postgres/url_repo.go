package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sharipovr/go-tinyurl/internal/domain/url"
)

type URLRepo struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *URLRepo { return &URLRepo{pool} }

func (r *URLRepo) Save(u url.URL) error {
	const q = `INSERT INTO urls(id, original, created_at) VALUES ($1,$2,$3)`
	_, err := r.pool.Exec(context.Background(), q, u.ID, u.Original, u.CreatedAt)
	return err
}

func (r *URLRepo) Find(id string) (url.URL, error) {
	const q = `SELECT id, original, created_at FROM urls WHERE id=$1`
	var u url.URL
	err := r.pool.QueryRow(context.Background(), q, id).Scan(&u.ID, &u.Original, &u.CreatedAt)
	return u, err
}
