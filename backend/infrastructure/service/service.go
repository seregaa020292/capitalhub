package service

type Email interface {
	SendConfirmedMail(email, code string) bool
}
