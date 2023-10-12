package api

import (
	"encoding/json"
	"net/http"

	"order-pack/internal/pack"
)

type PackService interface {
	Get() ([]pack.Pack, error)
	Find() (pack.Pack, error)
	Create(p pack.Pack) error
	Update(p pack.Pack) error
	Delete(p pack.Pack) error
}

type Api struct {
	PackSvc PackService
}

func NewApi(packSvc PackService) *Api {
	return &Api{
		PackSvc: packSvc,
	}
}

func (h *Api) RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
}

func (h *Api) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world!"))
}

func (h *Api) GetPackagesHandler(w http.ResponseWriter, r *http.Request) {
	packs, err := h.PackSvc.Get()
	if err != nil {
		panic(err)
	}

	packBytes, _ := json.Marshal(packs)

	w.WriteHeader(http.StatusOK)
	w.Write(packBytes)
}
