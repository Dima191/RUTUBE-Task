package smtpmanager

type Manager interface {
	SendMail(to string, msg string) error
}
