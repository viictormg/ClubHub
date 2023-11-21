package adapters

import "net/http"

type IHTTPPortClient interface {
	Do(req *http.Request) (*http.Response, error)
}
