package helper

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func SentEmailConfirmation(fullname string, date string, email string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "felix.fatahillah22@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "ANB Payment Approval")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Hi "+fullname+" Your holiday on"+date+" is approved. For more information, send Whatsapp message to https://wa.me/+6281332122490")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 465, "felix.fatahillah22@gmail.com", "uflsowzfnnsllwjx")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
