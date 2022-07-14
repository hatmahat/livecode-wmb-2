package dto

import "go-api-with-gin2/model"

type MenuResponse struct {
	Code    string     `json:"responseCode"`
	Message string     `json:"responseMessage"`
	Data    model.Menu `json:"data"`
}

type ListMenuResponse struct {
	Code    string       `json:"responseCode"`
	Message string       `json:"responseMessage"`
	Data    []model.Menu `json:"data"`
}

type ErrorCode struct {
	Code    string `json:"responseCode"`
	Message string `json:"responseMessage"`
}
