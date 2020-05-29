// Code generated by go-swagger; DO NOT EDIT.

package rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ahamtat/logicfm/pkg/models"
)

// AddNewReader is a Reader for the AddNew structure.
type AddNewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddNewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddNewCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddNewDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddNewCreated creates a AddNewCreated with default headers values
func NewAddNewCreated() *AddNewCreated {
	return &AddNewCreated{}
}

/*AddNewCreated handles this case with default header values.

Created
*/
type AddNewCreated struct {
}

func (o *AddNewCreated) Error() string {
	return fmt.Sprintf("[POST /rule/add/{musrvId}][%d] addNewCreated ", 201)
}

func (o *AddNewCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddNewDefault creates a AddNewDefault with default headers values
func NewAddNewDefault(code int) *AddNewDefault {
	return &AddNewDefault{
		_statusCode: code,
	}
}

/*AddNewDefault handles this case with default header values.

error
*/
type AddNewDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add new default response
func (o *AddNewDefault) Code() int {
	return o._statusCode
}

func (o *AddNewDefault) Error() string {
	return fmt.Sprintf("[POST /rule/add/{musrvId}][%d] addNew default  %+v", o._statusCode, o.Payload)
}

func (o *AddNewDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddNewDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
