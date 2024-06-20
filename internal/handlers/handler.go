package handlers

import (
	"net/http"
)

type Handler interface {
	Register(router http.Handler) error
	Employee() http.HandlerFunc
	Employees() http.HandlerFunc
	LogOut() http.HandlerFunc
	Refresh() http.HandlerFunc
	SignIn() http.HandlerFunc
	SignUp() http.HandlerFunc
	Subscribe() http.HandlerFunc
	Unsubscribe() http.HandlerFunc
	Subscriptions() http.HandlerFunc
}
