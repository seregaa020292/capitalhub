package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User full model
type User struct {
	UserID    uuid.UUID `json:"userId" db:"user_id" redis:"user_id" validate:"omitempty"`
	Name      string    `json:"name" db:"name" redis:"name" validate:"required,lte=30"`
	Email     string    `json:"email,omitempty" db:"email" redis:"email" validate:"omitempty,lte=60,email"`
	Password  string    `json:"password,omitempty" db:"password" redis:"password" validate:"omitempty,required,gte=6"`
	Role      *string   `json:"role,omitempty" db:"role" redis:"role" validate:"omitempty,lte=10"`
	Avatar    *string   `json:"avatar,omitempty" db:"avatar" redis:"avatar" validate:"omitempty,lte=512,url"`
	Confirmed *string   `json:"-" db:"confirmed"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" db:"updated_at" redis:"updated_at"`
}

// Хеширование пароля пользователя с помощью bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Сравнить пароль пользователя с переданным паролем
func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

// Проверка на подверждение почты
func (u *User) IsConfirmed() bool {
	return u.Confirmed == nil
}

// Очистить пароль пользователя
func (u *User) SanitizePassword() {
	u.Password = ""
}

// Готовим пользователя к созданию
func (u *User) PrepareCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}

	if u.Role != nil {
		*u.Role = strings.ToLower(strings.TrimSpace(*u.Role))
	}
	return nil
}

// Готовим пользователя к обновлению
func (u *User) PrepareUpdate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	if u.Role != nil {
		*u.Role = strings.ToLower(strings.TrimSpace(*u.Role))
	}
	return nil
}

// Формат ответа, все пользователи
type UsersList struct {
	TotalCount int     `json:"totalCount"`
	TotalPages int     `json:"totalPages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"hasMore"`
	Users      []*User `json:"users"`
}

type AccessToken struct {
	Token       string `json:"token"`
	PrefixToken string `json:"prefixToken"`
}

type RefreshToken struct {
	Token string `json:"token"`
}

type Tokens struct {
	AccessToken  *AccessToken  `json:"accessToken"`
	RefreshToken *RefreshToken `json:"refreshToken"`
}

type UserBase struct {
	User *User `json:"user"`
}

type UserWithToken struct {
	User         *User         `json:"user"`
	AccessToken  *AccessToken  `json:"accessToken"`
	RefreshToken *RefreshToken `json:"refreshToken"`
}
