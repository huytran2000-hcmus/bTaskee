// Code generated by go-swagger; DO NOT EDIT.

package dto

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelCreateTaskRequest model create task request
//
// swagger:model model.CreateTaskRequest
type ModelCreateTaskRequest struct {

	// assigned location
	// Required: true
	AssignedLocation *ModelLocation `json:"assigned_location"`

	// customer
	// Required: true
	Customer *ModelCustomer `json:"customer"`

	// note
	// Example: Be here ontime
	Note string `json:"note,omitempty"`

	// tasker
	// Required: true
	Tasker *ModelTasker `json:"tasker"`

	// working details
	// Required: true
	WorkingDetails *ModelWorkingDetails `json:"working_details"`
}

// Validate validates this model create task request
func (m *ModelCreateTaskRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTasker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkingDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelCreateTaskRequest) validateAssignedLocation(formats strfmt.Registry) error {

	if err := validate.Required("assigned_location", "body", m.AssignedLocation); err != nil {
		return err
	}

	if m.AssignedLocation != nil {
		if err := m.AssignedLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assigned_location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assigned_location")
			}
			return err
		}
	}

	return nil
}

func (m *ModelCreateTaskRequest) validateCustomer(formats strfmt.Registry) error {

	if err := validate.Required("customer", "body", m.Customer); err != nil {
		return err
	}

	if m.Customer != nil {
		if err := m.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *ModelCreateTaskRequest) validateTasker(formats strfmt.Registry) error {

	if err := validate.Required("tasker", "body", m.Tasker); err != nil {
		return err
	}

	if m.Tasker != nil {
		if err := m.Tasker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tasker")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tasker")
			}
			return err
		}
	}

	return nil
}

func (m *ModelCreateTaskRequest) validateWorkingDetails(formats strfmt.Registry) error {

	if err := validate.Required("working_details", "body", m.WorkingDetails); err != nil {
		return err
	}

	if m.WorkingDetails != nil {
		if err := m.WorkingDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("working_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("working_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this model create task request based on the context it is used
func (m *ModelCreateTaskRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignedLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTasker(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkingDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelCreateTaskRequest) contextValidateAssignedLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.AssignedLocation != nil {

		if err := m.AssignedLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assigned_location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assigned_location")
			}
			return err
		}
	}

	return nil
}

func (m *ModelCreateTaskRequest) contextValidateCustomer(ctx context.Context, formats strfmt.Registry) error {

	if m.Customer != nil {

		if err := m.Customer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *ModelCreateTaskRequest) contextValidateTasker(ctx context.Context, formats strfmt.Registry) error {

	if m.Tasker != nil {

		if err := m.Tasker.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tasker")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tasker")
			}
			return err
		}
	}

	return nil
}

func (m *ModelCreateTaskRequest) contextValidateWorkingDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkingDetails != nil {

		if err := m.WorkingDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("working_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("working_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelCreateTaskRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelCreateTaskRequest) UnmarshalBinary(b []byte) error {
	var res ModelCreateTaskRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
