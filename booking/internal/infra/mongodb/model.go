package mongodb

import (
	"time"

	"github.com/google/uuid"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
)

type Task struct {
	ID               uuid.UUID      `bson:"_id"`
	Customer         Customer       `bson:"customer"`
	AssignedLocation Location       `bson:"assigned_location"`
	WorkingDetails   WorkingDetails `bson:"working_time"`
	Note             string         `bson:"note"`
	Tasker           Tasker         `bson:"tasker"`
	Total            float64        `bson:"total"`
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
	Name  string `bson:"customer_name"`
	Email string `bson:"email"`
	Phone string `bson:"phone"`
}

type Location struct {
	Address string `bson:"address"`
}

type WorkingDetails struct {
	HouseType    string    `bson:"house_type"`
	ServiceTypes []string  `bson:"service_types"`
	From         time.Time `bson:"from"`
	To           time.Time `bson:"to"`
}

func toMongoDBTask(task *model.Task) *Task {
	return &Task{
		ID: task.ID,
		Customer: Customer{
			Name:  task.Customer.Name,
			Email: task.Customer.Email,
			Phone: task.Customer.Phone,
		},
		AssignedLocation: Location{
			Address: task.AssignedLocation.Address,
		},
		WorkingDetails: WorkingDetails{
			HouseType:    string(task.WorkingDetails.HouseType),
			ServiceTypes: toServiceTypeString(task.WorkingDetails.ServiceTypes),
			From:         task.WorkingDetails.FromTime,
			To:           task.WorkingDetails.ToTime,
		},
		Tasker: Tasker{
			ID:    task.Tasker.ID,
			Name:  task.Tasker.Name,
			Email: task.Tasker.Email,
			Phone: task.Tasker.Phone,
			Identification: Identification{
				CMND: task.Tasker.Identification.CMND,
				CCCD: task.Tasker.Identification.CCCD,
			},
		},
		Note:  task.Note,
		Total: task.Total,
	}
}

func toModelTask(task *Task) *model.Task {
	return &model.Task{
		ID: task.ID,
		Customer: model.Customer{
			Name:  task.Customer.Name,
			Email: task.Customer.Email,
			Phone: task.Customer.Phone,
		},
		AssignedLocation: model.Location{
			Address: task.AssignedLocation.Address,
		},
		WorkingDetails: model.WorkingDetails{
			HouseType:    model.HouseType(task.WorkingDetails.HouseType),
			ServiceTypes: toModelServiceType(task.WorkingDetails.ServiceTypes),
			FromTime:     task.WorkingDetails.From,
			ToTime:       task.WorkingDetails.To,
		},
		Tasker: model.Tasker{
			ID:    task.Tasker.ID,
			Name:  task.Tasker.Name,
			Email: task.Tasker.Email,
			Phone: task.Tasker.Phone,
			Identification: model.Identification{
				CMND: task.Tasker.Identification.CMND,
				CCCD: task.Tasker.Identification.CCCD,
			},
		},
		Note:  task.Note,
		Total: task.Total,
	}
}

func toServiceTypeString(typs []model.ServiceType) []string {
	res := make([]string, 0, len(typs))
	for _, typ := range typs {
		res = append(res, string(typ))
	}

	return res
}

func toModelServiceType(typs []string) []model.ServiceType {
	res := make([]model.ServiceType, 0, len(typs))
	for _, typ := range typs {
		res = append(res, model.ServiceType(typ))
	}

	return res
}
