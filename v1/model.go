package v1

// ---
// ---
// ---

import (
	"log"
	
	// ---
	
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ---
// ---
// ---

type Id string
type Email string
type UserId string
type Description string
type StripeCustomerId string

// ---
// ---
// ---

type StripeUserEntry struct {
	ObjectId bson.ObjectId `bson:"_id"`
	Id Id `bson:"id"`
	UserId UserId `bson:"userId"`
	StripeCustomerId StripeCustomerId `bson:"stripeCustomerId"`
}

// ---
// ---
// ---

func initModel() {
	index := mgo.Index{
		Key: []string{"id", "userId", "stripeCustomerId"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
	
	// ---
	
	ensureErr := MongoCollection.EnsureIndex(index)
	
	if ensureErr != nil {
		log.Fatal(ensureErr)
	}
}

// ---
