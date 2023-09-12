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

// PutKdoctoragentReader is a Reader for the PutKdoctoragent structure.
type PutKdoctoragentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutKdoctoragentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutKdoctoragentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPutKdoctoragentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutKdoctoragentOK creates a PutKdoctoragentOK with default headers values
func NewPutKdoctoragentOK() *PutKdoctoragentOK {
	return &PutKdoctoragentOK{}
}

/*
PutKdoctoragentOK describes a response with status code 200, with default header values.

Success
*/
type PutKdoctoragentOK struct {
}

// IsSuccess returns true when this put kdoctoragent o k response has a 2xx status code
func (o *PutKdoctoragentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put kdoctoragent o k response has a 3xx status code
func (o *PutKdoctoragentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put kdoctoragent o k response has a 4xx status code
func (o *PutKdoctoragentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put kdoctoragent o k response has a 5xx status code
func (o *PutKdoctoragentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put kdoctoragent o k response a status code equal to that given
func (o *PutKdoctoragentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put kdoctoragent o k response
func (o *PutKdoctoragentOK) Code() int {
	return 200
}

func (o *PutKdoctoragentOK) Error() string {
	return fmt.Sprintf("[PUT /kdoctoragent][%d] putKdoctoragentOK ", 200)
}

func (o *PutKdoctoragentOK) String() string {
	return fmt.Sprintf("[PUT /kdoctoragent][%d] putKdoctoragentOK ", 200)
}

func (o *PutKdoctoragentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutKdoctoragentInternalServerError creates a PutKdoctoragentInternalServerError with default headers values
func NewPutKdoctoragentInternalServerError() *PutKdoctoragentInternalServerError {
	return &PutKdoctoragentInternalServerError{}
}

/*
PutKdoctoragentInternalServerError describes a response with status code 500, with default header values.

Failed
*/
type PutKdoctoragentInternalServerError struct {
}

// IsSuccess returns true when this put kdoctoragent internal server error response has a 2xx status code
func (o *PutKdoctoragentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put kdoctoragent internal server error response has a 3xx status code
func (o *PutKdoctoragentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put kdoctoragent internal server error response has a 4xx status code
func (o *PutKdoctoragentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put kdoctoragent internal server error response has a 5xx status code
func (o *PutKdoctoragentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put kdoctoragent internal server error response a status code equal to that given
func (o *PutKdoctoragentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put kdoctoragent internal server error response
func (o *PutKdoctoragentInternalServerError) Code() int {
	return 500
}

func (o *PutKdoctoragentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /kdoctoragent][%d] putKdoctoragentInternalServerError ", 500)
}

func (o *PutKdoctoragentInternalServerError) String() string {
	return fmt.Sprintf("[PUT /kdoctoragent][%d] putKdoctoragentInternalServerError ", 500)
}

func (o *PutKdoctoragentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}