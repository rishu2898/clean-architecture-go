package handler

import "net/http"

type Handle interface {
	ReturnProductResult(w http.ResponseWriter, r *http.Request)
}
