package financials

import (
	"log"
	"net/http"
)

type API struct {
	Key       string
	BaseURL   string
	AuthParam string
}

func (api *API) Request(paramString string) http.Response {
	resp, err := http.Get(api.BaseURL + paramString + api.AuthParam + api.Key)
	if err != nil {
		log.Fatalln(err)
	}

	return *resp
}
