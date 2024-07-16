package dto

import (
	"github.com/LGuilhermeMoreira/url-shortener/internal/infra/model"
	"github.com/LGuilhermeMoreira/url-shortener/internal/utils"
)

type InputUrl struct {
	URL string `json:"url"`
}

type OutputUrl struct {
	URL     string `json:"complete_url"`
	ShortID string `json:"short_id"`
	Status  uint   `json:"status"`
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

func (i InputUrl) ConvertToOutput(status uint, short_id string) OutputUrl {
	return OutputUrl{
		URL:     i.URL,
		ShortID: short_id,
		Status:  status,
	}
}
