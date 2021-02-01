package handler

import (
	"Project_store/models"
	"Project_store/service"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// function for testing {/product} endpoint
func TestHandler_ReturnProductResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	ps := service.NewMockService(ctrl)
	handler := Handler{ps}

	product := []models.Product{
		{1, "bat", 1},
		{2, "ball", 2},
		{3, "wicket", 1},
	}

	for _, tc := range product {
		url := "/product/%v"
		req, err := http.NewRequest("GET", fmt.Sprintf(url, tc.Id), nil)
		if err != nil {
			t.Fatalf("an error '%s' was not expected while creating request", err)
		}
		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.Itoa(int(tc.Id)),
		})
		// returns an initialized ResponseRecorder
		w := httptest.NewRecorder()
		handler.ReturnProductResult(w, req)
		if w.Code != 200 {
			t.Fatalf("expected status code to be 200, but got: %d", w.Code)
		}
	}
}

// function for testing {/insert/product} endpoint
func TestHandler_InsertProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	ps := service.NewMockService(ctrl)
	handler := Handler{ps}

	testcases := []struct {
		productName, brandName string
	} {
		{"ball", "spartan"},
		{"wicket", "kashmiri"},
	}
	for _, tc := range testcases {
		url := "/product/insert"
		req, err := http.NewRequest("POST", fmt.Sprintf(url, tc), nil)
		if err != nil {
			t.Fatalf("an error '%s' was not expected while creating request", err)
		}

		w := httptest.NewRecorder()
		handler.InsertProduct(w, req)
		if w.Code != 201 {
			t.Fatalf("expected status code to be 201, but got: %d", w.Code)
		}
	}
}