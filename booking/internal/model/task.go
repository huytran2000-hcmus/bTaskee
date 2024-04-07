package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	Cleaning  ServiceType = "cleaning"
	ChildCare ServiceType = "childcare"

	TwoRoom   HouseType = "two_room"
	ThreeRoom HouseType = "three_room"
	FourRoom  HouseType = "four_room"
)

type ServiceType string

type HouseType string

type CreateTaskRequest struct {
	Customer         Customer       `json:"customer" validate:"required"`
	AssignedLocation Location       `json:"assigned_location" validate:"required"`
	WorkingDetails   WorkingDetails `json:"working_details" validate:"required"`
	Tasker           Tasker         `json:"tasker" validate:"required"`
	Note             string         `json:"note" example:"Be here ontime"`
}

type Task struct {
	ID               uuid.UUID      `json:"id"`
	Customer         Customer       `json:"customer"`
	AssignedLocation Location       `json:"assigned_location"`
	WorkingDetails   WorkingDetails `json:"working_details"`
	Tasker           Tasker         `json:"tasker"`
	Note             string         `json:"note"`
	Total            float64        `json:"total"`
}

type Tasker struct {
	ID             string         `json:"id" validate:"required" example:"1"`
	Name           string         `json:"name" validate:"required" example:"huy tran"`
	Email          string         `json:"email" validate:"required,email" example:"minhhuydev2000@gmail.com"`
	Phone          string         `json:"phone" validate:"required,e164" example:"+84948337945"`
	Identification Identification `json:"identification" validate:"required"`
}

type Identification struct {
	CMND string `json:"cmnd" validate:"number,len=9" example:"888888888"`
	CCCD string `json:"cccd" validate:"number,len=12" example:"999999999999"`
}

type Customer struct {
	Name  string `json:"name" validate:"required" example:"huy tran"`
	Email string `json:"email" validate:"required,email" example:"minhhuy@gmail.com"`
	Phone string `json:"phone" validate:"required,e164" example:"+84948337945"`
}

type Location struct {
	Address string `json:"address" validate:"required" example:"60 Bến Thành, Quận 1, TP.Hồ Chí Minh"`
}

type WorkingDetails struct {
	HouseType    HouseType     `json:"house_type" validate:"required" example:"two_room"`
	ServiceTypes []ServiceType `json:"service_types" validate:"required,gte=1" example:"cleaning"`
	FromTime     time.Time     `json:"from_time" validate:"required,ltfield=ToTime" example:"2024-04-07T09:00:00+07:00"`
	ToTime       time.Time     `json:"to_time" validate:"required,gtfield=FromTime" example:"2024-04-07T11:00:00+07:00"`
}

func CreateTaskRequestToTask(req *CreateTaskRequest) *Task {
	return &Task{
		Customer: Customer{
			Name:  req.Customer.Name,
			Email: req.Customer.Email,
			Phone: req.Customer.Phone,
		},
		AssignedLocation: req.AssignedLocation,
		WorkingDetails:   req.WorkingDetails,
		Tasker:           req.Tasker,
		Note:             req.Note,
	}
}
