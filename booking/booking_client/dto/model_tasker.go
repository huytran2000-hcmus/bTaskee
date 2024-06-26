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

// ModelTasker model tasker
//
// swagger:model model.Tasker
type ModelTasker struct {

	// email
	// Example: \u003center your email here\u003e
	// Required: true
	Email *string `json:"email"`

	// id
	// Example: 1
	// Required: true
	ID *string `json:"id"`

	// identification
	// Required: true
	Identification *ModelIdentification `json:"identification"`

	// name
	// Example: huy tran
	// Required: true
	Name *string `json:"name"`

	// phone
	// Example: +84948447525
	// Required: true
	Phone *string `json:"phone"`
}

// Validate validates this model tasker
func (m *ModelTasker) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelTasker) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *ModelTasker) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ModelTasker) validateIdentification(formats strfmt.Registry) error {

	if err := validate.Required("identification", "body", m.Identification); err != nil {
		return err
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *ModelTasker) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ModelTasker) validatePhone(formats strfmt.Registry) error {

	if err := validate.Required("phone", "body", m.Phone); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this model tasker based on the context it is used
func (m *ModelTasker) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIdentification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelTasker) contextValidateIdentification(ctx context.Context, formats strfmt.Registry) error {

	if m.Identification != nil {

		if err := m.Identification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelTasker) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelTasker) UnmarshalBinary(b []byte) error {
	var res ModelTasker
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
