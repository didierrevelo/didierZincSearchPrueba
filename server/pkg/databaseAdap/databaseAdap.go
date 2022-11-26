package databaseadap

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dghubble/sling"
)

const (
	headerType  = "Content-Type"
	contentType = "application/json"
)

var errorMeth = errors.New("wrong method")

// DatabaseAdap is the structure that wraps the basic methods of a database adapter
type DatabaseAdap struct {
	// httpClient is the http client that will be used to make the requests
	httpClient *http.Client
	Sling      *sling.Sling
}

// NewDatabaseAdap creates a new DatabaseAdap
func NewDatabaseAdap(httpClient *http.Client, url string) *DatabaseAdap {
	s := sling.New().Client(httpClient).Base(url)

	return &DatabaseAdap{
		httpClient: httpClient,
		Sling:      s,
	}
}

// HeaderAdap sets a header for the DatabaseAdap
func (a *DatabaseAdap) HeaderAdap(header, value string) {
	a.Sling.Set(header, value)
}

// SetAuth sets the basic auth DatabaseAdap
func (a *DatabaseAdap) SetAuth(username, password string) {
	a.Sling.SetBasicAuth(username, password)
}

// SetBodyReq sets the body of the DatabaseAdap
func (a *DatabaseAdap) SetBodyReq(method string, path string, body interface{}) (*http.Request, error) {
	value, err := a.GetReq(method, path)
	if err != nil {
		return &http.Request{}, err
	}

	if err := SetBody(value, body); err != nil {
		return &http.Request{}, err
	}

	request, err := value.Request()
	if err != nil {
		return &http.Request{}, err
	}

	return request, nil

}

// GetReq gets the request of the DatabaseAdap
func (a *DatabaseAdap) GetReq(method string, path string) (*sling.Sling, error) {
	switch method {
	case http.MethodGet:
		return a.Sling.New().Get(path), nil
	case http.MethodPost:
		return a.Sling.New().Post(path), nil
	case http.MethodPut:
		return a.Sling.New().Put(path), nil
	case http.MethodDelete:
		return a.Sling.New().Delete(path), nil
	case http.MethodPatch:
		return a.Sling.New().Patch(path), nil
	default:
		return a.Sling, errorMeth
	}
}

// SetBody sets the body of the DatabaseAdap
func SetBody(sling *sling.Sling, body interface{}) error {
	if body == nil {
		return nil
	}
	
	jsonBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	sling.Set(headerType, contentType).Body(bytes.NewReader(jsonBody))
	
	return nil
}
