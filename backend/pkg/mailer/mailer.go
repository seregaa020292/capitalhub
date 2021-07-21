package mailer

import (
	"errors"
	"fmt"
	"net/smtp"

	"github.com/seregaa020292/capitalhub/config"
)

type mailer struct {
	cfg config.MailerConfig
}

func NewClient(cfg config.MailerConfig) *mailer {
	return &mailer{cfg: cfg}
}

func (mailer *mailer) Send(input *EmailInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	var auth smtp.Auth
	addr := fmt.Sprintf("%s:%d", mailer.cfg.Host, mailer.cfg.Port)

	switch mailer.cfg.Mechanism {
	case "CRAMMD5":
		auth = smtp.CRAMMD5Auth(mailer.cfg.User, mailer.cfg.Password)
	case "PLAIN":
		auth = smtp.PlainAuth("", mailer.cfg.User, mailer.cfg.Password, mailer.cfg.Host)
	default:
		return errors.New("Error - invalid authentication mechanism")
	}

	return smtp.SendMail(addr, auth, mailer.cfg.FromEmail, input.to, input.body)
}
