// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/huytran2000-hcmus/bTaskee/booking/booking_client/dto"
)

// GetAPIV1TasksIDReader is a Reader for the GetAPIV1TasksID structure.
type GetAPIV1TasksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV1TasksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV1TasksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAPIV1TasksIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/tasks/{id}] GetAPIV1TasksID", response, response.Code())
	}
}

// NewGetAPIV1TasksIDOK creates a GetAPIV1TasksIDOK with default headers values
func NewGetAPIV1TasksIDOK() *GetAPIV1TasksIDOK {
	return &GetAPIV1TasksIDOK{}
}

/*
GetAPIV1TasksIDOK describes a response with status code 200, with default header values.

success
*/
type GetAPIV1TasksIDOK struct {
	Payload *GetAPIV1TasksIDOKBody
}

// IsSuccess returns true when this get Api v1 tasks Id o k response has a 2xx status code
func (o *GetAPIV1TasksIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v1 tasks Id o k response has a 3xx status code
func (o *GetAPIV1TasksIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v1 tasks Id o k response has a 4xx status code
func (o *GetAPIV1TasksIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v1 tasks Id o k response has a 5xx status code
func (o *GetAPIV1TasksIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v1 tasks Id o k response a status code equal to that given
func (o *GetAPIV1TasksIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v1 tasks Id o k response
func (o *GetAPIV1TasksIDOK) Code() int {
	return 200
}

func (o *GetAPIV1TasksIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/tasks/{id}][%d] getApiV1TasksIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIV1TasksIDOK) String() string {
	return fmt.Sprintf("[GET /api/v1/tasks/{id}][%d] getApiV1TasksIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIV1TasksIDOK) GetPayload() *GetAPIV1TasksIDOKBody {
	return o.Payload
}

func (o *GetAPIV1TasksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAPIV1TasksIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV1TasksIDNotFound creates a GetAPIV1TasksIDNotFound with default headers values
func NewGetAPIV1TasksIDNotFound() *GetAPIV1TasksIDNotFound {
	return &GetAPIV1TasksIDNotFound{}
}

/*
GetAPIV1TasksIDNotFound describes a response with status code 404, with default header values.

bad request
*/
type GetAPIV1TasksIDNotFound struct {
	Payload *GetAPIV1TasksIDNotFoundBody
}

// IsSuccess returns true when this get Api v1 tasks Id not found response has a 2xx status code
func (o *GetAPIV1TasksIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api v1 tasks Id not found response has a 3xx status code
func (o *GetAPIV1TasksIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v1 tasks Id not found response has a 4xx status code
func (o *GetAPIV1TasksIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api v1 tasks Id not found response has a 5xx status code
func (o *GetAPIV1TasksIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v1 tasks Id not found response a status code equal to that given
func (o *GetAPIV1TasksIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Api v1 tasks Id not found response
func (o *GetAPIV1TasksIDNotFound) Code() int {
	return 404
}

func (o *GetAPIV1TasksIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/tasks/{id}][%d] getApiV1TasksIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV1TasksIDNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/tasks/{id}][%d] getApiV1TasksIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV1TasksIDNotFound) GetPayload() *GetAPIV1TasksIDNotFoundBody {
	return o.Payload
}

func (o *GetAPIV1TasksIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAPIV1TasksIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetAPIV1TasksIDNotFoundBody get API v1 tasks ID not found body
swagger:model GetAPIV1TasksIDNotFoundBody
*/
type GetAPIV1TasksIDNotFoundBody struct {
	dto.RestfulErrorResponse

	// data
	Data string `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetAPIV1TasksIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// GetAPIV1TasksIDNotFoundBodyAO0
	var getAPIV1TasksIDNotFoundBodyAO0 dto.RestfulErrorResponse
	if err := swag.ReadJSON(raw, &getAPIV1TasksIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.RestfulErrorResponse = getAPIV1TasksIDNotFoundBodyAO0

	// GetAPIV1TasksIDNotFoundBodyAO1
	var dataGetAPIV1TasksIDNotFoundBodyAO1 struct {
		Data string `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetAPIV1TasksIDNotFoundBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetAPIV1TasksIDNotFoundBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetAPIV1TasksIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getAPIV1TasksIDNotFoundBodyAO0, err := swag.WriteJSON(o.RestfulErrorResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getAPIV1TasksIDNotFoundBodyAO0)
	var dataGetAPIV1TasksIDNotFoundBodyAO1 struct {
		Data string `json:"data,omitempty"`
	}

	dataGetAPIV1TasksIDNotFoundBodyAO1.Data = o.Data

	jsonDataGetAPIV1TasksIDNotFoundBodyAO1, errGetAPIV1TasksIDNotFoundBodyAO1 := swag.WriteJSON(dataGetAPIV1TasksIDNotFoundBodyAO1)
	if errGetAPIV1TasksIDNotFoundBodyAO1 != nil {
		return nil, errGetAPIV1TasksIDNotFoundBodyAO1
	}
	_parts = append(_parts, jsonDataGetAPIV1TasksIDNotFoundBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get API v1 tasks ID not found body
func (o *GetAPIV1TasksIDNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with dto.RestfulErrorResponse
	if err := o.RestfulErrorResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this get API v1 tasks ID not found body based on the context it is used
func (o *GetAPIV1TasksIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with dto.RestfulErrorResponse
	if err := o.RestfulErrorResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *GetAPIV1TasksIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAPIV1TasksIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetAPIV1TasksIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetAPIV1TasksIDOKBody get API v1 tasks ID o k body
swagger:model GetAPIV1TasksIDOKBody
*/
type GetAPIV1TasksIDOKBody struct {
	dto.RestfulSuccessResponse

	// data
	Data *dto.ModelTask `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetAPIV1TasksIDOKBody) UnmarshalJSON(raw []byte) error {
	// GetAPIV1TasksIDOKBodyAO0
	var getAPIV1TasksIDOKBodyAO0 dto.RestfulSuccessResponse
	if err := swag.ReadJSON(raw, &getAPIV1TasksIDOKBodyAO0); err != nil {
		return err
	}
	o.RestfulSuccessResponse = getAPIV1TasksIDOKBodyAO0

	// GetAPIV1TasksIDOKBodyAO1
	var dataGetAPIV1TasksIDOKBodyAO1 struct {
		Data *dto.ModelTask `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetAPIV1TasksIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetAPIV1TasksIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetAPIV1TasksIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getAPIV1TasksIDOKBodyAO0, err := swag.WriteJSON(o.RestfulSuccessResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getAPIV1TasksIDOKBodyAO0)
	var dataGetAPIV1TasksIDOKBodyAO1 struct {
		Data *dto.ModelTask `json:"data,omitempty"`
	}

	dataGetAPIV1TasksIDOKBodyAO1.Data = o.Data

	jsonDataGetAPIV1TasksIDOKBodyAO1, errGetAPIV1TasksIDOKBodyAO1 := swag.WriteJSON(dataGetAPIV1TasksIDOKBodyAO1)
	if errGetAPIV1TasksIDOKBodyAO1 != nil {
		return nil, errGetAPIV1TasksIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetAPIV1TasksIDOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get API v1 tasks ID o k body
func (o *GetAPIV1TasksIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with dto.RestfulSuccessResponse
	if err := o.RestfulSuccessResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAPIV1TasksIDOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getApiV1TasksIdOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getApiV1TasksIdOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get API v1 tasks ID o k body based on the context it is used
func (o *GetAPIV1TasksIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with dto.RestfulSuccessResponse
	if err := o.RestfulSuccessResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAPIV1TasksIDOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getApiV1TasksIdOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getApiV1TasksIdOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAPIV1TasksIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAPIV1TasksIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetAPIV1TasksIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
