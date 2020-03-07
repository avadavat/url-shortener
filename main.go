package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dpgil/url-shortener/handlers"
)

func main() {
	// config, err := ioutil.ReadFile(".config")
	// if err != nil {
	// 	log.Fatal("No config file found", err)
	// }

	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://127.0.0.1:8000"),
		//EndPoint: aws.String("https://dynamodb.us-east-1.amazonaws.com"),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)

	http.HandleFunc("/d/", handlers.Decode(db))
	http.HandleFunc("/e/", handlers.Encode(db))
	http.HandleFunc("/r/", handlers.Redirect(db))

	// todo: is 9090 the right port?
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
