package postgres

import (
	"context"
	"database/sql"

	"github.com/backend-magang/eniqilo-store/models/entity"
)

func (r *repository) FindUserByPhoneNumber(ctx context.Context, phoneNumber string) (result entity.User, err error) {
	query := `
		SELECT * FROM users WHERE
		phone_number = $1
	`

	err = r.db.QueryRowxContext(ctx, query, phoneNumber).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Errorf("[Repository][User][FindUserByPhoneNumber] failed to find user by phone number %s, err: %s", phoneNumber, err.Error())
		return
	}

	return
}

func (r *repository) FindUserByID(ctx context.Context, id string) (result entity.User, err error) {
	query := `
		SELECT * FROM users WHERE
		id = $1
	`

	err = r.db.QueryRowxContext(ctx, query, id).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Errorf("[Repository][User][FindUserByID] failed to find user id %s, err: %s", id, err.Error())
		return
	}

	return
}

func (r *repository) InsertUser(ctx context.Context, req entity.User) (result entity.User, err error) {

	query := `INSERT INTO users (id, name, phone_number, role, password, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING *`

	err = r.db.QueryRowxContext(ctx,
		query,
		req.ID,
		req.Name,
		req.PhoneNumber,
		req.Role,
		req.Password,
		req.CreatedAt,
	).StructScan(&result)

	if err != nil {
		r.logger.Errorf("[Repository][User][InsertUser] failed to insert new user, err: %s", err.Error())
		return
	}

	return
}
