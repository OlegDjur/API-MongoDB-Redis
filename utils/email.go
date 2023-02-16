package utils

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"

	"github.com/OlegDjur/API-MongoDB-Redis/configs"
	"github.com/OlegDjur/API-MongoDB-Redis/models"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func SendEmail(user *models.DBResponse, data *EmailData, temp *template.Template, templateName string) error {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("cloud not load config", err)
	}

	// Sender data
	from := config.EmailFrom
	smtpPass := config.SMTPPass
	smtpUser := config.SMTPUser
	to := user.Email
	smtpHost := config.SMTPHost
	smtpPort := config.SMTPPort

	var body bytes.Buffer

	if err := temp.ExecuteTemplate(&body, templateName, &data); err != nil {
		log.Fatal("Cloud not execute template", err)
	}

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
