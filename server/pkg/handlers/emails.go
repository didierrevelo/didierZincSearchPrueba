package handlers

import (
	"net/http"

	"github.com/didierrevelo/didierZincSearchPrueba/server/pkg/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// MailHandler is the handler for the email
type MailHandler struct {
	emailService EmailService
}

// NewMailHandler creates a new email handler
func NewMailHandler(emailService EmailService) *MailHandler {
	return &MailHandler{
		emailService: emailService,
	}
}

type UserResponse struct {
	Users []string `json:"users"`
}

// GetUsers returns the users
func (h *MailHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	records, err := h.emailService.GetUsers()
	if err != nil {
		NewError(w, r, http.StatusInternalServerError, err)
		return
	}
	
	res := &UserResponse{
		Users: records,
	}

	render.JSON(w, r, res)
}

type EmailResponse struct {
	Emails []domain.Email `json:"emails"`
}

// GetEmails returns the emails
func (h *MailHandler) GetEmails(w http.ResponseWriter, r *http.Request) {
	idUser := chi.URLParam(r, "userID")

	emails, err := h.emailService.ExtractIntoMail(idUser)
		if err != nil {
			NewError(w, r, http.StatusInternalServerError, err)
			return
	}

	res := &EmailResponse{
		Emails: emails,
	}

	render.JSON(w, r, res)

}

type SearchIntoEmail struct{
	Emails []domain.Email `json:"emails"`
}

func (h *MailHandler) SearchIntoEmail(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("q")

	emails, err := h.emailService.SearchIntoEmail(domain.Index, term)
	if err != nil {
		NewError(w, r, http.StatusInternalServerError, err)
		return
	}

	res := &SearchIntoEmail{
		Emails: emails,
	}

	render.JSON(w, r, res)
}

type ErrRes struct {
	Status int `json:"status"`
	Err string `json:"error"`
}

// NewError creates a new error

func NewError(w http.ResponseWriter, r *http.Request, status int, err error) {
	resErr := &ErrRes{
		Status: status,
		Err: err.Error(),
	}

	render.Status(r, status)
	render.JSON(w, r, resErr)
}
