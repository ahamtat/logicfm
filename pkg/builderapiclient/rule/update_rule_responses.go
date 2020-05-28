// Code generated by go-swagger; DO NOT EDIT.

package rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"../internal/models"
)

// UpdateRuleReader is a Reader for the UpdateRule structure.
type UpdateRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRuleOK creates a UpdateRuleOK with default headers values
func NewUpdateRuleOK() *UpdateRuleOK {
	return &UpdateRuleOK{}
}

/*UpdateRuleOK handles this case with default header values.

OK
*/
type UpdateRuleOK struct {
}

func (o *UpdateRuleOK) Error() string {
	return fmt.Sprintf("[PUT /rules/{id}][%d] updateRuleOK ", 200)
}

func (o *UpdateRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRuleDefault creates a UpdateRuleDefault with default headers values
func NewUpdateRuleDefault(code int) *UpdateRuleDefault {
	return &UpdateRuleDefault{
		_statusCode: code,
	}
}

/*UpdateRuleDefault handles this case with default header values.

error
*/
type UpdateRuleDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update rule default response
func (o *UpdateRuleDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /rules/{id}][%d] updateRule default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}