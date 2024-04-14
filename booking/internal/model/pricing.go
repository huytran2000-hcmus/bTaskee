package model

import "time"

type CalculatePriceRequest struct {
	HouseType    string        `json:"house_type"`
	ServiceTypes []ServiceType `json:"service_types"`
	FromTime     time.Time     `json:"from_time"`
	ToTime       time.Time     `json:"to_time"`
}

type CalculatePriceResponse struct {
	Total float64 `json:"total"`
}

type CalculatePricePayload struct {
	Data  CalculatePriceResponse `json:"data"`
	Error string                 `json:"error"`
}
