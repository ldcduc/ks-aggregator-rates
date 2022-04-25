package client

import "C"
import (
	"encoding/json"
	"ks-aggregator-rates/internal/pkg/client/requests"
	"log"
	"net/http"
)

type Client struct {
	Request requests.ClientRequest
}

func (c Client) Run() error {
	request := c.Request.ParseRequest()
	// fmt.Printf("Info: %s\n\t- Request: %s\n", c.Request.RequestInfo(), request)
	resp, err := http.Get(request)
	if err != nil {
		log.Fatalln(err)
	}
	var body interface{}
	// fmt.Println("resp.Body", resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return err
	}
	// fmt.Printf("Info: %s\n\t- Response %s\n", c.Request.RequestInfo(), body)

	return nil
}
