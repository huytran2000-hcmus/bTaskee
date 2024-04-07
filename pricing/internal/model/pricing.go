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
	HouseType    HouseType     `json:"house_type" validate:"required" example:"three_room"`
	ServiceTypes []ServiceType `json:"service_types" validate:"required,gte=1" example:"childcare"`
	FromTime     time.Time     `json:"from_time" validate:"required,ltfield=ToTime" example:"2024-04-07T09:00:00+07:00"`
	ToTime       time.Time     `json:"to_time" validate:"required,gtfield=FromTime" example:"2024-04-07T11:00:00+07:00"`
}

type CalculatePriceResponse struct {
	Total float64 `json:"total"`
}
