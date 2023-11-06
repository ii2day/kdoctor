// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new echo API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for echo API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Delete(params *DeleteParams, opts ...ClientOption) (*DeleteOK, error)

	DeleteKdoctoragent(params *DeleteKdoctoragentParams, opts ...ClientOption) (*DeleteKdoctoragentOK, error)

	Get(params *GetParams, opts ...ClientOption) (*GetOK, error)

	GetKdoctoragent(params *GetKdoctoragentParams, opts ...ClientOption) (*GetKdoctoragentOK, error)

	Head(params *HeadParams, opts ...ClientOption) (*HeadOK, error)

	HeadKdoctoragent(params *HeadKdoctoragentParams, opts ...ClientOption) (*HeadKdoctoragentOK, error)

	Options(params *OptionsParams, opts ...ClientOption) (*OptionsOK, error)

	OptionsKdoctoragent(params *OptionsKdoctoragentParams, opts ...ClientOption) (*OptionsKdoctoragentOK, error)

	Patch(params *PatchParams, opts ...ClientOption) (*PatchOK, error)

	PatchKdoctoragent(params *PatchKdoctoragentParams, opts ...ClientOption) (*PatchKdoctoragentOK, error)

	Post(params *PostParams, opts ...ClientOption) (*PostOK, error)

	PostKdoctoragent(params *PostKdoctoragentParams, opts ...ClientOption) (*PostKdoctoragentOK, error)

	Put(params *PutParams, opts ...ClientOption) (*PutOK, error)

	PutKdoctoragent(params *PutKdoctoragentParams, opts ...ClientOption) (*PutKdoctoragentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Delete cleans http request counts

clean http request counts
*/
func (a *Client) Delete(params *DeleteParams, opts ...ClientOption) (*DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Delete",
		Method:             "DELETE",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
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
	success, ok := result.(*DeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteKdoctoragent cleans http request counts

clean http request counts
*/
func (a *Client) DeleteKdoctoragent(params *DeleteKdoctoragentParams, opts ...ClientOption) (*DeleteKdoctoragentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKdoctoragentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteKdoctoragent",
		Method:             "DELETE",
		PathPattern:        "/kdoctoragent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteKdoctoragentReader{formats: a.formats},
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
	success, ok := result.(*DeleteKdoctoragentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteKdoctoragent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Get echos http request

echo http request
*/
func (a *Client) Get(params *GetParams, opts ...ClientOption) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
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
	success, ok := result.(*GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetKdoctoragent echos http request

echo http request
*/
func (a *Client) GetKdoctoragent(params *GetKdoctoragentParams, opts ...ClientOption) (*GetKdoctoragentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKdoctoragentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetKdoctoragent",
		Method:             "GET",
		PathPattern:        "/kdoctoragent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetKdoctoragentReader{formats: a.formats},
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
	success, ok := result.(*GetKdoctoragentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetKdoctoragent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Head echos http request

echo http request
*/
func (a *Client) Head(params *HeadParams, opts ...ClientOption) (*HeadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Head",
		Method:             "HEAD",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeadReader{formats: a.formats},
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
	success, ok := result.(*HeadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Head: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HeadKdoctoragent echos http request

echo http request
*/
func (a *Client) HeadKdoctoragent(params *HeadKdoctoragentParams, opts ...ClientOption) (*HeadKdoctoragentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadKdoctoragentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HeadKdoctoragent",
		Method:             "HEAD",
		PathPattern:        "/kdoctoragent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeadKdoctoragentReader{formats: a.formats},
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
	success, ok := result.(*HeadKdoctoragentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HeadKdoctoragent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Options echos http request

echo http request
*/
func (a *Client) Options(params *OptionsParams, opts ...ClientOption) (*OptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Options",
		Method:             "OPTIONS",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OptionsReader{formats: a.formats},
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
	success, ok := result.(*OptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Options: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OptionsKdoctoragent echos http request

echo http request
*/
func (a *Client) OptionsKdoctoragent(params *OptionsKdoctoragentParams, opts ...ClientOption) (*OptionsKdoctoragentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionsKdoctoragentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OptionsKdoctoragent",
		Method:             "OPTIONS",
		PathPattern:        "/kdoctoragent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OptionsKdoctoragentReader{formats: a.formats},
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
	success, ok := result.(*OptionsKdoctoragentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OptionsKdoctoragent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Patch echos http request

echo http request
*/
func (a *Client) Patch(params *PatchParams, opts ...ClientOption) (*PatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Patch",
		Method:             "PATCH",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchReader{formats: a.formats},
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
	success, ok := result.(*PatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Patch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchKdoctoragent echos http request

echo http request
*/
func (a *Client) PatchKdoctoragent(params *PatchKdoctoragentParams, opts ...ClientOption) (*PatchKdoctoragentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchKdoctoragentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchKdoctoragent",
		Method:             "PATCH",
		PathPattern:        "/kdoctoragent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchKdoctoragentReader{formats: a.formats},
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
	success, ok := result.(*PatchKdoctoragentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchKdoctoragent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Post echos http request counts

echo http request counts
*/
func (a *Client) Post(params *PostParams, opts ...ClientOption) (*PostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post",
		Method:             "POST",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostReader{formats: a.formats},
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
	success, ok := result.(*PostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostKdoctoragent echos http request counts

echo http request counts
*/
func (a *Client) PostKdoctoragent(params *PostKdoctoragentParams, opts ...ClientOption) (*PostKdoctoragentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostKdoctoragentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostKdoctoragent",
		Method:             "POST",
		PathPattern:        "/kdoctoragent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostKdoctoragentReader{formats: a.formats},
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
	success, ok := result.(*PostKdoctoragentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostKdoctoragent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Put echos http request

echo http request
*/
func (a *Client) Put(params *PutParams, opts ...ClientOption) (*PutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Put",
		Method:             "PUT",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutReader{formats: a.formats},
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
	success, ok := result.(*PutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutKdoctoragent echos http request

echo http request
*/
func (a *Client) PutKdoctoragent(params *PutKdoctoragentParams, opts ...ClientOption) (*PutKdoctoragentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutKdoctoragentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutKdoctoragent",
		Method:             "PUT",
		PathPattern:        "/kdoctoragent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutKdoctoragentReader{formats: a.formats},
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
	success, ok := result.(*PutKdoctoragentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutKdoctoragent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}