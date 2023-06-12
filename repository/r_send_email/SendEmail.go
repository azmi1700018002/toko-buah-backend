package r_send_email

import (
	"net/smtp"
	"toko-buah/model/m_send_email"
)

type EmailRepository interface {
	SendEmail(email *m_send_email.Email) error
}

type GmailRepository struct {
	SMTPServer   string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

func NewGmailRepository(server, port, username, password string) *GmailRepository {
	return &GmailRepository{
		SMTPServer:   server,
		SMTPPort:     port,
		SMTPUsername: username,
		SMTPPassword: password,
	}
}

func (r *GmailRepository) SendEmail(email *m_send_email.Email) error {
	auth := smtp.PlainAuth("", r.SMTPUsername, r.SMTPPassword, r.SMTPServer)
	msg := "From: " + r.SMTPUsername + "\n" +
		"To: " + email.To + "\n" +
		"Subject: " + email.Subject + "\n\n" +
		email.Body

	err := smtp.SendMail(r.SMTPServer+":"+r.SMTPPort, auth, r.SMTPUsername, []string{email.To}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
