# Audit Server

### Local Deployment
1. Clone repo
2. Grab all dependencies need. To do so, run the following commands:


		go get "github.com/gin-gonic/gin"

	      go get "gopkg.in/mgo.v2"

	      go get "gopkg.in/mgo.v2/bson"

	      go get "github.com/gorilla/mux"

	      go get "github.com/golang/glog" 

     		

3. Running the server: `go run *.go -logtostderr=true` 

Server will be up on [http://localhost:8082/](http://localhost:8082/)


### Supported Endpoints

The endpoints supported by the Audit Server are:

      GET /
      
      GET /api/log
      
      GET /api/cleardb
      
      POST /api/systemevent
      
      POST /api/usercommand
      
      POST /api/errorevent
      
      POST /api/quoteserver
      
      POST /api/accounttransaction
		
		
### GET `/api/cleardb`

Call this when in doubt. Will clear the collection used by MongoDb (All records are under one collection now).


### `GET /api/log`

Will spit out the XML logfile in the root of the project
