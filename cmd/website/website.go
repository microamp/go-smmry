package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const smmryAPIKeyEnv = "SMMRY_API_KEY"

// SmmryClient ...
type SmmryClient struct {
	baseURL string
	apiKey  string
}

// SmmryResult ...
type SmmryResult struct {
	SmAPICharacterCount string `json:"sm_api_character_count"`
	SmAPITitle          string `json:"sm_api_title"`
	SmAPIContent        string `json:"sm_api_content"`
	SmAPILimitation     string `json:"sm_api_limitation"`
}

// NewSmmryClient ...
func NewSmmryClient() (*SmmryClient, error) {
	apiKey := os.Getenv(smmryAPIKeyEnv)
	if apiKey == "" {
		return nil, fmt.Errorf("Invalid env value for %s", smmryAPIKeyEnv)
	}
	return &SmmryClient{
		baseURL: "http://api.smmry.com/",
		apiKey:  apiKey,
	}, nil
}

// SummaryByWebsite ...
func (client *SmmryClient) SummaryByWebsite(sentences int, websiteURL string) (*SmmryResult, error) {
	url := fmt.Sprintf("%s&SM_API_KEY=%s&SM_LENGTH=%d&SM_URL=%s", client.baseURL, client.apiKey, sentences, websiteURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Printf("Error closing body: %+v", err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result SmmryResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {
	client, err := NewSmmryClient()
	if err != nil {
		log.Panicf("Error instantiating SMMRY client: %+v", err)
	}

	summary, err := client.SummaryByWebsite(10, "http://www.bbc.com/news/election-us-2016-37604731")
	if err != nil {
		log.Panicf("Error getting summary by website: %+v", err)
	}

	log.Printf(summary.SmAPICharacterCount)
	log.Printf(summary.SmAPITitle)
	log.Printf(summary.SmAPIContent)
	log.Printf(summary.SmAPILimitation)
}
