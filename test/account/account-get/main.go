package main

import (
	"log"
	"os"

	api "github.com/danikarik/clicksend-go/client"
	auth "github.com/go-openapi/runtime/client"
)

var (
	username = os.Getenv("CLICKSEND_USERNAME")
	apiKey   = os.Getenv("CLICKSEND_API_KEY")
)

func main() {
	// basic auth credentials
	basicAuth := auth.BasicAuth(username, apiKey)
	// default api client
	client := api.Default
	// call api
	resp, err := client.Account.AccountGet(nil, basicAuth)
	if err != nil {
		log.Fatalf("could not get response: %v", err)
	}
	// work with payload
	log.Printf("%v", resp.Payload)
}
