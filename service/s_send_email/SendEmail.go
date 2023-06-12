package s_send_email

import (
	"toko-buah/model/m_send_email"
	"toko-buah/repository/r_send_email"
)

type EmailService struct {
	emailRepo r_send_email.EmailRepository
}

func NewEmailService(repo r_send_email.EmailRepository) *EmailService {
	return &EmailService{
		emailRepo: repo,
	}
}

func (s *EmailService) SendEmail(to, subject, body string) error {
	email := &m_send_email.Email{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	err := s.emailRepo.SendEmail(email)
	if err != nil {
		return err
	}

	return nil
}
