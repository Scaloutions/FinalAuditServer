package main

import (
	"time"

	"github.com/golang/glog"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	Database     = "mongoTest"
	Collection   = "sysEvent_collection_100"
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

func findAll(eventsCollection *mgo.Collection) {
	glog.Info("Now retrieving all records")
	var result []interface{}
	iter := eventsCollection.Find(nil)
	err := iter.All(&result)
	if err != nil {
		glog.Info("ERRRPR", err)
	} else {
		for event := range result {
			// glog.Info("Event:", event, reflect.TypeOf(result[event]))
			var tempEvent GenericEventType
			bsonBytes, _ := bson.Marshal(result[event])
			bson.Unmarshal(bsonBytes, &tempEvent)
			glog.Info(tempEvent)

			logEventToXML(tempEvent, result[event])

		}
	}
}

func insertRecord(eventsCollection *mgo.Collection, logMsg interface{}) {
	if err := eventsCollection.Insert(logMsg); err != nil {
		glog.Info("Error in insertion : ", err)
	} else {
		glog.Info("Success! Inserted event!")
	}
}
