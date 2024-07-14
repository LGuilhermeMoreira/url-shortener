package database

import (
	"database/sql"
	"log"

	"github.com/LGuilhermeMoreira/url-shortener/internal/infra/model"
)

type UrlDb struct {
	Db *sql.DB
}

func (d UrlDb) Store(entity *model.Url) error {
	stmt, err := d.Db.Prepare("INSERT INTO urls(short_id,url) VALUES ($1,$2)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(entity.ShortID, entity.CompleteUrl)
	if err != nil {
		return err
	}

	return nil
}

func (d UrlDb) FindByID(key string) (*model.Url, error) {
	var id, url string

	err := d.Db.QueryRow("SELECT short_id,url FROM urls WHERE short_id = $1", key).Scan(&id, &url)
	if err != nil {
		return nil, err
	}

	return &model.Url{
		CompleteUrl: url,
		ShortID:     id,
	}, nil
}

func exist(db *sql.DB, url string) bool {
	var id string
	err := db.QueryRow("SELECT short_id FROM urls WHERE url = $1").Scan(&id)
	log.Printf("Erro in exist function, querying this %v url: %v\n", url, err)
	if id == "" {
		return false
	} else {
		return true
	}
}
