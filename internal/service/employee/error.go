package employeesrv

import "errors"

var (
	ErrInvalidLoginOrPassword = errors.New("invalid login or password")
	ErrTokenExpired           = errors.New("token is expired. authorize again")
	ErrSelfSubscription       = errors.New("self subscription error")
	ErrInternal               = errors.New("internal error")
	ErrInvalidEmployeeID      = errors.New("invalid employee id")
)
