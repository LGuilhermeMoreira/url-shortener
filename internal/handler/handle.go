package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LGuilhermeMoreira/url-shortener/internal/dto"
	"github.com/LGuilhermeMoreira/url-shortener/internal/infra/database"
	"github.com/LGuilhermeMoreira/url-shortener/pkg/entity"
)

type Handler struct {
	DB database.Repository
}

func NewHandler(db database.Repository) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h *Handler) HandleGenerateShortID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var input dto.InputUrl
	json.NewDecoder(r.Body).Decode(&input)
	model, err := input.ConvertToModel()
	log.Println(model)
	if err != nil {
		msg := entity.NewHandleError("Erro ao converter input para model: "+err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = h.DB.Store(model)
	if err != nil {
		msg := entity.NewHandleError("Erro ao inserir no banco de dados: "+err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}
	json.NewEncoder(w).Encode(model)
}
func (h *Handler) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Println(id)
	model, err := h.DB.FindByID(id)
	if err != nil {
		msg := entity.NewHandleError("Erro ao buscar ID no banco de dados", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if model == nil {
		msg := entity.NewHandleError("ID não encontrado", http.StatusNotFound)
		json.NewEncoder(w).Encode(msg)
		return
	}
	log.Println(model.CompleteUrl)
	http.Redirect(w, r, model.CompleteUrl, http.StatusMovedPermanently)
}
