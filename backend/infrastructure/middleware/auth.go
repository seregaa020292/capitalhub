package middleware

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/auth"
	"github.com/seregaa020292/capitalhub/internal/auth/model"
	"github.com/seregaa020292/capitalhub/pkg/httpErrors"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// JWT способ аутентификации с использованием cookie или заголовка авторизации
func (mw *MiddlewareManager) AuthJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if err := mw.validateJWTToken(mw.authUC, ctx, mw.cfg); err != nil {
			mw.logger.Error("middleware validateJWTToken", zap.String("headerJWT", err.Error()))
			return ctx.JSON(http.StatusUnauthorized, httpErrors.NewUnauthorizedError(httpErrors.Unauthorized))
		}

		return next(ctx)
	}
}

// Роль администратора
func (mw *MiddlewareManager) AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*model.User)
		if !ok || *user.Role != "admin" {
			return c.JSON(http.StatusForbidden, httpErrors.NewUnauthorizedError(httpErrors.PermissionDenied))
		}
		return next(c)
	}
}

// Middleware для аутентификации на основе роли admin
func (mw *MiddlewareManager) OwnerOrAdminMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*model.User)
			if !ok {
				mw.logger.Errorf("Error c.Get(user) RequestID: %s, ERROR: %s,", utils.GetRequestID(c), "invalid user ctx")
				return c.JSON(http.StatusUnauthorized, httpErrors.NewUnauthorizedError(httpErrors.Unauthorized))
			}

			if *user.Role == "admin" {
				return next(c)
			}

			if user.UserID.String() != c.Param("user_id") {
				mw.logger.Errorf("Error c.Get(user) RequestID: %s, UserID: %s, ERROR: %s,",
					utils.GetRequestID(c),
					user.UserID.String(),
					"invalid user ctx",
				)
				return c.JSON(http.StatusForbidden, httpErrors.NewForbiddenError(httpErrors.Forbidden))
			}

			return next(c)
		}
	}
}

// Middleware для аутентификации на основе ролей пользователя
func (mw *MiddlewareManager) RoleBasedAuthMiddleware(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*model.User)
			if !ok {
				mw.logger.Errorf("Error c.Get(user) RequestID: %s, UserID: %s, ERROR: %s,",
					utils.GetRequestID(c),
					user.UserID.String(),
					"invalid user ctx",
				)
				return c.JSON(http.StatusUnauthorized, httpErrors.NewUnauthorizedError(httpErrors.Unauthorized))
			}

			for _, role := range roles {
				if role == *user.Role {
					return next(c)
				}
			}

			mw.logger.Errorf("Error c.Get(user) RequestID: %s, UserID: %s, ERROR: %s,",
				utils.GetRequestID(c),
				user.UserID.String(),
				"invalid user ctx",
			)

			return c.JSON(http.StatusForbidden, httpErrors.NewForbiddenError(httpErrors.PermissionDenied))
		}
	}
}

func (mw *MiddlewareManager) validateJWTToken(authUC auth.UseCase, ctx echo.Context, cfg *config.Config) error {
	claims, err := utils.ExtractJWTFromRequest(ctx, cfg)
	if err != nil {
		return err
	}

	userUUID, err := utils.ParseUserIDFromJWT(claims["id"])
	if err != nil {
		return err
	}

	user, err := authUC.GetByID(ctx.Request().Context(), userUUID)
	if err != nil {
		return err
	}

	token, _ := utils.ExtractBearerToken(ctx)

	ctx.Set("user", user)
	ctx.Set("sid", token)
	ctxUser := context.WithValue(ctx.Request().Context(), utils.UserCtxKey{}, user)
	ctx.SetRequest(ctx.Request().WithContext(ctxUser))

	mw.logger.Info(
		"SessionMiddleware, RequestID: %s,  IP: %s, UserID: %s, AccessToken: %s",
		utils.GetRequestID(ctx),
		utils.GetIPAddress(ctx),
		user.UserID.String(),
		token,
	)

	return nil
}
