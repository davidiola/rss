package main

import (
	"encoding/json"
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/sfreiberg/gotwilio"
	"net/http"
	"strconv"
)

func main() {

	message := ""

	addSportsMessage(&message)
	addBTCMessage(&message)
	addETHMessage(&message)

	fmt.Println(message)

	accountSID, authToken := "insert_SID_here", "insert_token_here"
	twilio := gotwilio.NewTwilioClient(accountSID, authToken)

	from := "+12706814675"
	to := "+12148429453"
	twilio.SendSMS(from, to, message, "", "")

}

func addSportsMessage(message *string) {

	*message = *message + "SPORTS\n"
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://www.espn.com/espn/rss/news")
	for i, val := range feed.Items {
		if i == 10 {
			break
		}
		s := strconv.Itoa(i + 1)
		*message = *message + s + ". " + val.Description + "\n"
	}
}

type BTC struct {
	Price string `json:"high"`
}

type ETH struct {
	Price float64 `json:"USD"`
}

func addBTCMessage(message *string) {

	*message = *message + "CRYPTOCURRENCIES\n"
	BTCurl := "https://www.bitstamp.net/api/v2/ticker/btcusd/"
	req, _ := http.NewRequest("GET", BTCurl, nil)
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	var btcprice BTC
	json.NewDecoder(resp.Body).Decode(&btcprice)
	*message = *message + "1 Bitcoin = $" + btcprice.Price + "\n"

}

func addETHMessage(message *string) {

	ETHurl := "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=BTC,USD,EUR"
	req, _ := http.NewRequest("GET", ETHurl, nil)
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	var ethprice ETH
	json.NewDecoder(resp.Body).Decode(&ethprice)
	eth := strconv.FormatFloat(ethprice.Price, 'f', -1, 64)
	*message = *message + "1 Ethereum = $" + eth

}
