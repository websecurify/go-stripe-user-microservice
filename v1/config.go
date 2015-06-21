package v1

// ---
// ---
// ---

import (
	"os"
)

// ---
// ---
// ---

var StripeKey string
var MongoServers string
var MongoDatabase string

// ---
// ---
// ---

func init() {
	StripeKey = os.Getenv("STRIPE_KEY")
	MongoServers = os.Getenv("MONGO_SERVERS")
	MongoDatabase = os.Getenv("MONGO_DATABASE")
}

// ---
