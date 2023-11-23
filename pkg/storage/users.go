package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"smashedbits.com/shorty/pkg/model"
)

const findUserByIDQuery = "select id, email from users where id = $1"
const findUserByEmailQuery = "select id, email from users where email = $1"
const insertUserQuery = "insert into users (email) VALUES ($1)"

type users struct {
	conn *pgx.Conn
}

func NewUsers(conn *pgx.Conn) users {
	return users{
		conn: conn,
	}
}

func (u users) GetByID(ctx context.Context, userId string) (model.User, error) {
	user := model.User{}
	if err := u.conn.QueryRow(ctx, findUserByIDQuery, userId).Scan(
		&user.ID,
		&user.Email,
	); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u users) GetByEmail(ctx context.Context, email string) (model.User, error) {
	user := model.User{}
	if err := u.conn.QueryRow(ctx, findUserByEmailQuery, email).Scan(
		&user.ID,
		&user.Email,
	); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u users) Insert(ctx context.Context, user model.User) error {
	_, err := u.conn.Exec(ctx, insertUserQuery, user.Email)
	if err != nil {
		return err
	}

	return nil
}
