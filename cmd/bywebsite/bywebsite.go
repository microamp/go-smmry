package main

/*
Example:
go run cmd/bywebsite/bywebsite.go https://en.wikipedia.org/wiki/Richard_Stallman 5
*/

import (
	"log"
	"os"

	"github.com/microamp/go-smmry/smmry"
)

func main() {
	args := os.Args[1:]
	websiteURL, length := args[0], args[1]

	client, err := smmry.NewSmmryClient()
	if err != nil {
		log.Panicf("Error instantiating SMMRY client: %+v", err)
	}

	summary, err := client.SummaryByWebsite(websiteURL, length)
	if err != nil {
		log.Panicf("Error getting summary by website: %+v", err)
	}

	log.Printf("Character count: %s", summary.SmAPICharacterCount)
	log.Printf("Title: %s", summary.SmAPITitle)
	log.Printf("Summary: %s", summary.SmAPIContent)
	log.Printf("Quota info: %s", summary.SmAPILimitation)
}
