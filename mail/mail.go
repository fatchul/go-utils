package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
	"strings"
)

var auth smtp.Auth

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

// NewRequest is default format for sending emails
func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

// parseTemplate is use for parse from
func (r *Request) parseTemplate(templateData string, data interface{}) error {
	t, err := template.ParseFiles(templateData)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}

	r.body = buf.String()
	return nil
}

// SendingEmail is function to sending email contain html tag with params data
func SendingEmail(to string, subjectEmail string, body string, templateData interface{}) error {
	r := NewRequest([]string{to}, subjectEmail, "")
	err := r.parseTemplate(body, templateData)
	if err != nil {
		log.Printf("parsing template error : %s", err)
		return err
	}

	ok, err := r.SendEmail(to)
	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}
	if ok {
		return nil
	}

	return nil
}

// encodeRFC2047 reformat string with `<>` notation
func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

// SendEmail is process send email with destination email
func (r *Request) SendEmail(to string) (bool, error) {
	from := "xxx@gmail.com"
	host := "smtp.gmail.com"
	port := "587"
	pass := "Y9sa8jPassword@"
	name := "Test"

	sender := mail.Address{name, from}
	header := make(map[string]string)
	header["From"] = sender.String()
	header["To"] = to
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	header["Subject"] = r.subject

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(r.body))

	addr := host + ":" + port

	auth = smtp.PlainAuth("", sender.Address, pass, host)

	err := smtp.SendMail(addr, auth, from, r.to, []byte(message))
	if err != nil {
		return false, err
	}

	return true, nil
}
