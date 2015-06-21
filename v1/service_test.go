package v1

// ---
// ---
// ---

import (
	"testing"
	
	// ---
	
	"gopkg.in/mgo.v2/bson"
	
	// ---
	
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	
	// ---
	
	"github.com/websecurify/go-dockertest"
)

// ---
// ---
// ---

func findById(id Id) (StripeUserEntry, error) {
	u := StripeUserEntry{
	}
	
	e := MongoCollection.Find(bson.M{"id": id}).One(&u)
	
	return u, e
}

// ---
// ---
// ---

func stripeCustomer(stripeCustomerId StripeCustomerId) (*stripe.Customer, error) {
	customerParams := &stripe.CustomerParams{
	}
	
	// ---
	
	return customer.Get(string(stripeCustomerId), customerParams)
}

// ---
// ---
// ---

func create(userId UserId, email Email, description Description) (CreateReply, error) {
	s := StripeUserMicroservice{}
	
	a := CreateArgs{
		Email: email,
		UserId: userId,
		Description: description,
	}
	
	r := CreateReply{
	}
	
	e := s.Create(nil, &a, &r)
	
	return r, e
}

func destroy(id Id) (DestroyReply, error) {
	s := StripeUserMicroservice{}
	
	a := DestroyArgs{
		Id: id,
	}
	
	r := DestroyReply{
	}
	
	e := s.Destroy(nil, &a, &r)
	
	return r, e
}

func destroyByUserId(userId UserId) (DestroyByUserIdReply, error) {
	s := StripeUserMicroservice{}
	
	a := DestroyByUserIdArgs{
		UserId: userId,
	}
	
	r := DestroyByUserIdReply{
	}
	
	e := s.DestroyByUserId(nil, &a, &r)
	
	return r, e
}

func destroyByStripeCustomerId(stripeCustomerId StripeCustomerId) (DestroyByStripeCustomerIdReply, error) {
	s := StripeUserMicroservice{}
	
	a := DestroyByStripeCustomerIdArgs{
		StripeCustomerId: stripeCustomerId,
	}
	
	r := DestroyByStripeCustomerIdReply{
	}
	
	e := s.DestroyByStripeCustomerId(nil, &a, &r)
	
	return r, e
}

func query(id Id) (QueryReply, error) {
	s := StripeUserMicroservice{}
	
	a := QueryArgs{
		Id: id,
	}
	
	r := QueryReply{
	}
	
	e := s.Query(nil, &a, &r)
	
	return r, e
}

func queryByUserId(userId UserId) (QueryByUserIdReply, error) {
	s := StripeUserMicroservice{}
	
	a := QueryByUserIdArgs{
		UserId: userId,
	}
	
	r := QueryByUserIdReply{
	}
	
	e := s.QueryByUserId(nil, &a, &r)
	
	return r, e
}

func queryByStripeCustomerId(stripeCustomerId StripeCustomerId) (QueryByStripeCustomerIdReply, error) {
	s := StripeUserMicroservice{}
	
	a := QueryByStripeCustomerIdArgs{
		StripeCustomerId: stripeCustomerId,
	}
	
	r := QueryByStripeCustomerIdReply{
	}
	
	e := s.QueryByStripeCustomerId(nil, &a, &r)
	
	return r, e
}

// ---
// ---
// ---

func TestEndToEnd(t *testing.T) {
	cid, cip := dockertest.SetupMongoContainer(t)
	
	defer cid.KillRemove(t)
	
	// ---
	
	MongoServers = cip
	MongoDatabase = "testing"
	
	InitMongo()
	InitStripe()
	
	// ---
	
	userId := UserId("test")
	email := Email("test@test")
	description := Description("desc")
	
	// ---
	
	cr, ce := create(userId, email, description)
	
	if ce != nil {
		t.Error(ce)
	}
	
	// ---
	
	fr, fe := findById(cr.Id)
	
	if fe != nil {
		t.Error(fe)
	}
	
	if fr.UserId != userId {
		t.Error("user id mismatch")
	}
	
	if fr.StripeCustomerId == "" {
		t.Error("stripe customer id mismatch")
	}
	
	// ---
	
	qr, qe := query(cr.Id)
	
	if qe != nil {
		t.Error(qe)
	}
	
	if qr.UserId != userId {
		t.Error("user id mismatch")
	}
	
	if qr.StripeCustomerId == "" {
		t.Error("stripe customer id mismatch")
	}
	
	// ---
	
	qbuir, qbuie := queryByUserId(userId)
	
	if qbuie != nil {
		t.Error(qbuie)
	}
	
	if qbuir.Id != cr.Id {
		t.Error("user id mismatch")
	}
	
	if qbuir.StripeCustomerId == "" {
		t.Error("stripe customer id mismatch")
	}
	
	// ---
	
	qbscir, qbscie := queryByStripeCustomerId(qbuir.StripeCustomerId)
	
	if qbscie != nil {
		t.Error(qbscie)
	}
	
	if qbscir.Id != cr.Id {
		t.Error("user id mismatch")
	}
	
	if qbscir.UserId != userId {
		t.Error("stripe customer id mismatch")
	}
	
	// ---
	
	_, de := destroy(cr.Id)
	
	if de != nil {
		t.Error(de)
	}
	
	// ---
	
	_, fe2 := findById(cr.Id)
	
	if fe2 == nil {
		if fe2 != nil {
			t.Error(fe2)
		} else {
			t.Error("entry found")
		}
	}
	
	// ---
	
	_, ce2 := create(userId, email, description)
	
	if ce2 != nil {
		t.Error(ce2)
	}
	
	// ---
	
	_, de2 := destroyByUserId(userId)
	
	if de2 != nil {
		t.Error(de2)
	}
	
	// ---
	
	cr3, ce3 := create(userId, email, description)

	if ce3 != nil {
		t.Error(ce3)
	}
	
	// ---
	
	_, de3 := destroyByStripeCustomerId(cr3.StripeCustomerId)
	
	if de3 != nil {
		t.Error(de3)
	}
}

// ---
