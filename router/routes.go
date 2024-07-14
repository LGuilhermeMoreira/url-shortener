package router

import (
	"net/http"

	"github.com/LGuilhermeMoreira/url-shortener/internal/handler"
	"github.com/LGuilhermeMoreira/url-shortener/internal/infra/database"
)

func CreateRouter(db database.UrlDb) *http.ServeMux {
	h := handler.NewHandler(db)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /encurtar", h.HandleGenerateShortID)
	mux.HandleFunc("GET /{id}", h.HandleRedirect)
	return mux
}
