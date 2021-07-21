package service

import (
	"fmt"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/mailer"
)

const confirmedLinkTmpl = "%s/confirmed/%s" // <frontend_url>/confirmed/<confirmed_code>

type EmailService struct {
	cfg    *config.Config
	client mailer.Sender
	logger logger.Logger
}

type confirmedEmailInput struct {
	VerificationLink string
}

func NewEmailService(cfg *config.Config, client mailer.Sender, logger logger.Logger) *EmailService {
	return &EmailService{
		cfg:    cfg,
		client: client,
		logger: logger,
	}
}

func (service EmailService) SendConfirmedMail(email, code string) bool {
	templateInput := confirmedEmailInput{service.createVerificationLink(code)}
	sendInput, err := mailer.Html(service.cfg.Email, templateInput)

	if err != nil {
		service.logger.Error(err)
		return false
	}

	input := mailer.NewEmailInput(service.cfg.Mailer.FromEmail, []string{email}, "Регистрация", sendInput)

	if err := service.client.Send(input); err != nil {
		service.logger.Error(err)
		return false
	}

	return true
}

func (service EmailService) createVerificationLink(code string) string {
	return fmt.Sprintf(confirmedLinkTmpl, service.cfg.Server.FrontendUrl, code)
}
