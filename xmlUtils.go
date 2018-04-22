package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/golang/glog"

	"gopkg.in/mgo.v2/bson"
)

const (
	OUTPUT_FILE_PATH = "logFile.xml"
)

var file, fileErr = os.Create(OUTPUT_FILE_PATH)
var fileWriter *bufio.Writer

func initXml() {
	glog.Info("Initializing XML")
	if fileErr != nil {
		glog.Info("Error opening file: ", fileErr)
	} else {
		glog.Info("XML File created!")
	}

	openXMLFile()
}

func openXMLFile() {
	if fileWriter != nil {
		return
	}
	fileWriter = bufio.NewWriter(file)
	fmt.Fprintln(fileWriter, "<?xml version=\"1.0\"?>")
	fmt.Fprintln(fileWriter, "<log>")
	fileWriter.Flush()
}

func performXMLWrite(currentEvent interface{}) {
	output, err := xml.MarshalIndent(currentEvent, "", "    ")
	if err == nil {
		glog.Info("Writing to file: \n", string(output))
		fmt.Fprintln(fileWriter, string(output))
	} else {
		glog.Info("ERROR writing: ", err)
	}
}

func logEventToXML(genericEventType GenericEventType, event interface{}) {
	glog.Info("Starting to log event...")
	switch genericEventType.EventType {
	case ACCOUNT_TRANSACTION_EVENT:
		var currentEvent AccountTransactionEvent
		bsonBytes, _ := bson.Marshal(event)
		bson.Unmarshal(bsonBytes, &currentEvent)

		performXMLWrite(currentEvent)
		break

	case SYSTEM_EVENT:
		var currentEvent SystemEvent
		bsonBytes, _ := bson.Marshal(event)
		bson.Unmarshal(bsonBytes, &currentEvent)
		performXMLWrite(currentEvent)
		break

	case ERROR_EVENT:
		var currentEvent ErrorEvent
		bsonBytes, _ := bson.Marshal(event)
		bson.Unmarshal(bsonBytes, &currentEvent)
		performXMLWrite(currentEvent)
		break

	case QUOTE_SERVER_EVENT:
		// UNCOMMENTING BECUASE CRYPTOKEYS DONT WORK
		// var currentEvent QuoteServerEvent
		// bsonBytes, _ := bson.Marshal(event)
		// bson.Unmarshal(bsonBytes, &currentEvent)
		// performXMLWrite(currentEvent)
		break

	case USER_COMMAND:
		var currentEvent UserCommand
		bsonBytes, _ := bson.Marshal(event)
		bson.Unmarshal(bsonBytes, &currentEvent)
		performXMLWrite(currentEvent)
		break
	}
}

func closeXMLFile() {
	fmt.Fprintln(fileWriter, "</log>")
	fileWriter.Flush()
	file.Close()
}
