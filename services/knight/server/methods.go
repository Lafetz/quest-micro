package knightserver

import (
	"context"
	"errors"

	commonerrors "github.com/lafetz/quest-micro/common/errors"
	knightv1 "github.com/lafetz/quest-micro/proto/gen/knight/v1"

	knight "github.com/lafetz/quest-micro/services/knight/core"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrValidate = errors.New("there was a problem with the provided data")
)

func (g *KnightServer) AddKnight(ctx context.Context, req *knightv1.AddKnightRequest) (*knightv1.AddKnightResponse, error) {

	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req ")
	}
	validationErrors := validateAddknight(req.Email, req.Name)
	if len(validationErrors) > 0 {
		stat := status.New(codes.InvalidArgument, ErrValidate.Error())
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return nil, s.Err()
	}
	knt := knight.NewKnight(req.Name, req.Email)
	kntR, err := g.knightService.AddKnight(ctx, knt)
	if err != nil {
		if errors.Is(err, knight.ErrEmailUnique) {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &knightv1.AddKnightResponse{Id: kntR.Id.String(), Name: kntR.Name, Email: kntR.Email, IsActive: kntR.IsActive}, nil
}

func (g *KnightServer) GetKnightStatus(ctx context.Context, req *knightv1.GetKnightStatusRequest) (*knightv1.GetKnightStatusResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req ")
	}
	validationErrors := validateGet(req.Email)
	if len(validationErrors) > 0 {
		stat := status.New(codes.InvalidArgument, ErrValidate.Error())
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return nil, s.Err()

	}
	isActive, err := g.knightService.KnightStatus(ctx, req.Email)
	if err != nil && errors.Is(err, commonerrors.ErrKnightNotFound) {

		return nil, status.Errorf(codes.NotFound, err.Error())

	} else if err != nil {

		g.logger.Error(err.Error())

		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &knightv1.GetKnightStatusResponse{IsActive: isActive}, nil
}

func (g *KnightServer) UpdateStatus(ctx context.Context, req *knightv1.UpdateStatusRequest) (*knightv1.UpdateStatusResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req")
	}
	err := g.knightService.UpdateStatus(ctx, req.Email, req.Active)
	if err != nil && errors.Is(err, commonerrors.ErrKnightNotFound) {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &knightv1.UpdateStatusResponse{}, nil
}
func (g *KnightServer) GetKnights(ctx context.Context, req *knightv1.GetKnightsRequest) (*knightv1.GetKnightsResponse, error) {
	knights, err := g.knightService.GetKnights(ctx)
	if err != nil {
		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	var responseKnights []*knightv1.Knight
	for _, knight := range knights {
		responseKnights = append(responseKnights, &knightv1.Knight{
			Id:       knight.Id.String(),
			Name:     knight.Name,
			Email:    knight.Email,
			IsActive: knight.IsActive,
		})
	}

	return &knightv1.GetKnightsResponse{
		Knights: responseKnights,
	}, nil
}
func (g *KnightServer) GetKnight(ctx context.Context, req *knightv1.GetKnightRequest) (*knightv1.GetKnightResponse, error) {
	knight, err := g.knightService.GetKnight(ctx, req.GetEmail())
	if err != nil && errors.Is(err, commonerrors.ErrKnightNotFound) {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &knightv1.GetKnightResponse{
		Knight: &knightv1.Knight{
			Id:       knight.Id.String(),
			Name:     knight.Name,
			Email:    knight.Email,
			IsActive: knight.IsActive,
		},
	}, nil
}
func (g *KnightServer) DeleteKnight(ctx context.Context, req *knightv1.DeleteKnightRequest) (*knightv1.DeleteKnightResponse, error) {
	err := g.knightService.DeleteKnight(ctx, req.GetEmail())
	if err != nil && errors.Is(err, commonerrors.ErrKnightNotFound) {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &knightv1.DeleteKnightResponse{}, nil
}
