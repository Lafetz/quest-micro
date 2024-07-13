package grpcserver

import (
	"context"
	"errors"
	"runtime/debug"

	commonerrors "github.com/lafetz/quest-demo/common/errors"
	protoknight "github.com/lafetz/quest-demo/proto/knight"
	knight "github.com/lafetz/quest-demo/services/knight/core"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GrpcServer) AddKnight(ctx context.Context, req *protoknight.AddKnightReq) (*protoknight.AddKnightRes, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req ")
	}
	validationErrors := validateAddknight(req.Email, req.Username, req.Password)
	if len(validationErrors) > 0 {
		stat := status.New(codes.InvalidArgument, "invalid knight request")
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return nil, s.Err()
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {

		g.logger.Error("err", err, "stack", debug.Stack())

		return nil, status.Errorf(codes.Internal, "internal error")
	}
	knt := knight.NewKnight(req.Username, req.Email, hashedPassword)
	kntR, err := g.service.AddKnight(ctx, knt)
	if err != nil {
		if errors.Is(err, knight.ErrEmailUnique) || errors.Is(err, knight.ErrUsernameUnique) {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		g.logger.Error("err", err, "stack", debug.Stack())
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &protoknight.AddKnightRes{Id: kntR.Id.String(), Username: kntR.Username, Email: kntR.Email, IsActive: kntR.IsActive}, nil
}
func (g *GrpcServer) GetKnightStatus(ctx context.Context, req *protoknight.KnightStatusReq) (*protoknight.KnightStatusRes, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req ")
	}
	validationErrors := validateGet(req.Username)
	if len(validationErrors) > 0 {
		stat := status.New(codes.InvalidArgument, "invalid knight request")
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return nil, s.Err()
	}
	isActive, err := g.service.KnightStatus(ctx, req.Username)
	if err != nil && errors.Is(err, commonerrors.ErrKnightNotFound) {

		return nil, status.Errorf(codes.NotFound, err.Error())

	} else if err != nil {

		g.logger.Error("err", err, "stack", debug.Stack())

		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &protoknight.KnightStatusRes{IsActive: isActive}, nil
}
func (g *GrpcServer) UpdateStatus(ctx context.Context, req *protoknight.UpdateStatusReq) (*protoknight.UpdateStatusRes, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req ")
	}
	err := g.service.UpdateStatus(ctx, req.Username, req.Active)
	if err != nil && errors.Is(err, commonerrors.ErrKnightNotFound) {

		g.logger.Error("err", err, "stack", debug.Stack())
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {

		g.logger.Error("err", err, "stack", debug.Stack())
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &protoknight.UpdateStatusRes{}, nil
}
