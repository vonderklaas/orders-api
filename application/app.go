package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

// This method belongs to App struct
func (a *App) Start(ctx context.Context) error {
	server := http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()

	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	// Return `nil` for success case
	return nil
}
