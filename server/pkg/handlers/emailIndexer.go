package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/didierrevelo/didierZincSearchPrueba/server/pkg/domain"
	"github.com/go-chi/render"
)

// Indexer is the handler for the indexer
type Indexer struct {
	indexerService IndexerService
	emailService   EmailService
}

// NewIndexer creates a new indexer handler
func NewIndexer(indexerService IndexerService, emailService EmailService) *Indexer {
	return &Indexer{
		indexerService: indexerService,
		emailService:   emailService,
	}
}

// IndexEmails indexes the emails
func (h *Indexer) IndexEmails(w http.ResponseWriter, r *http.Request) {
	emailsBYId, err := h.emailService.GetUsers()
	if err != nil {
		NewError(w, r, http.StatusInternalServerError, err)
		return
	}

	var wg sync.WaitGroup
	for _, userID := range emailsBYId {
		wg.Add(1)
		go h.mailByUserID(userID, &wg)
	}

	wg.Wait()

	render.Status(r, http.StatusNoContent)
}

func (h *Indexer) mailByUserID(emailUserId string, wg *sync.WaitGroup) {
	defer wg.Done()
	emails, err := h.emailService.ExtractIntoMail(emailUserId)
	if err != nil {
		log.Println(err)
		return
	}
	if err = h.indexerService.IndexEmails(domain.Index, emails); err != nil {
		log.Println(err)
	}
}
