package handler

import (
	"Project_store/models"
	"Project_store/service"
	"bytes"
	"encoding/json"
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
	brand := []models.Brand{
		{1, "reebok"},
		{2, "sparten"},
		{},
	}
	expect := []models.Result{
		{1, "bat", brand[0].Name},
		{2, "ball", brand[1].Name},
		{3, "wicket", brand[0].Name},
	}

	for i, tc := range product {
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
		ps.EXPECT().GetProductDetails(tc.Id).Return(expect[i], err)
		handler.ReturnProductResult(w, req)

		if w.Code != 200 || err != nil {
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
		payload bucket
		expected models.Result
	}{
		{
			payload: bucket {
				ProductName: "ball",
				BrandName:   "spartan",
			},
			expected: models.Result {
				Id:    1,
				Name:  "ball",
				Bname: "spartan",
			},
		},
	}

	for _, tc := range testcases {
		url := "/product/insert"

		b, _ := json.Marshal(tc.payload)

		req, err := http.NewRequest("POST", fmt.Sprintf(url), bytes.NewBuffer(b))
		if err != nil {
			t.Fatalf("an error '%s' was not expected while creating request", err)
		}

		w := httptest.NewRecorder()

		ps.EXPECT().InsertProduct(tc.payload.ProductName, tc.payload.BrandName).Return(tc.expected, nil)
		handler.InsertProduct(w, req)

		if w.Code != 201 || err != nil {
			t.Fatalf("expected status code to be 201, but got: %d", w.Code)
		}
	}
}
