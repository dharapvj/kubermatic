// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// CreateMeteringReportConfigurationReader is a Reader for the CreateMeteringReportConfiguration structure.
type CreateMeteringReportConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMeteringReportConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMeteringReportConfigurationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateMeteringReportConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateMeteringReportConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateMeteringReportConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMeteringReportConfigurationCreated creates a CreateMeteringReportConfigurationCreated with default headers values
func NewCreateMeteringReportConfigurationCreated() *CreateMeteringReportConfigurationCreated {
	return &CreateMeteringReportConfigurationCreated{}
}

/* CreateMeteringReportConfigurationCreated describes a response with status code 201, with default header values.

MeteringReportConfiguration
*/
type CreateMeteringReportConfigurationCreated struct {
	Payload *models.MeteringReportConfiguration
}

func (o *CreateMeteringReportConfigurationCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/metering/configurations/reports/{name}][%d] createMeteringReportConfigurationCreated  %+v", 201, o.Payload)
}
func (o *CreateMeteringReportConfigurationCreated) GetPayload() *models.MeteringReportConfiguration {
	return o.Payload
}

func (o *CreateMeteringReportConfigurationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MeteringReportConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMeteringReportConfigurationUnauthorized creates a CreateMeteringReportConfigurationUnauthorized with default headers values
func NewCreateMeteringReportConfigurationUnauthorized() *CreateMeteringReportConfigurationUnauthorized {
	return &CreateMeteringReportConfigurationUnauthorized{}
}

/* CreateMeteringReportConfigurationUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateMeteringReportConfigurationUnauthorized struct {
}

func (o *CreateMeteringReportConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/metering/configurations/reports/{name}][%d] createMeteringReportConfigurationUnauthorized ", 401)
}

func (o *CreateMeteringReportConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMeteringReportConfigurationForbidden creates a CreateMeteringReportConfigurationForbidden with default headers values
func NewCreateMeteringReportConfigurationForbidden() *CreateMeteringReportConfigurationForbidden {
	return &CreateMeteringReportConfigurationForbidden{}
}

/* CreateMeteringReportConfigurationForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateMeteringReportConfigurationForbidden struct {
}

func (o *CreateMeteringReportConfigurationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/metering/configurations/reports/{name}][%d] createMeteringReportConfigurationForbidden ", 403)
}

func (o *CreateMeteringReportConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMeteringReportConfigurationDefault creates a CreateMeteringReportConfigurationDefault with default headers values
func NewCreateMeteringReportConfigurationDefault(code int) *CreateMeteringReportConfigurationDefault {
	return &CreateMeteringReportConfigurationDefault{
		_statusCode: code,
	}
}

/* CreateMeteringReportConfigurationDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateMeteringReportConfigurationDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create metering report configuration default response
func (o *CreateMeteringReportConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *CreateMeteringReportConfigurationDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/metering/configurations/reports/{name}][%d] createMeteringReportConfiguration default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMeteringReportConfigurationDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateMeteringReportConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateMeteringReportConfigurationBody create metering report configuration body
swagger:model CreateMeteringReportConfigurationBody
*/
type CreateMeteringReportConfigurationBody struct {

	// interval
	Interval int32 `json:"interval,omitempty"`

	// retention
	Retention int32 `json:"retention,omitempty"`

	// schedule
	Schedule string `json:"schedule,omitempty"`

	// types
	Types []string `json:"types"`
}

// Validate validates this create metering report configuration body
func (o *CreateMeteringReportConfigurationBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create metering report configuration body based on context it is used
func (o *CreateMeteringReportConfigurationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateMeteringReportConfigurationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateMeteringReportConfigurationBody) UnmarshalBinary(b []byte) error {
	var res CreateMeteringReportConfigurationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
