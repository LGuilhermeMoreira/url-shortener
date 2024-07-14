package database

import "github.com/LGuilhermeMoreira/url-shortener/internal/infra/model"

type Repository interface {
	Store(*model.Url) error
	FindByID(string) (*model.Url, error)
}
