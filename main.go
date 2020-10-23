package main

import (
	"github.com/hodgesbe/aggregation_services/financials"
)

func main() {
	//financials.GetIEXticker("aapl")
	financials.GetHistIntraday("snap", "2017-01-01", "2020-10-26", "2hour")
}
