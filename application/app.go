package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

// Orchestration usually done with KuberNetes and ECs instances for DataStores
type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}

	return app
}

// receiver -> owner
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		// graceful shutdown
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("Starting server")

	// GO routine -> channel usage:
	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			// sending the error to anyone who is listening...
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()
	// select -> works like switch cases only for channels
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
	// return nil
}
