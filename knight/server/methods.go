package grpcserver

import (
	"context"
	"errors"

	commonerrors "github.com/lafetz/quest-micro/common/errors"
	knight "github.com/lafetz/quest-micro/knight/core"
	protoknight "github.com/lafetz/quest-micro/proto/gen"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GrpcServer) AddKnight(ctx context.Context, req *protoknight.AddKnightReq) (*protoknight.AddKnightRes, error) {

	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req ")
	}
	validationErrors := validateAddknight(req.Email, req.Name)
	if len(validationErrors) > 0 {
		stat := status.New(codes.InvalidArgument, "There was a problem with the provided data")
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
	return &protoknight.AddKnightRes{Id: kntR.Id.String(), Name: kntR.Name, Email: kntR.Email, IsActive: kntR.IsActive}, nil
}

func (g *GrpcServer) GetKnightStatus(ctx context.Context, req *protoknight.KnightStatusReq) (*protoknight.KnightStatusRes, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req ")
	}
	validationErrors := validateGet(req.Email)
	if len(validationErrors) > 0 {
		stat := status.New(codes.InvalidArgument, "invalid knight request")
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
	return &protoknight.KnightStatusRes{IsActive: isActive}, nil
}

func (g *GrpcServer) UpdateStatus(ctx context.Context, req *protoknight.UpdateStatusReq) (*protoknight.UpdateStatusRes, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req ")
	}
	err := g.knightService.UpdateStatus(ctx, req.Email, req.Active)
	if err != nil && errors.Is(err, commonerrors.ErrKnightNotFound) {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &protoknight.UpdateStatusRes{}, nil
}
func (g *GrpcServer) GetKnights(ctx context.Context, req *protoknight.GetKnightsReq) (*protoknight.GetKnightsRes, error) {
	knights, err := g.knightService.GetKnights(ctx)
	if err != nil {
		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	var responseKnights []*protoknight.Knight
	for _, knight := range knights {
		responseKnights = append(responseKnights, &protoknight.Knight{
			Id:       knight.Id.String(),
			Name:     knight.Name,
			Email:    knight.Email,
			IsActive: knight.IsActive,
		})
	}

	return &protoknight.GetKnightsRes{
		Knights: responseKnights,
	}, nil
}
func (g *GrpcServer) GetKnight(ctx context.Context, req *protoknight.GetKnightReq) (*protoknight.GetKnightRes, error) {
	knight, err := g.knightService.GetKnight(ctx, req.GetEmail())
	if err != nil && errors.Is(err, commonerrors.ErrKnightNotFound) {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &protoknight.GetKnightRes{
		Knight: &protoknight.Knight{
			Id:       knight.Id.String(),
			Name:     knight.Name,
			Email:    knight.Email,
			IsActive: knight.IsActive,
		},
	}, nil
}
func (g *GrpcServer) DeleteKnight(ctx context.Context, req *protoknight.DeleteKnightReq) (*protoknight.DeleteKnightRes, error) {
	err := g.knightService.DeleteKnight(ctx, req.GetEmail())
	if err != nil && errors.Is(err, commonerrors.ErrKnightNotFound) {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {

		g.logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &protoknight.DeleteKnightRes{}, nil
}
