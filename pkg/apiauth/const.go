package apiauth

import (
	"net/http"
	"net/url"
)

type JmsAPI interface {
	MakeRequest(method, endpoint string, body interface{}) (*http.Request, error)
	DoRequest(req *http.Request, result interface{}) error
	SetQuery(req *http.Request, v url.Values) *http.Request
	GetEndpoint() string
}
