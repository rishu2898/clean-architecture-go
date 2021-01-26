package handler

import (
	"Project_store/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	s service.Service
}
func New(s service.Service) Handle {
	return &Handler{s}
}
func (empHandler Handler) ReturnProductResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		key, _ := strconv.Atoi(vars["id"])
		data := empHandler.s.GetProductDetails(key)
		json.NewEncoder(w).Encode(data)
	}
}