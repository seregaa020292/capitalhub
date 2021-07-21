package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/infrastructure/service"
	mockSess "github.com/seregaa020292/capitalhub/infrastructure/session/mock"
	"github.com/seregaa020292/capitalhub/internal/auth/mock"
	"github.com/seregaa020292/capitalhub/internal/auth/model"
	"github.com/seregaa020292/capitalhub/pkg/converter"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/mailer"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

func TestAuthHandlers_Register(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUC := mock.NewMockUseCase(ctrl)
	mockSessUC := mockSess.NewMockUseCase(ctrl)

	cfg := &config.Config{
		Session: config.SessionConfig{
			Expire: 10,
		},
		Logger: config.LoggerConfig{
			Development: true,
		},
		Email: config.EmailConfig{
			ConfirmedPartial: "./templates/mails/account_confirmed.html",
		},
		Mailer: config.MailerConfig{
			FromEmail: "test@mail.ru",
		},
		Server: config.ServerConfig{
			FrontendUrl: "",
		},
	}

	apiLogger := logger.NewApiLogger(cfg)
	emailSender := mailer.NewMailer(cfg)
	emailService := service.NewEmailService(cfg, emailSender, apiLogger)

	authHandlers := NewAuthHandlers(cfg, mockAuthUC, mockSessUC, emailService, apiLogger)

	user := &model.User{
		Name:     "Name",
		Email:    "email@gmail.com",
		Password: "123456",
	}

	buf, err := converter.AnyToBytesBuffer(user)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", strings.NewReader(buf.String()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	ctx := utils.GetRequestCtx(c)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "auth.Register")
	defer span.Finish()

	handlerFunc := authHandlers.Register()

	userUID := uuid.New()
	userWithToken := &model.UserWithToken{
		User: &model.User{
			UserID: userUID,
		},
	}
	sess := &model.Session{
		UserID: userUID,
	}
	session := "session"

	mockAuthUC.EXPECT().Register(ctxWithTrace, gomock.Eq(user)).Return(userWithToken, nil)
	mockSessUC.EXPECT().CreateSession(ctxWithTrace, gomock.Eq(sess), 10).Return(session, nil)

	err = handlerFunc(c)
	require.NoError(t, err)
	require.Nil(t, err)
}

func TestAuthHandlers_Login(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUC := mock.NewMockUseCase(ctrl)
	mockSessUC := mockSess.NewMockUseCase(ctrl)

	cfg := &config.Config{
		Session: config.SessionConfig{
			Expire: 10,
		},
		Logger: config.LoggerConfig{
			Development: true,
		},
		Email: config.EmailConfig{
			ConfirmedPartial: "./templates/mails/account_confirmed.html",
		},
		Mailer: config.MailerConfig{
			FromEmail: "test@mail.ru",
		},
		Server: config.ServerConfig{
			FrontendUrl: "",
		},
	}

	apiLogger := logger.NewApiLogger(cfg)
	emailSender := mailer.NewMailer(cfg)
	emailService := service.NewEmailService(cfg, emailSender, apiLogger)

	authHandlers := NewAuthHandlers(cfg, mockAuthUC, mockSessUC, emailService, apiLogger)

	type Login struct {
		Email    string `json:"email" db:"email" validate:"omitempty,lte=60,email"`
		Password string `json:"password,omitempty" db:"password" validate:"required,gte=6"`
	}

	login := &Login{
		Email:    "email@mail.com",
		Password: "123456",
	}

	user := &model.User{
		Email:    login.Email,
		Password: login.Password,
	}

	buf, err := converter.AnyToBytesBuffer(user)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(buf.String()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	ctx := utils.GetRequestCtx(c)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "auth.Login")
	defer span.Finish()

	handlerFunc := authHandlers.Login()

	userUID := uuid.New()
	userWithToken := &model.UserWithToken{
		User: &model.User{
			UserID: userUID,
		},
	}
	sess := &model.Session{
		UserID: userUID,
	}
	session := "session"

	mockAuthUC.EXPECT().Login(ctxWithTrace, gomock.Eq(user)).Return(userWithToken, nil)
	mockSessUC.EXPECT().CreateSession(ctxWithTrace, gomock.Eq(sess), 10).Return(session, nil)

	err = handlerFunc(c)
	require.NoError(t, err)
	require.Nil(t, err)
}

func TestAuthHandlers_Logout(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUC := mock.NewMockUseCase(ctrl)
	mockSessUC := mockSess.NewMockUseCase(ctrl)

	cfg := &config.Config{
		Session: config.SessionConfig{
			Expire: 10,
		},
		Logger: config.LoggerConfig{
			Development: true,
		},
	}

	apiLogger := logger.NewApiLogger(cfg)
	emailSender := mailer.NewMailer(cfg)
	emailService := service.NewEmailService(cfg, emailSender, apiLogger)

	authHandlers := NewAuthHandlers(cfg, mockAuthUC, mockSessUC, emailService, apiLogger)

	sessionKey := "session-id"
	cookieValue := "cookieValue"

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/logout", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.AddCookie(&http.Cookie{Name: sessionKey, Value: cookieValue})

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	ctx := utils.GetRequestCtx(c)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "auth.Logout")
	defer span.Finish()

	userID := uuid.New()

	logout := authHandlers.Logout()

	cookie, err := req.Cookie(sessionKey)
	require.NoError(t, err)
	require.NotNil(t, cookie)
	require.NotEqual(t, cookie.Value, "")
	require.Equal(t, cookie.Value, cookieValue)

	mockSessUC.EXPECT().DeleteByID(ctxWithTrace, gomock.Eq(userID), gomock.Eq(cookie.Value)).Return(nil)

	err = logout(c)
	require.NoError(t, err)
	require.Nil(t, err)
}
