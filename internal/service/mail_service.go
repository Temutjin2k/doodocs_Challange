package service

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
)

type MailServiceImpl interface {
	SendFile(mails []string, filename, mimeType string, filedata []byte) error
}

type mailService struct {
	smtpHost string
	smtpPort string
	email    string
	password string
	smptAddr string
	auth     smtp.Auth
}

func NewMailService(smtpHost, smtpPort, email, password string) (*mailService, error) {
	auth := smtp.PlainAuth("", email, password, smtpHost)

	// test if given parametrs correct
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{"kayixa2211@cironex.com"}, []byte("Test request"))
	if err != nil {
		return nil, fmt.Errorf("couldn't create mail service. Error: %v", err)
	}

	return &mailService{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		email:    email,
		password: password,
		smptAddr: smtpHost + ":" + smtpPort,
		auth:     auth,
	}, nil
}

func (s *mailService) SendFile(mails []string, filename, mimeType string, filedata []byte) error {
	subject := "Bing!"
	var msg bytes.Buffer
	msg.WriteString(fmt.Sprintf("From: %s\r\n", s.email))
	msg.WriteString("Subject: " + subject + "\r\n")
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: multipart/mixed; boundary=\"boundary\"\r\n")
	msg.WriteString("\r\n--boundary\r\n")
	msg.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	msg.WriteString("Content-Transfer-Encoding: 7bit\r\n")
	msg.WriteString("\r\n")
	msg.WriteString("Hello! Someone send u this file.\r\n")
	msg.WriteString("\r\n--boundary\r\n")
	msg.WriteString(fmt.Sprintf("Content-Type: %s; name=\"%s\"\r\n", mimeType, filename))
	msg.WriteString("Content-Transfer-Encoding: base64\r\n")
	msg.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", filename))
	msg.WriteString("\r\n")

	// Encode file content as base64
	encodedFile := base64.StdEncoding.EncodeToString(filedata)
	chunkSize := 76
	for i := 0; i < len(encodedFile); i += chunkSize {
		end := i + chunkSize
		if end > len(encodedFile) {
			end = len(encodedFile)
		}
		msg.WriteString(encodedFile[i:end] + "\r\n")
	}
	msg.WriteString("--boundary--\r\n")

	return smtp.SendMail(s.smptAddr, s.auth, s.email, mails, msg.Bytes())
}
