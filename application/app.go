package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

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

// This method belongs to App struct
func (a *App) Start(ctx context.Context) error {
	server := http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("> starting server")

	// Make channel to send info back to main thread
	ch := make(chan error, 1)

	// Run server concurrently
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			// Send arrow to anyone who is listening to this buffer channel
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		// Channel is done, so inform reader to stop expecting data
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}

	// Return `nil` for success case
	return nil
}
