package features

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"swaggen/models"
)

// PutAPIFeaturesIDReader is a Reader for the PutAPIFeaturesID structure.
type PutAPIFeaturesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIFeaturesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAPIFeaturesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutAPIFeaturesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutAPIFeaturesIDOK creates a PutAPIFeaturesIDOK with default headers values
func NewPutAPIFeaturesIDOK() *PutAPIFeaturesIDOK {
	return &PutAPIFeaturesIDOK{}
}

/*PutAPIFeaturesIDOK handles this case with default header values.

Full feature representation
*/
type PutAPIFeaturesIDOK struct {
	Payload *models.FeatureFull
}

func (o *PutAPIFeaturesIDOK) Error() string {
	return fmt.Sprintf("[PUT /api/features/{id}][%d] putApiFeaturesIdOK  %+v", 200, o.Payload)
}

func (o *PutAPIFeaturesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FeatureFull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIFeaturesIDDefault creates a PutAPIFeaturesIDDefault with default headers values
func NewPutAPIFeaturesIDDefault(code int) *PutAPIFeaturesIDDefault {
	return &PutAPIFeaturesIDDefault{
		_statusCode: code,
	}
}

/*PutAPIFeaturesIDDefault handles this case with default header values.

Unexpected exception
*/
type PutAPIFeaturesIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put API features ID default response
func (o *PutAPIFeaturesIDDefault) Code() int {
	return o._statusCode
}

func (o *PutAPIFeaturesIDDefault) Error() string {
	return fmt.Sprintf("[PUT /api/features/{id}][%d] PutAPIFeaturesID default  %+v", o._statusCode, o.Payload)
}

func (o *PutAPIFeaturesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}