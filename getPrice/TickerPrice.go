package getPrice

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

/*
API_KEY='D7...Ejj',
API_SECRET='gwQ...u3A'

print('ticker/price', bot.tickerPrice(
symbol='BNBBTC'
))
*/

type Response struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

var responce []Response

func formHandler(requestType string, pair string) string {

	baseURL := "https://api.binance.com/api/v3/"

	performUrl := baseURL + requestType + "?symbol=" + pair
	//fmt.Println(performUrl)
	return performUrl
}

func TickerPrice(pair string) string {

	//url := formHandler("ticker/price", pair)
	//resp := http.Client{Timeout: 2 * time.Second}
	//if _, err := resp.Get(url); err != nil {
	//	log.Fatal(err)
	//}

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	url := formHandler("ticker/price", pair)
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	var myVal Response
	err = json.Unmarshal(body, &myVal)
	return fmt.Sprintln(myVal.Price)
}
