package utils

import (
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/pkg/httpErrors"
)

// JWT Claims struct
type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

// Создание токенов
func GenerateTokens(user *models.User, config *config.Config) (Tokens, error) {
	accessToken, err := GenerateAccessToken(user, config)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: GenerateRefreshToken(),
	}, nil
}

// Создание Access токена
func GenerateAccessToken(user *models.User, config *config.Config) (string, error) {
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Email: user.Email,
		ID:    user.UserID.String(),
		Role:  *user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * config.Auth.AccessTokenExpMinute).Unix(),
		},
	})

	return token.SignedString([]byte(config.Auth.AccessSecretKey))
}

// Создание Refresh токена
func GenerateRefreshToken() string {
	return uuid.New().String()
}

// Извлечь JWT из запроса
func ExtractJWTFromRequest(echoCtx echo.Context, config *config.Config) (map[string]interface{}, error) {
	// Инициализировать новый экземпляр `Claims` (здесь используется карта утверждений)
	claims := jwt.MapClaims{}
	token, err := ParseJWT(echoCtx, config, claims)

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ParseJWT(echoCtx echo.Context, config *config.Config, claims jwt.MapClaims) (*jwt.Token, error) {
	// Получить строку JWT
	tokenString, err := ExtractBearerToken(echoCtx)
	if err != nil {
		return nil, err
	}

	// Анализируем строку JWT и сохраняем результат в `claims`.
	// Обратите внимание, что мы также передаем ключ в этом методе. Этот метод вернет ошибку
	// если токен недействителен (если он истек в соответствии со сроком действия, который мы установили при входе в систему),
	// или если подпись не совпадает
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (jwtKey interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Auth.AccessSecretKey), nil
	})

	if err != nil {
		if ev, ok := err.(*jwt.ValidationError); ok {
			if ev.Errors == jwt.ValidationErrorExpired {
				return token, httpErrors.ExpiredJWTToken
			}
		}
	}

	return token, err
}

// Извлекаем токен из заголовка авторизации запроса
func ExtractBearerToken(echoCtx echo.Context) (string, error) {
	bearerHeader := echoCtx.Request().Header.Get("Authorization")
	if bearerHeader == "" {
		bearerHeader = echoCtx.QueryParam("authorization")
	}
	if bearerHeader == "" {
		return "", httpErrors.EmptyJWTToken
	}

	headerParts := strings.Split(bearerHeader, " ")
	if len(headerParts) != 2 {
		return "", errors.New("len(headerParts) != 2")
	}

	return html.EscapeString(headerParts[1]), nil
}

func ParseUserIDFromJWT(id interface{}) (uuid.UUID, error) {
	userID, ok := id.(string)
	if !ok {
		return [16]byte{}, httpErrors.InvalidJWTClaims
	}
	return uuid.Parse(userID)
}
