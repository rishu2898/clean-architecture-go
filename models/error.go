package models

type ReturnError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}
