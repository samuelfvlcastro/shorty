package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"smashedbits.com/shorty/pkg/model"
)

const findUserByIDQuery = "select id, email, dark_mode from users where id = $1"
const findUserByEmailQuery = "select id, email, dark_mode from users where email = $1"
const insertUserQuery = "insert into users (email) VALUES ($1)"
const updateDarkModeQuery = "UPDATE users SET email=$1, dark_mode=$2 WHERE id = $3"

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
		&user.DarkMode,
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
		&user.DarkMode,
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

func (u users) UpdateUser(ctx context.Context, user model.User) error {
	_, err := u.conn.Exec(ctx, updateDarkModeQuery, user.Email, user.DarkMode, user.ID)
	if err != nil {
		return err
	}

	return nil
}
