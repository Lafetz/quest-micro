package heraldserver

import (
	"context"
	"errors"

	mailv1 "github.com/lafetz/quest-micro/proto/gen/mail/v1"
	herald "github.com/lafetz/quest-micro/services/herald/core"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrValidate = errors.New("there was a problem with the provided data")
)

func (h *HeraldServer) SendEmail(ctx context.Context, req *mailv1.SendEmailRequest) (*mailv1.SendEmailResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req ")
	}
	validationErrors := validateSendEmailRequest(req.From, req.To, req.Subject, req.Body)
	if len(validationErrors) > 0 {
		stat := status.New(codes.InvalidArgument, ErrValidate.Error())
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return nil, s.Err()
	}
	err := h.HeraldService.SendQuestAssignmentEmail(herald.Email{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	})
	if err != nil {
		h.logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &mailv1.SendEmailResponse{}, nil
}
