package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/didierrevelo/didierZincSearchPrueba/server/pkg/databaseAdap/zincSearch"
	"github.com/didierrevelo/didierZincSearchPrueba/server/pkg/domain"
)

const (
	emailDetailSeparator         = "\r\n"
	emailDetailsContentSeparator = "\r\n\r\n"
	defaultEmailSearchType       = "matchphrase"
	defaultEmailMaxResults       = 5
)

// EmailService is the service for the email
type EmailService struct {
	zincSearchAdap ZincSearchAdap
}

// NewEmailService creates a new email service
func NewEmailService(zincSearchAdap ZincSearchAdap) *EmailService {
	return &EmailService{
		zincSearchAdap: zincSearchAdap,
	}
}

// GetUsers returns the users
func (s *EmailService) GetUsers() ([]string, error) {
	var users []string
	files, err := os.ReadDir(domain.FileFolderRoot)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			users = append(users, file.Name())
		}
	}

	return users, nil
}

// EmailFileDetails
func (s *EmailService) EmailFileDetails(filePathS string) (*domain.Email, error) {
	file, err := os.ReadFile(filepath.Clean(filePathS))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Split the file into the email details and the email content
	fileSplit := strings.SplitN(string(file), emailDetailsContentSeparator, 2)
	if len(fileSplit) != 2 {
		return nil, fmt.Errorf("invalid email file format")
	}

	allContent, content := fileSplit[0], fileSplit[1]

	contentArr := strings.Split(allContent, emailDetailSeparator)

	// Get the email details
	emailDetails := emailMap(contentArr)
	emailDetails.Content = content
	emailDetails.Filepath = filePathS

	return emailDetails, nil
}

// emailMap
func emailMap(emailDetails []string) *domain.Email {
	email := &domain.Email{}

	for i := 0; i < len(emailDetails); i++ {
		keyValue := strings.SplitN(emailDetails[i], ": ", 2)
		switch keyValue[0] {
		case "Message-ID":
			email.MessageID = keyValue[1]
		case "Date":
			email.Date = keyValue[1]
		case "From":
			email.From = keyValue[1]
		case "To":
			email.To = keyValue[1]
		case "Subject":
			email.Subject = keyValue[1]
		default:
			continue
		}
	}

	return email
}

// ExtractIntoMail returns the emails
func (s *EmailService) ExtractIntoMail(idUser string) ([]domain.Email, error) {
	var emails []domain.Email

	pathFolder := fmt.Sprintf("%s/%s", domain.FileFolderRoot, idUser)
	// Get the user folder

	err := filepath.Walk(pathFolder, s.processFile(&emails))
	if err != nil {
		log.Println(err)
	}

	return emails, nil

}

// processFile
func (s *EmailService) processFile(emails *[]domain.Email) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
		}
		if !info.IsDir() {
			email, err := s.EmailFileDetails(path)
			if err != nil || email == nil {
				return err
			}
			*emails = append(*emails, *email)
		}
		return nil
	}
}

// SearchIntoEmail returns the emails
func (s *EmailService) SearchIntoEmail(indexName string, term string) ([]domain.Email, error) {
	now := time.Now()
	timeStart := now.AddDate(0, 0, -7).Format("2022-01-02T15:04:05Z")
	timeEnd := now.Format("2022-01-02T15:04:05Z")

	// Search the emails
	body := zincsearch.DocumentFinderReq {
		SearchType: defaultEmailSearchType,
		Query: zincsearch.SearchDocumentsRequestQuery {
			Term:      term,
			StartTime: timeStart,
			EndTime:   timeEnd,
		},
		SortFields: []string{"-@timestamp"},
		From:       0,
		MaxResults: defaultEmailMaxResults,
	}

	res, err := s.zincSearchAdap.DocumentFinder(indexName, body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return processSearchResult(res), nil
}

// processSearchResult
func processSearchResult(res *zincsearch.DocumentFinderRes) []domain.Email {
	var emails []domain.Email

	for _, hit := range res.Hits.Hits {
		var email domain.Email

		emailBytes, _ := json.Marshal(hit.Source)

		err := json.Unmarshal(emailBytes, &email)
		if err != nil {
			log.Println("Error in mapZincSearchResponseToEmails: ", err)
			continue
		}

		emails = append(emails, email)
	}

	return emails

}
