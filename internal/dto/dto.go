package dto

import (
	"github.com/LGuilhermeMoreira/url-shortener/internal/infra/model"
	"github.com/LGuilhermeMoreira/url-shortener/internal/utils"
)

type InputUrl struct {
	URL string `json:"url"`
}

func (i InputUrl) ConvertToModel() (*model.Url, error) {
	id, err := utils.GenerateShortID()
	if err != nil {
		return nil, err
	}
	model := model.Url{
		CompleteUrl: i.URL,
		ShortID:     id,
	}
	return &model, nil
}
