// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_storage_capacity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudStoragecapacityTypesGetReader is a Reader for the PcloudStoragecapacityTypesGet structure.
type PcloudStoragecapacityTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudStoragecapacityTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudStoragecapacityTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPcloudStoragecapacityTypesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudStoragecapacityTypesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudStoragecapacityTypesGetOK creates a PcloudStoragecapacityTypesGetOK with default headers values
func NewPcloudStoragecapacityTypesGetOK() *PcloudStoragecapacityTypesGetOK {
	return &PcloudStoragecapacityTypesGetOK{}
}

/*PcloudStoragecapacityTypesGetOK handles this case with default header values.

OK
*/
type PcloudStoragecapacityTypesGetOK struct {
	Payload *models.StorageTypeCapacity
}

func (o *PcloudStoragecapacityTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageTypeCapacity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityTypesGetNotFound creates a PcloudStoragecapacityTypesGetNotFound with default headers values
func NewPcloudStoragecapacityTypesGetNotFound() *PcloudStoragecapacityTypesGetNotFound {
	return &PcloudStoragecapacityTypesGetNotFound{}
}

/*PcloudStoragecapacityTypesGetNotFound handles this case with default header values.

Not Found
*/
type PcloudStoragecapacityTypesGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudStoragecapacityTypesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityTypesGetInternalServerError creates a PcloudStoragecapacityTypesGetInternalServerError with default headers values
func NewPcloudStoragecapacityTypesGetInternalServerError() *PcloudStoragecapacityTypesGetInternalServerError {
	return &PcloudStoragecapacityTypesGetInternalServerError{}
}

/*PcloudStoragecapacityTypesGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudStoragecapacityTypesGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudStoragecapacityTypesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}