package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}
}

// Single handler -> for proper setting you can do it manually
// but it is better to use a library for faster development...
// Lets use Chi
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
