// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatasetUpdateRequest dataset update request
//
// swagger:model dataset.UpdateRequest
type DatasetUpdateRequest struct {

	// Public key of the encryption recipient
	EncryptionRecipients []string `json:"encryptionRecipients"`

	// EncryptionScript command to run for custom encryption
	EncryptionScript string `json:"encryptionScript,omitempty"`

	// Maximum size of the CAR files to be created
	MaxSize *string `json:"maxSize,omitempty"`

	// Output directory for CAR files. Do not set if using inline preparation
	OutputDirs []string `json:"outputDirs"`

	// Target piece size of the CAR files used for piece commitment calculation
	PieceSize string `json:"pieceSize,omitempty"`
}

// Validate validates this dataset update request
func (m *DatasetUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dataset update request based on context it is used
func (m *DatasetUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasetUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasetUpdateRequest) UnmarshalBinary(b []byte) error {
	var res DatasetUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
