package c_send_email

import (
	"encoding/json"
	"net/http"
	"toko-buah/service/s_send_email"
)

type EmailController struct {
	emailService *s_send_email.EmailService
}

func NewEmailController(service *s_send_email.EmailService) *EmailController {
	return &EmailController{
		emailService: service,
	}
}

func (c *EmailController) SendEmail(w http.ResponseWriter, r *http.Request) {
	var request struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.emailService.SendEmail(request.To, request.Subject, request.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}
