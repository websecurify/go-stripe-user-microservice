package v1

// ---
// ---
// ---

import (
	"log"
	
	// ---
	
	"github.com/stripe/stripe-go"
)

// ---
// ---
// ---

func InitStripe() {
	if StripeKey == "" {
		log.Fatal("stripe key not configured")
	}
	
	// ---
	
	stripe.Key = StripeKey
}

// ---
