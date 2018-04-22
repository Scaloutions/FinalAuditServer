package main

import (
	"time"

	"github.com/golang/glog"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	Database     = "mongoTest"
	Collection   = "systemEvents_collection"
	MongoDBHosts = "localhost:27017"
)

func getDBCollection() *mgo.Collection {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: Database,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		glog.Info("ERROR: CreateSession: %s\n", err)
	}

	eventsCollection := mongoSession.DB(Database).C(Collection)
	return eventsCollection
}

func findAllAndLogToFile(eventsCollection *mgo.Collection) {
	glog.Info("Retrieving all records")
	var result []interface{}
	iter := eventsCollection.Find(nil)
	err := iter.All(&result)
	if err != nil {
		glog.Info("Error: ", err)
	} else {
		glog.Info("Logging events to XML file")
		for event := range result {
			var tempEvent GenericEventType
			bsonBytes, _ := bson.Marshal(result[event])
			bson.Unmarshal(bsonBytes, &tempEvent)
			logEventToXML(tempEvent, result[event])
		}
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
