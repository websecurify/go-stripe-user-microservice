package main // import "github.com/websecurify/go-stripe-user-microservice"

// ---
// ---
// ---

import (
	"os"
	"log"
	"net/http"
	
	// ---
	
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/gorilla/handlers"
	
	// ---
	
	"v1"
)

// ---
// ---
// ---

func main() {
	v1.Start()
	
	// ---
	
	s := rpc.NewServer()
	
	// ---
	
	s.RegisterService(new(v1.StripeUserMicroservice), "1")
	
	// ---
	
	s.RegisterCodec(json.NewCodec(), "application/json")
	
	// ---
	
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, s))
	
	// ---
	
	address := os.Getenv("ADDRESS")
	
	if address == "" {
		address = ":8080"
	}
	
	// ---
	
	log.Println("starting server at", address)
	
	// ---
	
	log.Fatal(http.ListenAndServe(address, nil))
}

// ---
