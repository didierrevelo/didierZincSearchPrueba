package handlers

import (
	"github.com/didierrevelo/didierZincSearchPrueba/server/pkg/domain"
)

// EmailService is the interface for the email service
type EmailService interface {
	GetUsers() ([]string, error)
	ExtractIntoMail(userID  string) ([]domain.Email, error)
	SearchIntoEmail(indexName string, term string) ([]domain.Email, error)
}

// IndexerService is the interface for the indexer service
type IndexerService interface {
	IndexEmails(indexName string, emails []domain.Email) error
}
