// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVolumesClonePostReader is a Reader for the PcloudVolumesClonePost structure.
type PcloudVolumesClonePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVolumesClonePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudVolumesClonePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudVolumesClonePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPcloudVolumesClonePostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudVolumesClonePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudVolumesClonePostOK creates a PcloudVolumesClonePostOK with default headers values
func NewPcloudVolumesClonePostOK() *PcloudVolumesClonePostOK {
	return &PcloudVolumesClonePostOK{}
}

/*PcloudVolumesClonePostOK handles this case with default header values.

OK
*/
type PcloudVolumesClonePostOK struct {
	Payload *models.VolumesCloneResponse
}

func (o *PcloudVolumesClonePostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostOK  %+v", 200, o.Payload)
}

func (o *PcloudVolumesClonePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesCloneResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumesClonePostBadRequest creates a PcloudVolumesClonePostBadRequest with default headers values
func NewPcloudVolumesClonePostBadRequest() *PcloudVolumesClonePostBadRequest {
	return &PcloudVolumesClonePostBadRequest{}
}

/*PcloudVolumesClonePostBadRequest handles this case with default header values.

Bad Request
*/
type PcloudVolumesClonePostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudVolumesClonePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVolumesClonePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumesClonePostConflict creates a PcloudVolumesClonePostConflict with default headers values
func NewPcloudVolumesClonePostConflict() *PcloudVolumesClonePostConflict {
	return &PcloudVolumesClonePostConflict{}
}

/*PcloudVolumesClonePostConflict handles this case with default header values.

Conflict
*/
type PcloudVolumesClonePostConflict struct {
	Payload *models.Error
}

func (o *PcloudVolumesClonePostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostConflict  %+v", 409, o.Payload)
}

func (o *PcloudVolumesClonePostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumesClonePostInternalServerError creates a PcloudVolumesClonePostInternalServerError with default headers values
func NewPcloudVolumesClonePostInternalServerError() *PcloudVolumesClonePostInternalServerError {
	return &PcloudVolumesClonePostInternalServerError{}
}

/*PcloudVolumesClonePostInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudVolumesClonePostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudVolumesClonePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudVolumesClonePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVolumesClonePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
