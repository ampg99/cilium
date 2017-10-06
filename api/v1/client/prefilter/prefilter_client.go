// Code generated by go-swagger; DO NOT EDIT.

package prefilter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new prefilter API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for prefilter API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeletePrefilter deletes list of c ID rs
*/
func (a *Client) DeletePrefilter(params *DeletePrefilterParams) (*DeletePrefilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrefilterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePrefilter",
		Method:             "DELETE",
		PathPattern:        "/prefilter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePrefilterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePrefilterOK), nil

}

/*
GetPrefilter retrieves list of c ID rs
*/
func (a *Client) GetPrefilter(params *GetPrefilterParams) (*GetPrefilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrefilterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrefilter",
		Method:             "GET",
		PathPattern:        "/prefilter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrefilterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrefilterOK), nil

}

/*
PutPrefilter updates list of c ID rs
*/
func (a *Client) PutPrefilter(params *PutPrefilterParams) (*PutPrefilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutPrefilterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutPrefilter",
		Method:             "PUT",
		PathPattern:        "/prefilter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutPrefilterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutPrefilterOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}