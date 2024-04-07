package brevo

import (
	"context"
	"strings"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/huytran2000-hcmus/bTaskee/send/config"
	"github.com/huytran2000-hcmus/bTaskee/send/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/send/internal/repository"
)

var _ repository.Emai = (*Brevo)(nil)

type Brevo struct {
	client *brevo.APIClient
}

func New() *Brevo {
	cfg := brevo.NewConfiguration()
	cfg.AddDefaultHeader("api-key", config.LoadEnv().BrevoAPIKey)

	br := brevo.NewAPIClient(cfg)

	return &Brevo{
		client: br,
	}
}

func (b *Brevo) SendEmail(ctx context.Context, req model.SendTaskRequest) ([]string, error) {
	email, _, err := b.client.TransactionalEmailsApi.SendTransacEmail(ctx, brevo.SendSmtpEmail{
		To: []brevo.SendSmtpEmailTo{
			{
				Email: req.Tasker.Email,
				Name:  req.Tasker.Name,
			},
		},
		TemplateId: config.LoadEnv().BrevoTaskAssignedEmailTemplateID,
		Params:     fromSendTaskRequestToParams(req),
	})
	if err != nil {
		return nil, err
	}

	if len(email.MessageIds) != 0 {
		return email.MessageIds, nil
	}

	return []string{email.MessageId}, nil
}

func fromSendTaskRequestToParams(req model.SendTaskRequest) map[string]any {
	params := map[string]any{}

	params["id"] = req.ID

	params["customer"] = map[string]string{
		"name":  req.Customer.Name,
		"email": req.Customer.Email,
		"phone": req.Customer.Phone,
	}

	params["address"] = req.AssignedLocation.Address

	params["from_time"] = req.WorkingDetails.FromTime
	params["to_time"] = req.WorkingDetails.ToTime

	params["house_type"] = req.WorkingDetails.HouseType
	params["service_types"] = strings.Join(req.WorkingDetails.ServiceTypes, ",")

	params["note"] = req.Note
	params["total"] = req.Total

	return params
}
