package handler

import (
	"encoding/json"
	"net/http"

	"github.com/LGuilhermeMoreira/url-shortener/internal/dto"
	"github.com/LGuilhermeMoreira/url-shortener/internal/infra/database"
	"github.com/LGuilhermeMoreira/url-shortener/pkg/entity"
	"github.com/LGuilhermeMoreira/url-shortener/public"
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
	http.Redirect(w, r, model.CompleteUrl, http.StatusMovedPermanently)
}

func (h *Handler) HandleTempl(w http.ResponseWriter, r *http.Request) {
	public.Encurtar().Render(r.Context(), w)
}

func (h *Handler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("CotentType", "application/json")
	body := map[string]interface{}{
		"Ping": "Pong",
	}
	json.NewEncoder(w).Encode(body)
}
