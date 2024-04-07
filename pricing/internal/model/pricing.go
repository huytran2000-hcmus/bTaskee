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
	HouseType    HouseType     `json:"house_type" validate:"required"`
	ServiceTypes []ServiceType `json:"service_types" validate:"required,gte=1"`
	FromTime     time.Time     `json:"from_time" validate:"required,ltfield=ToTime"`
	ToTime       time.Time     `json:"to_time" validate:"required,gtfield=FromTime"`
}

type CalculatePriceResponse struct {
	Total float64 `json:"total"`
}
