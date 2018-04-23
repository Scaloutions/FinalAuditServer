package main

import (
	"os"
	"strconv"
	"time"

	"github.com/golang/glog"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	Database   = "auditServerEvents"
	Collection = "systemEvents_collection"
	MongoHost  = "mongodb:27017"
)

func getMongoHost() []string {
	prodEnv, err := strconv.ParseBool(os.Getenv("DEV_ENVIRONMENT"))

	if err != nil {
		glog.Info("Error reading env file! Using prod config")
		return []string{MongoHost}
	} else {
		glog.Info("PRODUCTION ENV: ", prodEnv)
	}

	if prodEnv == true {
		host := os.Getenv("MONGO_PROD_HOST") + os.Getenv("MONGO_PORT")
		glog.Info("Using mongo host:", host)
		return []string{host}
	} else {
		host := os.Getenv("MONGO_DEV_HOST") + os.Getenv("MONGO_PORT")
		glog.Info("Using mongo host:", host)
		return []string{}
	}
}

func getDBCollection() *mgo.Collection {
	glog.Info("Contacting mongo..")
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    getMongoHost(),
		Timeout:  10 * time.Second,
		Database: Database,
	}

	glog.Info("About to dial with info:", mongoDBDialInfo)
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		glog.Info("ERROR: CreateSession: %s\n", err)
	} else {
		eventsCollection := mongoSession.DB(Database).C(Collection)
		index := mgo.Index{
			Key:        []string{"eventtype"},
			Unique:     false,
			DropDups:   false,
			Background: false, // See notes.
			Sparse:     false,
		}

		eventsCollection.EnsureIndex(index)
		return eventsCollection
	}
	return nil
}

func findAllAndLogToFile(eventsCollection *mgo.Collection) {
	glog.Info("Retrieving all records")
	var result []interface{}
	iter := eventsCollection.Find(nil)

	glog.Info("Done retrieving")
	err := iter.All(&result)
	if err != nil {
		glog.Info("Error: ", err)
	} else {
		glog.Info("Logging events to XML file")
		openXMLFile()
		for event := range result {
			var tempEvent GenericEventType
			bsonBytes, _ := bson.Marshal(result[event])
			bson.Unmarshal(bsonBytes, &tempEvent)
			logEventToXML(tempEvent, result[event])
		}
		closeXMLFile()
		glog.Info("========= XML File is ready! :D ========= ")
	}
}

func insertRecord(eventsCollection *mgo.Collection, logMsg interface{}) {
	if err := eventsCollection.Insert(logMsg); err != nil {
		glog.Info("Error in insertion : ", err)
	} else {
		glog.Info("SUCCESS: Inserted event - ", logMsg)
	}
}
