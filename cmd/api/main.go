package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sharipovr/go-tinyurl/internal/config"
	"github.com/sharipovr/go-tinyurl/internal/logger"
	"go.uber.org/zap"
)

func main() {
	cfg := config.MustLoad()
	log := logger.New()

	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	log.Info("starting server", zap.String("port", cfg.Port))
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal("server error", zap.Error(err))
	}
}
