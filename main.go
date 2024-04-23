package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/vonderklaas/orders-api/application"
)

func main() {

	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// Call at the end of main function
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
