package model

import "time"

const (
	Cleaning  ServiceType = "cleaning"
	ChildCare ServiceType = "childcare"

	TwoRoom   HouseType = "two_room"
	ThreeRoom HouseType = "three_room"
	FourRoom  HouseType = "four_room"
)

type ServiceType string

type HouseType string

type CalculatePriceRequest struct {
	HouseType    HouseType     `json:"house_type"`
	ServiceTypes []ServiceType `json:"service_types"`
	From         time.Time     `json:"from"`
	To           time.Time     `json:"to"`
}

type PricingResponse struct {
	Total float64 `json:"total"`
}
