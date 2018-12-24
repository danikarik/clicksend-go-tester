package main

import (
	"log"
	"os"

	"github.com/danikarik/clicksend-go/client"
	apiclient "github.com/go-openapi/runtime/client"
)

var (
	username = os.Getenv("CLICKSEND_USERNAME")
	apiKey   = os.Getenv("CLICKSEND_API_KEY")
)

func main() {
	auth := apiclient.BasicAuth(username, apiKey)
	api := client.Default
	resp, err := api.Account.AccountGet(nil, auth)
	if err != nil {
		log.Fatalf("could not get response: %v", err)
	}
	log.Printf("%v", resp.Payload)
}
