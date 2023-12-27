package handler

import (
	"encoding/json"
	"net/http"

	"github.com/NikitaBarysh/wb_L0/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Service
}

func NewHandler(newService *service.Service) *Handler {
	return &Handler{services: newService}
}

func (h *Handler) getOrder(rw http.ResponseWriter, r *http.Request) {
	orderUid := chi.URLParam(r, "orderUID")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	order, err := h.services.Order.GetOrder(orderUid)
	if err != nil {
		http.Error(rw, `{"data":"Order not found"}`, http.StatusBadRequest)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(order)

}
