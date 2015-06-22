package v1

// ---
// ---
// ---

import (
	"log"
	
	// ---
	
	"gopkg.in/mgo.v2"
)

// ---
// ---
// ---

var MongoSession *mgo.Session
var MongoCollection *mgo.Collection

// ---
// ---
// ---

func InitMongo() {
	if MongoServers == "" {
		log.Fatal("mongo servers not configured")
	}
	
	if MongoDatabase == "" {
		log.Fatal("mongo database not configured")
	}
	
	// ---
	
	log.Println("connecting to mongo servers", MongoServers)
	
	// ---
	
	session, sessionErr := mgo.Dial(MongoServers)
	
	if sessionErr != nil {
		log.Fatal(sessionErr)
	}
	
	// ---
	
	collection := session.DB(MongoDatabase).C(MongoCollectionName)
	
	// ---
	
	MongoSession = session
	MongoCollection = collection
	
	// ---
	
	initModel()
}

// ---
