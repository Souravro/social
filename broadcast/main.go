package main

import (
	// "context"
	"fmt"
	"log"
	"net/http"
	"social/broadcast/routes"
)

func main() {
	log.Println("Starting broadcast service...")
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// register broadcast routes
	routes.RegisterRoutes()

	// Start http server
	fmt.Printf("Starting server at port 2121...\n")
	if err := http.ListenAndServe(":2121", nil); err != nil {
		log.Fatal(err)
	}
}
