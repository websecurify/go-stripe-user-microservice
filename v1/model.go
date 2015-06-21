package v1

// ---
// ---
// ---

import (
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
