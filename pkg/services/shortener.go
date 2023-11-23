package services

import (
	"context"

	"smashedbits.com/shorty/pkg/model"
)

type urlStorage interface {
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

func (s Shortener) GetUserURLs(ctx context.Context, userId string) ([]model.URL, error) {
	return s.store.GetURLs(ctx, userId)
}

func (s Shortener) InsertURL(ctx context.Context, userId string, longURL string) (model.URL, error) {
	hash := "fewg34t52"
	url := model.URL{
		UserID:  userId,
		Hash:    hash,
		LongURL: longURL,
	}
	if err := s.store.Insert(ctx, url); err != nil {
		return url, err
	}

	return url, nil
}
