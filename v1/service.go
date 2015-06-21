package v1

// ---
// ---
// ---

import (
	"sync"
	"net/http"
	
	// ---
	
	"gopkg.in/mgo.v2/bson"
	
	// ---
	
	"code.google.com/p/go-uuid/uuid"
	
	// ---
	
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

// ---
// ---
// ---

type StripeUserMicroservice struct {
}

// ---
// ---
// ---

type CreateArgs struct {
	UserId UserId `json:"userId"`
	Email Email `json:"email"`
	Description Description `json:"description"`
}

type CreateReply struct {
	Id Id `json:"id"`
	StripeCustomerId StripeCustomerId `json:"stripeCustomerId"`
}

// ---

func (s *StripeUserMicroservice) Create(r *http.Request, args *CreateArgs, reply *CreateReply) (error) {
	id := Id(uuid.NewRandom().String())
	
	// ---
	
	customerParams := &stripe.CustomerParams{
	}
	
	customer, customerErr := customer.New(customerParams)
	
	if customerErr != nil {
		return customerErr
	}
	
	// ---
	
	entry := StripeUserEntry{
		ObjectId: bson.NewObjectId(),
		Id: id,
		UserId: args.UserId,
		StripeCustomerId: StripeCustomerId(customer.ID),
	}
	
	// ---
	
	insertErr := MongoCollection.Insert(entry)
	
	if insertErr != nil {
		return insertErr
	}
	
	// ---
	
	reply.Id = entry.Id
	reply.StripeCustomerId = entry.StripeCustomerId
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type DestroyArgs struct {
	Id Id `json:"id"`
}

type DestroyReply struct {
}

// ---

func (s *StripeUserMicroservice) Destroy(r *http.Request, args *DestroyArgs, reply *DestroyReply) (error) {
	result := StripeUserEntry{}
	
	// ---
	
	findErr := MongoCollection.Find(bson.M{"id": args.Id}).One(&result)
	
	if findErr != nil {
		return findErr
	}
	
	// ---
	
	errChan := make(chan error)
	
	// ---
	
	var wg sync.WaitGroup
	
	// ---
	
	wg.Add(1)
	
	go func () {
		defer wg.Done()
		
		removeErr := MongoCollection.Remove(bson.M{"id": args.Id})
		
		if removeErr != nil {
			errChan <- removeErr
		}
	}()
	
	wg.Add(1)
	
	go func () {
		defer wg.Done()
		
		delErr := customer.Del(string(result.StripeCustomerId))
		
		if delErr != nil {
			errChan <- delErr
		}
	}()
	
	// ---
	
	wg.Wait()
	
	// ---
	
	close(errChan)
	
	// ---
	
	for err := range errChan {
		return err
	}
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type DestroyByUserIdArgs struct {
	UserId UserId `json:"userId"`
}

type DestroyByUserIdReply struct {
}

// ---

func (s *StripeUserMicroservice) DestroyByUserId(r *http.Request, args *DestroyByUserIdArgs, reply *DestroyByUserIdReply) (error) {
	result := StripeUserEntry{}
	
	// ---
	
	findErr := MongoCollection.Find(bson.M{"userId": args.UserId}).One(&result)
	
	if findErr != nil {
		return findErr
	}
	
	// ---
	
	errChan := make(chan error)
	
	// ---
	
	var wg sync.WaitGroup
	
	// ---
	
	wg.Add(1)
	
	go func () {
		defer wg.Done()
		
		removeErr := MongoCollection.Remove(bson.M{"id": result.Id})
		
		if removeErr != nil {
			errChan <- removeErr
		}
	}()
	
	wg.Add(1)
	
	go func () {
		defer wg.Done()
		
		delErr := customer.Del(string(result.StripeCustomerId))
		
		if delErr != nil {
			errChan <- delErr
		}
	}()
	
	// ---
	
	wg.Wait()
	
	// ---
	
	close(errChan)
	
	// ---
	
	for err := range errChan {
		return err
	}
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type DestroyByStripeCustomerIdArgs struct {
	StripeCustomerId StripeCustomerId `json:"stripeCustomerId"`
}

type DestroyByStripeCustomerIdReply struct {
}

// ---

func (s *StripeUserMicroservice) DestroyByStripeCustomerId(r *http.Request, args *DestroyByStripeCustomerIdArgs, reply *DestroyByStripeCustomerIdReply) (error) {
	errChan := make(chan error)
	
	// ---
	
	var wg sync.WaitGroup
	
	// ---
	
	wg.Add(1)
	
	go func () {
		defer wg.Done()
		
		removeErr := MongoCollection.Remove(bson.M{"stripeCustomerId": args.StripeCustomerId})
		
		if removeErr != nil {
			errChan <- removeErr
		}
	}()
	
	wg.Add(1)
	
	go func () {
		defer wg.Done()
		
		delErr := customer.Del(string(args.StripeCustomerId))
		
		if delErr != nil {
			errChan <- delErr
		}
	}()
	
	// ---
	
	wg.Wait()
	
	// ---
	
	close(errChan)
	
	// ---
	
	for err := range errChan {
		return err
	}
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type QueryArgs struct {
	Id Id `json:"id"`
}

type QueryReply struct {
	UserId UserId `json:"userId"`
	StripeCustomerId StripeCustomerId `json:"stripeCustomerId"`
}

// ---

func (s *StripeUserMicroservice) Query(r *http.Request, args *QueryArgs, reply *QueryReply) (error) {
	result := StripeUserEntry{}
	
	// ---
	
	findErr := MongoCollection.Find(bson.M{"id": args.Id}).One(&result)
	
	if findErr != nil {
		return findErr
	}
	
	// ---
	
	reply.UserId = result.UserId
	reply.StripeCustomerId = result.StripeCustomerId
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type QueryByUserIdArgs struct {
	UserId UserId `json:"userId"`
}

type QueryByUserIdReply struct {
	Id Id `json:"id"`
	StripeCustomerId StripeCustomerId `json:"stripeCustomerId"`
}

// ---

func (s *StripeUserMicroservice) QueryByUserId(r *http.Request, args *QueryByUserIdArgs, reply *QueryByUserIdReply) (error) {
	result := StripeUserEntry{}
	
	// ---
	
	findErr := MongoCollection.Find(bson.M{"userId": args.UserId}).One(&result)
	
	if findErr != nil {
		return findErr
	}
	
	// ---
	
	reply.Id = result.Id
	reply.StripeCustomerId = result.StripeCustomerId
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type QueryByStripeCustomerIdArgs struct {
	StripeCustomerId StripeCustomerId `json:"stripeCustomerId"`
}

type QueryByStripeCustomerIdReply struct {
	Id Id `json:"id"`
	UserId UserId `json:"userId"`
}

// ---

func (s *StripeUserMicroservice) QueryByStripeCustomerId(r *http.Request, args *QueryByStripeCustomerIdArgs, reply *QueryByStripeCustomerIdReply) (error) {
	result := StripeUserEntry{}
	
	// ---
	
	findErr := MongoCollection.Find(bson.M{"stripeCustomerId": args.StripeCustomerId}).One(&result)
	
	if findErr != nil {
		return findErr
	}
	
	// ---
	
	reply.Id = result.Id
	reply.UserId = result.UserId
	
	// ---
	
	return nil
}

// ---
// ---
// ---

func Start() {
	InitMongo()
	InitStripe()
}

// ---
