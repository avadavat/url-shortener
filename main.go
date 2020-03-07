package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dpgil/url-shortener/handlers"
)

func main() {
	// read config properties
	config, err := ioutil.ReadFile(".config")
	if err != nil {
		log.Fatal("No config file found. Expected file named .config with dynamo endpoint, region, and table name.", err)
	}
	configParams := strings.Split(string(config), "\n")
	if len(configParams) < 3 {
		log.Fatal("Missing config parameters")
	}

	endpoint := configParams[0]
	region := configParams[1]
	tableName := configParams[2]

	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)

	http.HandleFunc("/d/", handlers.Decode(db, tableName))
	http.HandleFunc("/e/", handlers.Encode(db, tableName))
	http.HandleFunc("/r/", handlers.Redirect(db, tableName))

	// todo: is 9090 the right port?
	err = http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
