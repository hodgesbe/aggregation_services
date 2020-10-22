package common

import (
	"net/http"
)

type API struct {
	Key       string
	BaseURL   string
	AuthParam string
}

func (api *API) Request(paramString string) (http.Response, error) {
	resp, err := http.Get(api.BaseURL + paramString + api.AuthParam + api.Key)

	return *resp, err
}
