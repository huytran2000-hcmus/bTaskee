// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new task API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for task API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIV1TasksID(params *GetAPIV1TasksIDParams, opts ...ClientOption) (*GetAPIV1TasksIDOK, error)

	PostAPIV1Tasks(params *PostAPIV1TasksParams, opts ...ClientOption) (*PostAPIV1TasksOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAPIV1TasksID gets a task by id

Get a task by id
*/
func (a *Client) GetAPIV1TasksID(params *GetAPIV1TasksIDParams, opts ...ClientOption) (*GetAPIV1TasksIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV1TasksIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV1TasksID",
		Method:             "GET",
		PathPattern:        "/api/v1/tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV1TasksIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIV1TasksIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV1TasksID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPIV1Tasks creates a new task

Create a new task
*/
func (a *Client) PostAPIV1Tasks(params *PostAPIV1TasksParams, opts ...ClientOption) (*PostAPIV1TasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV1TasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV1Tasks",
		Method:             "POST",
		PathPattern:        "/api/v1/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPIV1TasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAPIV1TasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIV1Tasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
