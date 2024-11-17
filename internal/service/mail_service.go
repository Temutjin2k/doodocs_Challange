package service

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/go-mail/mail/v2"
)

type MailServiceImpl interface {
	SendFile(mails []string, filename, mimeType string, filedata []byte) error
}

type mailService struct {
	smtpHost string
	smtpPort int
	email    string
	password string
	smptAddr string
	dialer   *mail.Dialer
}

func NewMailService(smtpHost, smtpPort, email, password string) (*mailService, error) {
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return nil, errors.New("Invalid port. Error: " + err.Error())
	}

	d := mail.NewDialer(smtpHost, port, email, password)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	return &mailService{
		smtpHost: smtpHost,
		smtpPort: port,
		email:    email,
		password: password,
		smptAddr: smtpHost + ":" + smtpPort,
		dialer:   d,
	}, nil
}

func (s *mailService) SendFile(mails []string, filename, mimeType string, filedata []byte) error {
	m := mail.NewMessage()

	m.SetHeader("From", s.email)

	m.SetHeader("To", mails...)
	m.SetHeader("Subject", "Bing!")
	m.SetBody("text/plain", "Hello! Someone sent you this file.")

	// Attaching file
	m.Attach(filename,
		mail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(filedata) // Write the file data directly to the writer
			return err
		}),
		mail.SetHeader(map[string][]string{
			"Content-Type":        {fmt.Sprintf("%s; name=\"%s\"", mimeType, filename)},
			"Content-Disposition": {fmt.Sprintf("attachment; filename=\"%s\"", filename)},
		}),
	)

	// Send the email
	if err := s.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
