package services

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"smashedbits.com/shorty/pkg/model"
)

const hashSalt = "aCoolHashSaltUsedForCollisions"

type urlStorage interface {
	GetURL(ctx context.Context, hash string) (model.URL, error)
	GetURLs(ctx context.Context, userId string) ([]model.URL, error)
	Insert(ctx context.Context, url model.URL) error
}

type Shortener struct {
	store urlStorage
}

func NewShortener(store urlStorage) Shortener {
	return Shortener{
		store: store,
	}
}

func (s Shortener) GetURL(ctx context.Context, hash string) (model.URL, error) {
	return s.store.GetURL(ctx, hash)
}

func (s Shortener) GetUserURLs(ctx context.Context, userId string) ([]model.URL, error) {
	return s.store.GetURLs(ctx, userId)
}

func (s Shortener) InsertURL(ctx context.Context, userId string, longURL string) (model.URL, error) {
	hash := s.hashFunc(longURL + hashSalt)
	url := model.URL{
		UserID:  userId,
		Hash:    hash[0:7],
		LongURL: longURL,
	}

	if err := s.store.Insert(ctx, url); err != nil {
		saltedUrl := longURL + hashSalt
		return s.InsertURL(ctx, userId, saltedUrl)
	}

	return url, nil
}

func (s Shortener) hashFunc(long string) string {
	hash := md5.Sum([]byte(long))
	return hex.EncodeToString(hash[:])
}
