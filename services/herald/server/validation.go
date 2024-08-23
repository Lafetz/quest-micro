package heraldserver

import (
	"regexp"
	"strings"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

var rxEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func validateSendEmailRequest(from, to, subject, body string) []*errdetails.BadRequest_FieldViolation {
	var validationErrors []*errdetails.BadRequest_FieldViolation

	if strings.TrimSpace(from) == "" {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "from",
				Description: "sender email is required",
			})
	} else if len(from) > 254 || !rxEmail.MatchString(from) {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "from",
				Description: "sender email is not valid",
			})
	}

	if strings.TrimSpace(to) == "" {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "to",
				Description: "recipient email is required",
			})
	} else if len(to) > 254 || !rxEmail.MatchString(to) {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "to",
				Description: "recipient email is not valid",
			})
	}

	if strings.TrimSpace(subject) == "" {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "subject",
				Description: "subject is required",
			})
	}

	if strings.TrimSpace(body) == "" {
		validationErrors = append(validationErrors,
			&errdetails.BadRequest_FieldViolation{
				Field:       "body",
				Description: "body is required",
			})
	}

	return validationErrors
}
