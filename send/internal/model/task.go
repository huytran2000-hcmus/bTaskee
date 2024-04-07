package model

import (
	"time"
)

type SendTaskRequest struct {
	ID               string         `json:"id"`
	Customer         Customer       `json:"customer"`
	Tasker           Tasker         `json:"tasker"`
	AssignedLocation Location       `json:"assigned_location"`
	WorkingDetails   WorkingDetails `json:"working_details"`
	Note             string         `json:"note"`
	Total            float64        `json:"total"`
}

type SendTaskResponse struct {
	EmailIDs []string `json:"email_ids"`
}

type Tasker struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
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
	HouseType    string    `json:"house_type"`
	ServiceTypes []string  `json:"service_types"`
	From         time.Time `json:"from"`
	To           time.Time `json:"to"`
}
