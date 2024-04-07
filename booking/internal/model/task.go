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
	Note             string         `json:"note"`
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
	ID             string         `json:"id" validate:"required"`
	Name           string         `json:"name" validate:"required"`
	Email          string         `json:"email" validate:"required,email"`
	Phone          string         `json:"phone" validate:"required,e164"`
	Identification Identification `json:"identification" validate:"required"`
}

type Identification struct {
	CMND string `json:"cmnd" validate:"number,len=9"`
	CCCD string `json:"cccd" validate:"number,len=12"`
}

type Customer struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required,e164"`
}

type Location struct {
	Address string `json:"address" validate:"required"`
}

type WorkingDetails struct {
	HouseType    HouseType     `json:"house_type" validate:"required"`
	ServiceTypes []ServiceType `json:"service_types" validate:"required,gte=1"`
	FromTime     time.Time     `json:"from_time" validate:"required,ltfield=ToTime"`
	ToTime       time.Time     `json:"to_time" validate:"required,gtfield=FromTime"`
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
