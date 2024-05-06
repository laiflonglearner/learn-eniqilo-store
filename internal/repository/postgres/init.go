package postgres

import (
	"github.com/backend-magang/cats-social-media/config"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RepositoryHandler interface {
	// GetListCat(ctx context.Context, req entity.GetListCatRequest) (result []entity.Cat, err error)
	// FindCatByID(ctx context.Context, id int) (entity.Cat, error)
	// FindUserCatByID(ctx context.Context, userId int, catId int) (result entity.Cat, err error)
	// InsertCat(ctx context.Context, req entity.Cat) (result entity.Cat, err error)
	// UpdateCat(ctx context.Context, req entity.Cat) (result entity.Cat, err error)
	// DeleteCat(ctx context.Context, req entity.Cat) (result entity.Cat, err error)

	// FindUserByEmail(ctx context.Context, email string) (result entity.User, err error)
	// InsertUser(ctx context.Context, req entity.User) (result entity.User, err error)

	// FindRequestedMatch(ctx context.Context, catId int) (result []entity.MatchCat, err error)
	// FindMatchByID(ctx context.Context, id int) (entity.MatchCat, error)
	// InsertMatchCat(ctx context.Context, req entity.MatchCat) (err error)
	// UpdateMatchCat(ctx context.Context, req entity.MatchCat) (err error)
	// GetListMatchCat(ctx context.Context, req entity.GetListMatchCatRequest) (result []entity.GetListMatchCatQueryResponse, err error)
	// FindMatchCatByID(ctx context.Context, id int) (result entity.MatchCat, err error)
	// ApproveMatch(ctx context.Context, matchCatId int) (err error)
	// UpdateCatsMatch(ctx context.Context, matchCat entity.MatchCat) (err error)
	// DeleteOtherMatch(ctx context.Context, catId int, matchCatId int) (err error)
	// FindRequestedMatchCat(ctx context.Context, catId int, matchCatId int) (result []entity.MatchCat, err error)
}

type repository struct {
	cfg    config.Config
	db     *sqlx.DB
	logger *logrus.Logger
}

func NewRepository(cfg config.Config, db *sqlx.DB, log *logrus.Logger) RepositoryHandler {
	return &repository{
		cfg:    cfg,
		db:     db,
		logger: log,
	}
}