package middleware

import (
	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/infrastructure/session"
	"github.com/seregaa020292/capitalhub/internal/auth"
	"github.com/seregaa020292/capitalhub/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	sessUC  session.UseCase
	authUC  auth.UseCase
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(
	sessUC session.UseCase,
	authUC auth.UseCase,
	cfg *config.Config,
	origins []string,
	logger logger.Logger,
) *MiddlewareManager {
	return &MiddlewareManager{
		sessUC:  sessUC,
		authUC:  authUC,
		cfg:     cfg,
		origins: origins,
		logger:  logger,
	}
}
