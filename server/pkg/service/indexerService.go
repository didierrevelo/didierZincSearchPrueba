package service

import (
	"log"

	"github.com/didierrevelo/didierZincSearchPrueba/server/pkg/domain"
)

// IndexerService is the interface for the IndexerService
type IndexerService struct {
	zincsearchAdap ZincSearchAdap
}

// NewIndexerService creates a new IndexerService
func NewIndexerService(zsa ZincSearchAdap) *IndexerService {
	return &IndexerService{
		zincsearchAdap: zsa,
	}
}

// IndexEmails indexes emails with the ZincSearch API
func (is *IndexerService) IndexEmails(indexName string, records []domain.Email) error {
	res, err := is.zincsearchAdap.DocumentCreator(indexName, records)
	if err != nil {
		return err
	}

	log.Printf("Indexed %d documents\n", res.RecordCount)

	return nil
}