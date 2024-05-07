package usecase

import (
	"context"

	"github.com/backend-magang/eniqilo-store/config"
	"github.com/backend-magang/eniqilo-store/internal/repository/postgres"
	"github.com/backend-magang/eniqilo-store/models"
	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/sirupsen/logrus"
)

type UsecaseHandler interface {
	RegisterStaff(ctx context.Context, req entity.RegisterStaffRequest) models.StandardResponseReq

	// GetListCat(ctx context.Context, req entity.GetListCatRequest) models.StandardResponseReq
	// CreateCat(ctx context.Context, req entity.CreateCatRequest) models.StandardResponseReq
	// UpdateCat(ctx context.Context, req entity.UpdateCatRequest) models.StandardResponseReq
	// DeleteCat(ctx context.Context, req entity.DeleteCatRequest) models.StandardResponseReq

	// MatchCat(ctx context.Context, req entity.MatchCatRequest) models.StandardResponseReq
	// RejectMatchCat(ctx context.Context, req entity.UpdateMatchCatRequest) models.StandardResponseReq
	// DeleteMatchCat(ctx context.Context, req entity.DeleteMatchCatRequest) models.StandardResponseReq
	// GetListMatchCat(ctx context.Context, req entity.GetListMatchCatRequest) models.StandardResponseReq
	// MatchApprove(ctx context.Context, req entity.MatchApproveRequest) models.StandardResponseReq

	// LoginUser(ctx context.Context, req entity.LoginUserRequest) models.StandardResponseReq
}

type usecase struct {
	cfg        config.Config
	logger     *logrus.Logger
	repository postgres.RepositoryHandler
}

func NewUsecase(cfg config.Config, log *logrus.Logger, repository postgres.RepositoryHandler) UsecaseHandler {
	return &usecase{
		cfg:        cfg,
		logger:     log,
		repository: repository,
	}
}
