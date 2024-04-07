package model

import (
	"time"
)

type SendTaskRequest struct {
	ID               string             `json:"id"`
	Customer         SendCustomer       `json:"customer"`
	Tasker           SendTasker         `json:"tasker"`
	AssignedLocation SendLocation       `json:"assigned_location"`
	WorkingDetails   SendWorkingDetails `json:"working_details"`
	Note             string             `json:"note"`
	Total            float64            `json:"total"`
}

type SendTaskPayload struct {
	Data  SendTaskResponse `json:"data"`
	Error string           `json:"error"`
}

type SendTaskResponse struct {
	MessageIDs []string `json:"message_ids"`
}

type SendTasker struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type SendCustomer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type SendLocation struct {
	Address string `json:"address"`
}

type SendWorkingDetails struct {
	HouseType    string    `json:"house_type" validate:"required"`
	ServiceTypes []string  `json:"service_types" validate:"required,gte=1"`
	FromTime     time.Time `json:"from_time" validate:"required,ltfield=ToTime"`
	ToTime       time.Time `json:"to_time" validate:"required,gtfield=FromTime"`
}

func ToSendTaskRequest(task *Task) SendTaskRequest {
	return SendTaskRequest{
		ID: task.ID.String(),
		Customer: SendCustomer{
			Name:  task.Customer.Name,
			Email: task.Customer.Email,
			Phone: task.Customer.Phone,
		},
		Tasker: SendTasker{
			Name:  task.Tasker.Name,
			Email: task.Tasker.Email,
			Phone: task.Tasker.Phone,
		},
		AssignedLocation: SendLocation{
			Address: task.AssignedLocation.Address,
		},
		WorkingDetails: SendWorkingDetails{
			HouseType:    string(task.WorkingDetails.HouseType),
			ServiceTypes: toSendServiceTypeStr(task.WorkingDetails.ServiceTypes),
			FromTime:     task.WorkingDetails.FromTime,
			ToTime:       task.WorkingDetails.ToTime,
		},
		Note:  task.Note,
		Total: task.Total,
	}
}

func toSendServiceTypeStr(services []ServiceType) []string {
	res := make([]string, 0, len(services))
	for _, svc := range services {
		res = append(res, string(svc))
	}

	return res
}
