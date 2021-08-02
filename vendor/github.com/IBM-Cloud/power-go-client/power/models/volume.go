// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Volume volume
// swagger:model Volume
type Volume struct {

	// Indicates if the volume is the server's boot volume
	BootVolume *bool `json:"bootVolume,omitempty"`

	// Indicates if the volume is boot capable
	Bootable *bool `json:"bootable,omitempty"`

	// Creation Date
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// Indicates if the volume should be deleted when the server terminates
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty"`

	// Type of Disk
	DiskType string `json:"diskType,omitempty"`

	// Last Update Date
	// Required: true
	// Format: date-time
	LastUpdateDate *strfmt.DateTime `json:"lastUpdateDate"`

	// Volume Name
	// Required: true
	Name *string `json:"name"`

	// List of PCloud PVM Instance attached to the volume
	PvmInstanceIds []string `json:"pvmInstanceIDs"`

	// Indicates if the volume is shareable between VMs
	Shareable *bool `json:"shareable,omitempty"`

	// Volume Size
	// Required: true
	Size *float64 `json:"size"`

	// Volume State
	State string `json:"state,omitempty"`

	// Volume ID
	// Required: true
	VolumeID *string `json:"volumeID"`

	// Volume type, name of storage template used to create the volume
	VolumeType string `json:"volumeType,omitempty"`

	// Volume world wide name
	Wwn string `json:"wwn,omitempty"`
}

// Validate validates this volume
func (m *Volume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Volume) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateLastUpdateDate(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdateDate", "body", m.LastUpdateDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastUpdateDate", "body", "date-time", m.LastUpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("volumeID", "body", m.VolumeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Volume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Volume) UnmarshalBinary(b []byte) error {
	var res Volume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
