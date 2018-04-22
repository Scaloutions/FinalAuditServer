package main

import (
	"strconv"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"

	"github.com/beevik/etree"
)

const (
	OUTPUT_FILE_PATH = "logFile.xml"
)

var doc = etree.NewDocument()

func initXml() {
	doc.CreateProcInst("xml", `version="1.0"`)

	// people := doc.CreateElement("People")
	// people.CreateComment("These are all known people")

	// jon := people.CreateElement("Person")
	// jon.CreateAttr("name", "Jon")

	// sally := people.CreateElement("Person")
	// sally.CreateAttr("name", "Sally")

	// doc.Indent(2)
	// doc.WriteToFile(OUTPUT_FILE_PATH)
}

func logEventToXML(genericEventType GenericEventType, event interface{}) {
	switch genericEventType.EventType {
	case ACCOUNT_TRANSACTION_EVENT:
		var currentEvent AccountTransactionEvent
		bsonBytes, _ := bson.Marshal(event)
		bson.Unmarshal(bsonBytes, &currentEvent)

		glog.Info("It is: ", currentEvent)
		accountTransactionTag := doc.CreateElement("accountTransaction")
		accountTransactionTag.CreateElement("timestamp").CreateCharData(strconv.FormatInt(currentEvent.Timestamp, 10))

		break
	case SYSTEM_EVENT:
		var currentEvent SystemEvent
		bsonBytes, _ := bson.Marshal(event)
		bson.Unmarshal(bsonBytes, &currentEvent)

		glog.Info("THIS IS A SYS EVTN", currentEvent)
		break
	}

	doc.Indent(2)
	doc.WriteToFile(OUTPUT_FILE_PATH)
}
