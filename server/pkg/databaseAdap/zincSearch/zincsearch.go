package zincsearch

import (
	"fmt"
	"net/http"
	"os"

	"github.com/didierrevelo/didierZincSearchPrueba/server/pkg/databaseAdap"
)

const (
	// DefaultHost is the default host for the server
	DefaultHost = "http://localhost:4080"
)

// ZincSearchClient
type ZincSearchClient struct {
	adapter *databaseadap.DatabaseAdap
}

// NewZincSearchClient creates a new ZincSearchClient
func NewZincSearchClient(c *http.Client) *ZincSearchClient {
	host := os.Getenv("ZINC_HOST")
	if host == "" {
		host = DefaultHost
	}

	a := databaseadap.NewDatabaseAdap(c, host)
	setHeaderAdap(a)
	return &ZincSearchClient{
		adapter: a,
	}
}

// setHeaderAdap sets the header for the ZincSearchClient
func setHeaderAdap(adapter *databaseadap.DatabaseAdap){
	username := os.Getenv("ZINC_FIRST_ADMIN_USER")
	password := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")

	if username == "" || password == "" {
		panic("ZINC_FIRST_ADMIN_USER and ZINC_FIRST_ADMIN_PASSWORD environment variables were not set")
	}

	adapter.SetAuth(username, password) 
}

// DocumentCreator creates a document
func (a *ZincSearchClient) DocumentCreator(indexName string, records interface{}) (*DocumentCreatorRes, error) {
	response := &DocumentCreatorRes{}
	errApiRes := &ErrorRes{}

	path := "/api/_bulkv2"
	body := DocumentcreatorReq{
		Index:   indexName,
		Records: records,
	}

	req, err := a.adapter.SetBodyReq(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	res, err := a.adapter.Sling.Do(req, response, errApiRes)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(errApiRes.ErrorMessage)
	}

	return response, nil
}

// DocumentFinder searches documents with zincSearch
func (a *ZincSearchClient) DocumentFinder(indexName string, body DocumentFinderReq) (*DocumentFinderRes, error) {
	response := &DocumentFinderRes{}
	errApiRes := &ErrorRes{}

	path := fmt.Sprintf("/api/%s/_search", indexName)

	req, err := a.adapter.SetBodyReq(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	res, err := a.adapter.Sling.Do(req, response, errApiRes)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(errApiRes.ErrorMessage)
	}

	return response, nil
}
