package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dpgil/url-shortener/handlers"
)

const defaultPort = "5000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	endpoint := os.Getenv("DYNAMO_ENDPOINT")
	region := os.Getenv("DYNAMO_REGION")
	tableName := os.Getenv("DYNAMO_TABLE_NAME")

	if endpoint == "" || region == "" || tableName == "" {
		// No config environment variables set. Check if a .config file exists.
		config, err := ioutil.ReadFile(".config")
		if err != nil {
			log.Fatal("No config file found. Expected file named .config with dynamo endpoint, region, and table name.", err)
		}
		configParams := strings.Split(string(config), "\n")
		if len(configParams) < 3 {
			log.Fatal("Missing config parameters")
		}

		endpoint = configParams[0]
		region = configParams[1]
		tableName = configParams[2]
	}

	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)

	http.HandleFunc("/d", handlers.Decode(db, tableName))
	http.HandleFunc("/e", handlers.Encode(db, tableName))

	fmt.Println("Running on port: " + port)
	err := http.ListenAndServe(":"+port, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
