package client

import "C"
import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"ks-aggregator-rates/internal/pkg/client/requests"
	"log"
	"net/http"
)

type Client struct {
	Request requests.ClientRequest
}

func (c Client) Run(db *gorm.DB, timestamp int64) error {
	request := c.Request.ParseRequest()
	// fmt.Printf("Info: %s\n\t- Request: %s\n", c.Request.RequestInfo(), request)
	resp, err := http.Get(request)
	if err != nil {
		log.Fatalln(err)
	}
	var body interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return err
	}
	price := c.Request.ParseResponse(body, timestamp)
	// fmt.Println("- ", price)
	// fmt.Println("price dbRecord", priceDbRecord)
	result := db.Create(price.ToPrice())
	fmt.Printf("Info: %s\n\t- result: %s\n", c.Request.RequestInfo(), result)

	return nil
}
