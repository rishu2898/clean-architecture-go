package handler

import (
	"Project_store/service"
	"encoding/json"
	"fmt"
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

type bucket struct {
	ProductName string	`json:"productName"`
	BrandName string	`json:"brandName"`
}
func (prod Handler) InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data bucket
		err := json.NewDecoder(r.Body).Decode(&data)
		fmt.Println(data)
		if err != nil {
			panic(err)
		}
		res, err := prod.s.InsertProduct(data.ProductName, data.BrandName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
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
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(data)
	}
}