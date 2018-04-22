package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

const (
	ACCOUNT_TRANSACTION_EVENT = "accountTransactionEvent"
	SYSTEM_EVENT              = "systemEvent"
	QUOTE_SERVER_EVENT        = "quoteServeEvent"
	ERROR_EVENT               = "errorEvent"
	USER_COMMAND              = "userCommand"
	LOG_EVENT                 = "log"
	LOG_BY_USER_EVENT         = "logByUser"
	TRANSACTION_HISTORY       = "transactionHistory"
)

func usage() {
	fmt.Println("usage: example -logtostderr=true -stderrthreshold=[INFO|WARN|FATAL|ERROR] -log_dir=[string]\n")
	flag.PrintDefaults()
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to the Audit Server!")
}

func parseEvent(w http.ResponseWriter, r *http.Request, eventsCollection *mgo.Collection, eventType string) {
	switch eventType {
	case SYSTEM_EVENT:
		eventMsg := SystemEvent{
			EventType: SYSTEM_EVENT,
		}
		err := json.NewDecoder(r.Body).Decode(&eventMsg)
		if err != nil {
			glog.Error("Error decoding!", err)
		}
		insertRecord(eventsCollection, eventMsg)
		break

	case ERROR_EVENT:
		eventMsg := ErrorEvent{
			EventType: ERROR_EVENT,
		}
		err := json.NewDecoder(r.Body).Decode(&eventMsg)
		if err != nil {
			glog.Error("Error decoding!", err)
		}
		insertRecord(eventsCollection, eventMsg)
		break

	case ACCOUNT_TRANSACTION_EVENT:
		eventMsg := AccountTransactionEvent{
			EventType: ACCOUNT_TRANSACTION_EVENT,
		}
		err := json.NewDecoder(r.Body).Decode(&eventMsg)
		if err != nil {
			glog.Error("Error decoding!", err)
		}
		insertRecord(eventsCollection, eventMsg)
		break

	case QUOTE_SERVER_EVENT:
		eventMsg := QuoteServerEvent{
			EventType: QUOTE_SERVER_EVENT,
		}
		err := json.NewDecoder(r.Body).Decode(&eventMsg)
		if err != nil {
			glog.Error("Error decoding!", err)
		}
		insertRecord(eventsCollection, eventMsg)
		break

	case USER_COMMAND:
		eventMsg := UserCommand{
			EventType: USER_COMMAND,
		}
		err := json.NewDecoder(r.Body).Decode(&eventMsg)
		if err != nil {
			glog.Error("Error decoding!", err)
		}
		insertRecord(eventsCollection, eventMsg)
		break

	case LOG_EVENT:
		findAllAndLogToFile(eventsCollection)
		break
	}

}

func clearDb(eventsCollection *mgo.Collection) {
	if err := eventsCollection.DropCollection(); err != nil {
		glog.Info("Error dropping collection! : ", err)
	} else {
		glog.Info("Successfully dropped collection!")
	}
}

func main() {
	router := mux.NewRouter()
	flag.Usage = usage
	flag.Parse()
	glog.Info("Spinning up the server..")

	var eventsCollection = getDBCollection()

	if eventsCollection == nil {
		glog.Info("Cound not connect! Returned nill!")
	}

	glog.Info("Connecting to mongoDb")
	clearDb(eventsCollection)
	initXml()

	router.HandleFunc("/", homeFunc).Methods("GET")

	router.HandleFunc("/api/systemevent", func(w http.ResponseWriter, r *http.Request) {
		parseEvent(w, r, eventsCollection, SYSTEM_EVENT)
	}).Methods("POST")

	router.HandleFunc("/api/usercommand", func(w http.ResponseWriter, r *http.Request) {
		parseEvent(w, r, eventsCollection, USER_COMMAND)
	}).Methods("POST")

	router.HandleFunc("/api/quoteserver", func(w http.ResponseWriter, r *http.Request) {
		parseEvent(w, r, eventsCollection, QUOTE_SERVER_EVENT)
	}).Methods("POST")

	router.HandleFunc("/api/errorevent", func(w http.ResponseWriter, r *http.Request) {
		parseEvent(w, r, eventsCollection, ERROR_EVENT)
	}).Methods("POST")

	router.HandleFunc("/api/accounttransaction", func(w http.ResponseWriter, r *http.Request) {
		parseEvent(w, r, eventsCollection, ACCOUNT_TRANSACTION_EVENT)
	}).Methods("POST")

	router.HandleFunc("/api/log", func(w http.ResponseWriter, r *http.Request) {
		parseEvent(w, r, eventsCollection, LOG_EVENT)
	}).Methods("GET")

	router.HandleFunc("/api/cleardb", func(w http.ResponseWriter, r *http.Request) {
		clearDb(eventsCollection)
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8082", router))
}
