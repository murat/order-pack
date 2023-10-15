package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"order-pack/internal/order"
	"order-pack/internal/product"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Api struct {
	ProductSvc *product.Service
	OrderSvc   *order.Service
}

func NewApi(productSvc *product.Service, orderSvc *order.Service) *Api {
	return &Api{
		ProductSvc: productSvc,
		OrderSvc:   orderSvc,
	}
}

func (h *Api) RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
}

func (h *Api) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(NewSuccessResponse("hello world!"))
}

func (h *Api) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	size, err := strconv.ParseInt(r.PostFormValue("size"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write(NewErrorResponse(err.Error()))
		return
	}

	p := product.Product{
		Name:        r.PostFormValue("name"),
		PackageSize: int(size),
	}
	if err := h.ProductSvc.Create(&p); err != nil {
		log.Printf("could not create product, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(NewErrorResponse(err.Error()))
		return
	}

	_, _ = w.Write(NewSuccessResponse(p))
}

func (h *Api) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	// ignored error due to route constraints
	val, _ := getParam(r, "id")
	id, _ := strconv.ParseUint(val, 10, 64)
	p, err := h.ProductSvc.Find(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			log.Printf("could not fetch product, %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		_, _ = w.Write(NewErrorResponse(err.Error()))
		return
	}

	_, _ = w.Write(NewSuccessResponse(p))
}

func (h *Api) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductSvc.Get()
	if err != nil {
		log.Printf("could not fetch products, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(NewErrorResponse(err.Error()))
		return
	}

	_, _ = w.Write(NewSuccessResponse(products))
}

type CreateOrderRequest struct {
	ItemCount int `json:"item_count" validate:"required"`
}

func (h *Api) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write(NewErrorResponse(err.Error()))
		return
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(req); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write(NewErrorResponse(err.Error()))
		return
	}

	products, err := h.ProductSvc.GetSortedBy("package_size DESC")
	if err != nil {
		log.Printf("could not fetch product, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(NewErrorResponse(err.Error()))
		return
	}

	_, _ = w.Write(NewSuccessResponse(order.FulfillOrder(&order.Order{ItemCount: req.ItemCount}, products)))
}

func getParam(r *http.Request, param string) (string, error) {
	val, ok := mux.Vars(r)[param]
	if !ok {
		return "", fmt.Errorf("%s parameter is not provided", param)
	}

	return strings.TrimSpace(val), nil
}
