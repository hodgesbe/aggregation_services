package financials

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hodgesbe/aggregation_services/common"
)

var tingoApi = common.API{
	Key:       os.Getenv("TINGOAPI_KEY"),
	BaseURL:   "https://api.tiingo.com/",
	AuthParam: "token=",
}

// GetIEXticker takes a given ticker (stock symbol) and returns the latest top-of-book price
func GetIEXticker(ticker string) {
	resp, err := tingoApi.Request("iex/" + ticker + "?")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(bodyBytes))
}

// GetHistIntraday returns the historical intraday prices for a stock (date fmt: yyyy-mm-dd)
func GetHistIntraday(ticker string, startDate string, endDate string, freq string) {
	//https://api.tiingo.com/iex/<ticker>/prices?startDate=2019-01-02&resampleFreq=5min
	resp, err := tingoApi.Request("iex/" + ticker + "/prices?startDate=" +
		startDate + "&endDate=" + endDate + "&resampleFreq=" + freq + "&")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(bodyBytes))
}
