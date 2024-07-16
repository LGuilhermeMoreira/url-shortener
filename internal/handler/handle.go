package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

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

	if !validate(input.URL) {
		msg := entity.NewHandleError("URL está com problema", http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(msg)
		return
	}

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
	output := input.ConvertToOutput(http.StatusCreated, model.ShortID)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		msg := entity.NewHandleError("Error ao fazer o encode: "+err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
	}
}
func (h *Handler) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	model, err := h.DB.FindByID(id)
	if err != nil {
		msg := entity.NewHandleError("Erro ao buscar ID no banco de dados: "+err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}
	fmt.Println(model)
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

func validate(url string) bool {
	if url == "" {
		return false
	}
	regex := regexp.MustCompile(`^(http|https):\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(\/\S*)?$`)
	isValid := regex.MatchString(url)
	return isValid
}
