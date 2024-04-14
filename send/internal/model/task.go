package model

import (
	"time"
)

type SendTaskRequest struct {
	ID               string         `json:"id" validate:"required" example:"c4824b96-e4ca-49cf-aaea-156604799612"`
	Customer         Customer       `json:"customer" validate:"required"`
	AssignedLocation Location       `json:"assigned_location" validate:"required"`
	WorkingDetails   WorkingDetails `json:"working_details" validate:"required"`
	Tasker           Tasker         `json:"tasker" validate:"required"`
	Note             string         `json:"note" example:"Be here ontime"`
	Total            float64        `json:"total" validate:"required" example:"400"`
}

type SendTaskResponse struct {
	MessageIDs []string `json:"message_ids"`
}

type Tasker struct {
	Name  string `json:"name" validate:"required" example:"huy tran"`
	Email string `json:"email" validate:"required,email" example:"<enter your email here>"`
	Phone string `json:"phone" validate:"required,e164" example:"+84948447525"`
}

type Customer struct {
	Name  string `json:"name" validate:"required" example:"huy tran"`
	Email string `json:"email" validate:"required,email" example:"huy@gmail.com"`
	Phone string `json:"phone" validate:"required,e164" example:"+84948447524"`
}

type Location struct {
	Address string `json:"address" validate:"required" example:"60 Bến Thành, Quận 1, TP.Hồ Chí Minh"`
}

type WorkingDetails struct {
	HouseType    string    `json:"house_type" validate:"required" example:"two_room"`
	ServiceTypes []string  `json:"service_types" validate:"required,gte=1" example:"cleaning,childcare"`
	FromTime     time.Time `json:"from_time" validate:"required,ltfield=ToTime" example:"2024-04-07T09:00:00+07:00"`
	ToTime       time.Time `json:"to_time" validate:"required,gtfield=FromTime" example:"2024-04-07T11:00:00+07:00"`
}
