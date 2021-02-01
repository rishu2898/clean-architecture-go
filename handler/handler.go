package handler

import (
	"Project_store/models"
	"Project_store/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	s service.Service
}

func New(s service.Service) Handler {
	return Handler{s}
}

type bucket struct {
	ProductName string	`json:"productName"`
	BrandName string	`json:"brandName"`
}
func (prod Handler) InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data bucket
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			temp := models.ReturnError{http.StatusInternalServerError, "error in decoding data"}
			err, _ := json.Marshal(temp)
			w.Write(err)
		}
		res, err := prod.s.InsertProduct(data.ProductName, data.BrandName)
		if err != nil {
			temp := models.ReturnError{http.StatusBadRequest, "error in insertion"}
			err, _ := json.Marshal(temp)
			w.Write(err)
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}

func (prod Handler) ReturnProductResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		key, _ := strconv.Atoi(vars["id"])
		data, err := prod.s.GetProductDetails(int64(key))
		if err != nil {
			temp := models.ReturnError{http.StatusInternalServerError, "error in fatching key id from url"}
			err, _ := json.Marshal(temp)
			w.Write(err)
		}
		json.NewEncoder(w).Encode(data)
	}
}