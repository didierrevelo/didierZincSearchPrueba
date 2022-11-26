package service

import (
	zincsearch "github.com/didierrevelo/didierZincSearchPrueba/server/pkg/databaseAdap/zincSearch"
)

// zincSearchAdap
type ZincSearchAdap interface {
	DocumentCreator(index string, emails interface{}) (*zincsearch.DocumentCreatorRes, error)
	DocumentFinder(indexName string, body zincsearch.DocumentFinderReq) (*zincsearch.DocumentFinderRes, error)
}
