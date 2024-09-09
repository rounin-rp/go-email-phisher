package services

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"time"
)

type EmailManager struct {
	smtpHost string
	smtpPort string
	from     string
	password string
	replyTo  string
	auth     smtp.Auth
}

func BuildEmailManager(host string, port string, from string, password string) EmailManager {

	manager := EmailManager{
		smtpHost: host,
		smtpPort: port,
		from:     from,
		password: password,
		replyTo:  from,
	}
	manager.auth = smtp.PlainAuth(
		"",
		from,
		password,
		host,
	)

	return manager
}

func (em EmailManager) SetReplyTo(replyTo string) {
	em.replyTo = replyTo
}

func (em EmailManager) SendMail(to string, subject string, body string) (bool, error) {
	// Set the email headers.
	headers := make(map[string]string)
	headers["From"] = em.from
	headers["To"] = to // Use the first recipient's email for the 'To' header
	headers["Reply-To"] = em.replyTo
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""
	headers["Content-Transfer-Encoding"] = "7bit"
	headers["Date"] = time.Now().Format(time.RFC1123Z)

	// Format headers into a single string.
	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + body

	// Connect to the SMTP server using TLS.
	conn, err := tls.Dial("tcp", em.smtpHost+":"+em.smtpPort, &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         em.smtpHost,
	})
	if err != nil {
		log.Println("Error connecting to SMTP server:", err)
		return false, err
	}
	defer conn.Close()

	// Create a new SMTP client.
	client, err := smtp.NewClient(conn, em.smtpHost)
	if err != nil {
		log.Println("Error creating SMTP client:", err)
		return false, err
	}
	defer client.Quit()

	// Authenticate with the SMTP server.
	if err = client.Auth(em.auth); err != nil {
		log.Println("Error authenticating with SMTP server:", err)
		return false, err
	}

	// Set the sender.
	if err = client.Mail(em.from); err != nil {
		log.Println("Error setting sender email:", err)
		return false, err
	}

	// set the recipient
	if err = client.Rcpt(to); err != nil {
		log.Println("Error setting recipient email:", err)
		return false, err
	}

	// Send the email data.
	wc, err := client.Data()
	if err != nil {
		log.Println("Error sending email data:", err)
		return false, err
	}
	// Write the headers and message content.
	_, err = wc.Write([]byte(message))
	if err != nil {
		log.Println("Error writing message content:", err)
		return false, nil
	}

	// Close the data writer.
	err = wc.Close()
	if err != nil {
		log.Println("Error closing message writer:", err)
		return false, nil
	}

	// Successfully sent.
	log.Println("Email sent successfully to multiple recipients!")

	return true, nil
}
