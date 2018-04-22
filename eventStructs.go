package main

type AccountTransactionEvent struct {
	Timestamp      int64  `bson:"timestamp" json:"timestamp"`
	Server         string `bson:"server" json:"server"`
	TransactionNum int    `bson:"transactionnum" json:"transactionnum"`
	Action         string `bson:"action" json:"action"`
	UserId         string `bson:"userid" json:"userid"`
	Funds          string `bson:"funds" json:"funds"`
	EventType      string `bson:"eventtype" json:"eventtype"`
}

type SystemEvent struct {
	Timestamp      int64  `bson:"timestamp" json:"timestamp"`
	Server         string `bson:"server" json:"server"`
	TransactionNum int    `bson:"transactionnum" json:"transactionnum"`
	Command        string `bson:"command" json:"command"`
	UserId         string `bson:"userid" json:"userid"`
	StockSymbol    string `bson:"stocksymbol" json:"stocksymbol"`
	Funds          string `bson:"funds" json:"funds"`
	EventType      string `bson:"eventtype" json:"eventtype"`
}

type QuoteServerEvent struct {
	Timestamp            int64  `bson:"timestamp" json:"timestamp"`
	Server               string `bson:"server" json:"server"`
	TransactionNum       int    `bson:"transactionnum" json:"transactionnum"`
	QuoteServerEventTime int64  `bson:"quoteservereventtime" json:"quoteservereventtime"`
	UserId               string `bson:"userid" json:"userid"`
	StockSymbol          string `bson:"stocksymbol" json:"stocksymbol"`
	Price                string `bson:"price" json:"price"`
	Cryptokey            string `bson:"cryptokey" json:"cryptokey"`
	EventType            string `bson:"eventtype" json:"eventtype"`
}

type ErrorEvent struct {
	Timestamp      int64  `bson:"timestamp" json:"timestamp"`
	Server         string `bson:"server" json:"server"`
	TransactionNum int    `bson:"transactionnum" json:"transactionnum"`
	Command        string `bson:"command" json:"command"`
	UserId         string `bson:"userid" json:"userid"`
	StockSymbol    string `bson:"stocksymbol" json:"stocksymbol"`
	Funds          string `bson:"funds" json:"funds"`
	ErrorMessage   string `bson:"errormessage" json:"errormessage"`
	EventType      string `bson:"eventtype" json:"eventtype"`
}

type UserCommand struct {
	Timestamp      int64  `bson:"timestamp" json:"timestamp"`
	Server         string `bson:"server" json:"server"`
	TransactionNum int    `bson:"transactionnum" json:"transactionnum"`
	Command        string `bson:"command" json:"command"`
	UserId         string `bson:"userid" json:"userid"`
	StockSymbol    string `bson:"stocksymbol" json:"stocksymbol"`
	Funds          string `bson:"funds" json:"funds"`
	EventType      string `bson:"eventtype" json:"eventtype"`
}

type GenericEventType struct {
	EventType string `bson:"eventtype" json:"eventtype"`
}
