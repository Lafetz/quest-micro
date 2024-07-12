package grpcserver

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

var rxEmail = regexp.MustCompile(`^[a-zA-Z-0-9\\_\\-\\.]+@[a-zA-Z\\_\\-\\.]+$`)

func validateAddknight(email string, username string, password string) []*errdetails.BadRequest_FieldViolation {
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
	if strings.TrimSpace(username) == "" {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "username",
				Description: "username is required",
			})

	}
	if utf8.RuneCountInString(password) < 8 {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "password",
				Description: "password too short",
			})
	}
	return validationErrors
}
func validateGet(username string) []*errdetails.BadRequest_FieldViolation {
	var validationErrors []*errdetails.BadRequest_FieldViolation
	if strings.TrimSpace(username) == "" {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "username",
				Description: "username is required",
			})

	}
	return validationErrors
}
