package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Nash1818/go_api.git/application"
	// "net/http"
	// "github.com/go-chi/chi/middleware"
	// "github.com/go-chi/chi/v5"
)

func main() {
	app := application.New()
	// Graceful shutdown -> SIGINT (interrupt)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.Start(ctx)
	// err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}

// func main() {
// 	router := chi.NewRouter() //New -> constructor
// 	router.Use(middleware.Logger)
// 	router.Get("/hello", basicHandler)

// 	server := &http.Server{
// 		Addr:    ":3000",
// 		Handler: router,
// 		// Handler: http.HandlerFunc(basicHandler),
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		fmt.Println("Failed to listen to server", err)
// 	}
// }

// // Single handler -> for proper setting you can do it manually
// // but it is better to use a library for faster development...
// // Lets use Chi
// func basicHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello World!"))
// }
