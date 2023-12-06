package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"smashedbits.com/shorty/pkg/model"
)

const insertUrlQuery = "insert into urls (user_id, hash, long_url) VALUES ($1, $2, $3)"
const fetchUserUrlsQuery = "select id, user_id, hash, long_url, created_at from urls where user_id=$1"
const fetchUrlByHash = "select id, user_id, hash, long_url, created_at from urls where hash=$1"

type urls struct {
	conn *pgx.Conn
}

func NewURLs(conn *pgx.Conn) urls {
	return urls{
		conn: conn,
	}
}

func (u urls) GetURL(ctx context.Context, hash string) (model.URL, error) {
	url := model.URL{}
	if err := u.conn.QueryRow(ctx, fetchUrlByHash, hash).Scan(
		&url.ID,
		&url.UserID,
		&url.Hash,
		&url.LongURL,
		&url.CreatedAt,
	); err != nil {
		return model.URL{}, err
	}

	return url, nil
}

func (u urls) GetURLs(ctx context.Context, userId string) ([]model.URL, error) {
	urls := []model.URL{}
	rows, err := u.conn.Query(ctx, fetchUserUrlsQuery, userId)
	if err != nil {
		return urls, err
	}

	for rows.Next() {
		url := model.URL{}
		if err := rows.Scan(&url.ID, &url.UserID, &url.Hash, &url.LongURL, &url.CreatedAt); err != nil {
			return urls, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (u urls) Insert(ctx context.Context, url model.URL) error {
	_, err := u.conn.Exec(ctx, insertUrlQuery, url.UserID, url.Hash, url.LongURL)
	if err != nil {
		return err
	}

	return nil
}
