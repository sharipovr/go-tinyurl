package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sharipovr/go-tinyurl/internal/domain/url"
	pg "github.com/sharipovr/go-tinyurl/internal/infrastructure/postgres"
	"github.com/sharipovr/go-tinyurl/internal/pkg/db"
)

func main() {
	dsn := "postgres://tinyurl_admin:tiny-tini-123@localhost:5432/tinyurl_db?sslmode=disable"
	pool := db.MustConnect(context.Background(), dsn)
	repo := pg.New(pool)

	u := url.URL{ID: "x1", Original: "https://example.com", CreatedAt: time.Now()}
	if err := repo.Save(u); err != nil {
		log.Fatal("save:", err)
	}

	got, err := repo.Find("x1")
	if err != nil {
		log.Fatal("find:", err)
	}
	fmt.Println("read back:", got.Original)
}
