package handler

import (
	"Project_store/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	s service.Service
}

func New(s service.Service) Handler {
	return Handler{s}
}

func (empHandler Handler) ReturnProductResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		key, _ := strconv.Atoi(vars["id"])
		data, err := empHandler.s.GetProductDetails(key)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(data)
	}
}