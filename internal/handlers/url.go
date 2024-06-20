package handlers

import "fmt"

var (
	EmployeesURL       = "/api/v1/employees"
	EmployeeIDURLParam = "employee_id"
	EmployeeURL        = fmt.Sprintf("%s/{%s}", EmployeesURL, EmployeeIDURLParam)
	SubscribeURL       = "/subscribe"
	UnsubscribeURL     = "/unsubscribe"
	SignInURL          = "/api/v1/sign-in"
	SignUpURL          = "/api/v1/sign-up"
	RefreshURL         = "/api/v1/refresh"
	SubscriptionsURL   = "/api/v1/subscriptions"
	LogOutURL          = "/api/v1/log-out"
)
