package main

import "encoding/xml"

type AccountTransactionEvent struct {
	XMLName        xml.Name `xml:"accountTransaction"`
	Timestamp      int64    `bson:"timestamp" json:"timestamp" xml:"timestamp"`
	Server         string   `bson:"server" json:"server" xml:"server"`
	TransactionNum int      `bson:"transactionnum" json:"transactionnum" xml:"transactionNum"`
	Action         string   `bson:"action" json:"action" xml:"action"`
	UserId         string   `bson:"userid" json:"userid" xml:"username"`
	Funds          string   `bson:"funds" json:"funds" xml:"funds,omitempty"`
	EventType      string   `bson:"eventtype" json:"eventtype" xml:"-"`
}

type SystemEvent struct {
	XMLName        xml.Name `xml:"systemEvent"`
	Timestamp      int64    `bson:"timestamp" json:"timestamp" xml:"timestamp"`
	Server         string   `bson:"server" json:"server" xml:"server"`
	TransactionNum int      `bson:"transactionnum" json:"transactionnum" xml:"transactionNum"`
	Command        string   `bson:"command" json:"command" xml:"command,omitempty"`
	UserId         string   `bson:"userid" json:"userid" xml:"username"`
	StockSymbol    string   `bson:"stocksymbol" json:"stocksymbol" xml:"stockSymbol,omitempty"`
	Funds          string   `bson:"funds" json:"funds" xml:"funds,omitempty"`
	EventType      string   `bson:"eventtype" json:"eventtype" xml:"-"`
}

type QuoteServerEvent struct {
	XMLName              xml.Name `xml:"quoteServer"`
	Timestamp            int64    `bson:"timestamp" json:"timestamp" xml:"timestamp"`
	Server               string   `bson:"server" json:"server" xml:"server"`
	TransactionNum       int      `bson:"transactionnum" json:"transactionnum" xml:"transactionNum"`
	QuoteServerEventTime int64    `bson:"quoteservereventtime" json:"quoteservereventtime" xml:"quoteServerTime"`
	UserId               string   `bson:"userid" json:"userid" xml:"username"`
	StockSymbol          string   `bson:"stocksymbol" json:"stocksymbol" xml:"stockSymbol,omitempty"`
	Price                string   `bson:"price" json:"price" xml:"price"`
	Cryptokey            string   `bson:"cryptokey" json:"cryptokey" xml:"cryptokey"`
	EventType            string   `bson:"eventtype" json:"eventtype" xml:"-"`
}

type ErrorEvent struct {
	XMLName        xml.Name `xml:"errorEvent"`
	Timestamp      int64    `bson:"timestamp" json:"timestamp" xml:"timestamp"`
	Server         string   `bson:"server" json:"server" xml:"server"`
	TransactionNum int      `bson:"transactionnum" json:"transactionnum" xml:"transactionNum"`
	Command        string   `bson:"command" json:"command" xml:"command,omitempty"`
	UserId         string   `bson:"userid" json:"userid" xml:"username"`
	StockSymbol    string   `bson:"stocksymbol" json:"stocksymbol" xml:"stockSymbol,omitempty"`
	Funds          string   `bson:"funds" json:"funds" xml:"funds,omitempty"`
	ErrorMessage   string   `bson:"errormessage" json:"errormessage" xml:"errorMessage"`
	EventType      string   `bson:"eventtype" json:"eventtype" xml:"-"`
}

type UserCommand struct {
	XMLName        xml.Name `xml:"userCommand"`
	Timestamp      int64    `bson:"timestamp" json:"timestamp" xml:"timestamp"`
	Server         string   `bson:"server" json:"server" xml:"server"`
	TransactionNum int      `bson:"transactionnum" json:"transactionnum" xml:"transactionNum"`
	Command        string   `bson:"command" json:"command" xml:"command,omitempty"`
	UserId         string   `bson:"userid" json:"userid" xml:"username,omitempty"`
	StockSymbol    string   `bson:"stocksymbol" json:"stocksymbol" xml:"stockSymbol,omitempty"`
	Funds          string   `bson:"funds" json:"funds" xml:"funds,omitempty"`
	EventType      string   `bson:"eventtype" json:"eventtype" xml:"-"`
}

type GenericEventType struct {
	EventType string `bson:"eventtype" json:"eventtype" xml:"-"`
}
