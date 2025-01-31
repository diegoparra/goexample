package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

////////VERSAO A SER REFATORADA
// func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
// 	jsonBytes, err := json.Marshal(body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
// 	if err != nil {
// 		return nil, err
// 	}
// 	request.Header = headers
// 	client := &http.Client{}
// 	return client.Do(request)
// }

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

// Post sends a post request to the URL with the body
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return Client.Do(request)
}

func main() {
	log.Println("jeff")
}
