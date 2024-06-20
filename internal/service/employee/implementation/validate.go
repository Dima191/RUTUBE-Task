package employeesrvimpl

import (
	"errors"
	"fmt"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/go-playground/validator"
	"github.com/hashicorp/go-multierror"
	"log/slog"
)

func (s *service) validateErrorHandler(err error) error {
	if err != nil {
		var validateErr validator.ValidationErrors
		if errors.As(err, &validateErr) {
			var validationErrors error
			for _, v := range validateErr {
				s.logger.Error("employee validation failed", slog.String("field", v.StructNamespace()), slog.String("value", v.Value().(string)))
				validationErrors = multierror.Append(validationErrors, fmt.Errorf(`value "%s" invalid`, v.Value().(string)))
			}
			return validationErrors
		}
		s.logger.Error("validation error", slog.String("error", err.Error()))
		return srv.ErrInternal
	}
	return nil
}
