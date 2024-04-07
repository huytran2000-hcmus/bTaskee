package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	Cleaning  ServiceType = "cleaning"
	ChildCare ServiceType = "childcare"

	TwoRoom   HouseType = "two_room"
	ThreeType HouseType = "three_room"
	FourType  HouseType = "four_room"
)

type ServiceType string

type HouseType string

type CreateTaskRequest struct {
	Customer         Customer       `json:"customer"`
	AssignedLocation Location       `json:"assigned_location"`
	WorkingDetails   WorkingDetails `json:"working_details"`
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
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	Phone          string         `json:"phone"`
	Identification Identification `json:"identification"`
}

type Identification struct {
	CMND string `json:"cmnd"`
	CCCD string `json:"cccd"`
}

type Customer struct {
	Name  string `json:"customer_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Location struct {
	Address string `json:"address"`
}

type WorkingDetails struct {
	HouseType    HouseType     `json:"house_type"`
	ServiceTypes []ServiceType `json:"service_types"`
	From         time.Time     `json:"from"`
	To           time.Time     `json:"to"`
}

func CreateTaskRequestToTask(req *CreateTaskRequest) *Task {
	return &Task{
		Customer: Customer{
			Name:  req.Customer.Name,
			Email: req.Customer.Email,
			Phone: req.Customer.Phone,
		},
		AssignedLocation: Location{
			Address: req.AssignedLocation.Address,
		},
		WorkingDetails: WorkingDetails{
			HouseType:    req.WorkingDetails.HouseType,
			ServiceTypes: req.WorkingDetails.ServiceTypes,
			From:         req.WorkingDetails.From,
			To:           req.WorkingDetails.To,
		},
		Note: req.Note,
	}
}
