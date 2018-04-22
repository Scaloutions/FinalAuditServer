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
var logTag *etree.Element = nil

func initXml() {
	glog.Info("XML File created")
	doc.CreateProcInst("xml", `version="1.0"`)
	logTag = doc.CreateElement("log")
}

func addTimeStampTag(parentTag *etree.Element, timestamp int64) {
	parentTag.
		CreateElement("timestamp").
		CreateCharData(strconv.FormatInt(timestamp, 10))
}

func addTransactionNumTag(parentTag *etree.Element, transactionNum int) {
	parentTag.
		CreateElement("transactionNum").
		CreateCharData(strconv.Itoa(transactionNum))
}

func addServerTag(parentTag *etree.Element, server string) {
	parentTag.
		CreateElement("server").
		CreateCharData(server)
}

func addUsernameTag(parentTag *etree.Element, username string) {
	parentTag.
		CreateElement("username").
		CreateCharData(username)
}

func addActionTag(parentTag *etree.Element, action string) {
	parentTag.
		CreateElement("action").
		CreateCharData(action)
}

func addFundsTag(parentTag *etree.Element, funds string) {
	parentTag.
		CreateElement("funds").
		CreateCharData(funds)
}

func addCommandTag(parentTag *etree.Element, command string) {
	parentTag.
		CreateElement("command").
		CreateCharData(command)
}

func addStockSymbolTag(parentTag *etree.Element, stockSymbol string) {
	parentTag.
		CreateElement("stockSymbol").
		CreateCharData(stockSymbol)
}

func addErrorMessageTag(parentTag *etree.Element, errorMessage string) {
	parentTag.
		CreateElement("errorMessage").
		CreateCharData(errorMessage)
}

func addQuoteServerTimeTag(parentTag *etree.Element, quoteServerTime int64) {
	parentTag.
		CreateElement("quoteServerTime").
		CreateCharData(strconv.FormatInt(quoteServerTime, 10))
}

func addPriceTag(parentTag *etree.Element, price string) {
	parentTag.
		CreateElement("price").
		CreateCharData(price)
}

func addCryptokeyTag(parentTag *etree.Element, cryptokey string) {
	parentTag.
		CreateElement("cryptokey").
		CreateCharData(cryptokey)
}

func logEventToXML(genericEventType GenericEventType, event interface{}) {
	switch genericEventType.EventType {
	case ACCOUNT_TRANSACTION_EVENT:
		var currentEvent AccountTransactionEvent
		bsonBytes, _ := bson.Marshal(event)
		bson.Unmarshal(bsonBytes, &currentEvent)

		accountTransactionTag := logTag.CreateElement("accountTransaction")
		addTimeStampTag(accountTransactionTag, currentEvent.Timestamp)
		addServerTag(accountTransactionTag, currentEvent.Server)
		addTransactionNumTag(accountTransactionTag, currentEvent.TransactionNum)
		addUsernameTag(accountTransactionTag, currentEvent.UserId)
		addActionTag(accountTransactionTag, currentEvent.Action)
		addFundsTag(accountTransactionTag, currentEvent.Funds)
		break

	case SYSTEM_EVENT:
		var currentEvent SystemEvent
		bsonBytes, _ := bson.Marshal(event)
		bson.Unmarshal(bsonBytes, &currentEvent)

		systemEventTag := logTag.CreateElement("systemEvent")
		addTimeStampTag(systemEventTag, currentEvent.Timestamp)
		addServerTag(systemEventTag, currentEvent.Server)
		addTransactionNumTag(systemEventTag, currentEvent.TransactionNum)
		addCommandTag(systemEventTag, currentEvent.Command)
		addUsernameTag(systemEventTag, currentEvent.UserId)
		addStockSymbolTag(systemEventTag, currentEvent.StockSymbol)
		addFundsTag(systemEventTag, currentEvent.Funds)
		break

	case ERROR_EVENT:
		var currentEvent ErrorEvent
		bsonBytes, _ := bson.Marshal(event)
		bson.Unmarshal(bsonBytes, &currentEvent)

		errorEventTag := logTag.CreateElement("errorEvent")
		addTimeStampTag(errorEventTag, currentEvent.Timestamp)
		addServerTag(errorEventTag, currentEvent.Server)
		addTransactionNumTag(errorEventTag, currentEvent.TransactionNum)
		addCommandTag(errorEventTag, currentEvent.Command)
		addUsernameTag(errorEventTag, currentEvent.UserId)
		addFundsTag(errorEventTag, currentEvent.Funds)
		addErrorMessageTag(errorEventTag, currentEvent.ErrorMessage)
		break

	case QUOTE_SERVER_EVENT:
		var currentEvent QuoteServerEvent
		bsonBytes, _ := bson.Marshal(event)
		bson.Unmarshal(bsonBytes, &currentEvent)

		quoteServerEventTag := logTag.CreateElement("quoteServer")
		addTimeStampTag(quoteServerEventTag, currentEvent.Timestamp)
		addServerTag(quoteServerEventTag, currentEvent.Server)
		addTransactionNumTag(quoteServerEventTag, currentEvent.TransactionNum)
		addQuoteServerTimeTag(quoteServerEventTag, currentEvent.QuoteServerEventTime)
		addUsernameTag(quoteServerEventTag, currentEvent.UserId)
		addStockSymbolTag(quoteServerEventTag, currentEvent.StockSymbol)
		addPriceTag(quoteServerEventTag, currentEvent.Price)
		addCryptokeyTag(quoteServerEventTag, currentEvent.Cryptokey)
		break

	case USER_COMMAND:
		var currentEvent UserCommand
		bsonBytes, _ := bson.Marshal(event)
		bson.Unmarshal(bsonBytes, &currentEvent)

		userCommandEventTag := logTag.CreateElement("userCommand")
		addTimeStampTag(userCommandEventTag, currentEvent.Timestamp)
		addServerTag(userCommandEventTag, currentEvent.Server)
		addTransactionNumTag(userCommandEventTag, currentEvent.TransactionNum)
		addCommandTag(userCommandEventTag, currentEvent.Command)
		addUsernameTag(userCommandEventTag, currentEvent.UserId)
		addStockSymbolTag(userCommandEventTag, currentEvent.StockSymbol)
		addFundsTag(userCommandEventTag, currentEvent.Funds)
		break
	}

	doc.Indent(4)
	doc.WriteToFile(OUTPUT_FILE_PATH)
}
