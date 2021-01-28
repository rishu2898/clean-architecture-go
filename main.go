package main

import (
	"Project_store/handler"
	"Project_store/service"
	"Project_store/store/brand"
	"Project_store/store/product"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "rishabh:Rishu2898@@(127.0.0.1)/store")
	if err != nil {
		log.Fatal(err)
	}
	product := product.New(db)
	//fmt.Println(product.GetById(1))
	brand := brand.New(db)
	//fmt.Println(brand.GetById(1))

	service := service.New(product, brand)
	h := handler.New(service)

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/product/{id}", h.ReturnProductResult).Methods("GET")
	myRouter.HandleFunc("/product/insert", h.InsertProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
	defer db.Close()
}