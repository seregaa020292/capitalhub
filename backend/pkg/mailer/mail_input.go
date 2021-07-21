package mailer

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/seregaa020292/capitalhub/config"
	"html/template"
	"strings"
)

type EmailInput struct {
	from string
	to   []string
	msg  []byte
	body []byte
}

type Sender interface {
	Send(input *EmailInput) error
}

func NewEmailInput(from string, to []string, subject string, msg []byte) *EmailInput {
	body := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"From: %s\r\n"+
			"Subject: %s\r\n"+
			"Content-Type: text/html; charset=\\\"UTF-8\\\";\r\n\r\n"+
			"%s\r\n",
		strings.Join(to, ","), from, subject, msg),
	)

	return &EmailInput{
		from: from,
		to:   to,
		msg:  msg,
		body: body,
	}
}

func (emailInput *EmailInput) Validate() error {
	if len(emailInput.to) == 0 || len(emailInput.body) == 0 {
		return errors.New("empty to/body")
	}

	for _, email := range emailInput.to {
		if !IsEmailValid(email) {
			return errors.New("invalid from email")
		}
	}

	return nil
}

func Html(config config.EmailConfig, data interface{}) ([]byte, error) {
	t, err := template.ParseFiles(config.ConfirmedPartial, config.BaseLayout)
	if err != nil {
		return []byte(""), err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return []byte(""), err
	}

	return buf.Bytes(), nil
}
