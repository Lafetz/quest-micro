package grpcserver

import (
	"regexp"
	"strings"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

var rxEmail = regexp.MustCompile(`^[a-zA-Z-0-9\\_\\-\\.]+@[a-zA-Z\\_\\-\\.]+$`)

func validateAddknight(email string, name string) []*errdetails.BadRequest_FieldViolation {
	var validationErrors []*errdetails.BadRequest_FieldViolation
	if strings.TrimSpace(email) == "" {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "email",
				Description: "email is required",
			})

	} else if len(email) > 254 || !rxEmail.MatchString(email) {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "email",
				Description: "email is not valid",
			})
	}
	if strings.TrimSpace(name) == "" {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "name",
				Description: "name is required",
			})

	}

	return validationErrors
}
func validateGet(name string) []*errdetails.BadRequest_FieldViolation {
	var validationErrors []*errdetails.BadRequest_FieldViolation
	if strings.TrimSpace(name) == "" {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "name",
				Description: "name is required",
			})

	}
	return validationErrors
}
