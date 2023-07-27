// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DatasourceOnedriveRequest datasource onedrive request
//
// swagger:model datasource.OnedriveRequest
type DatasourceOnedriveRequest struct {

	// Set scopes to be requested by rclone.
	AccessScopes *string `json:"accessScopes,omitempty"`

	// Auth server URL.
	AuthURL string `json:"authUrl,omitempty"`

	// Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
	ChunkSize *string `json:"chunkSize,omitempty"`

	// OAuth Client Id.
	ClientID string `json:"clientId,omitempty"`

	// OAuth Client Secret.
	ClientSecret string `json:"clientSecret,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// Disable the request for Sites.Read.All permission.
	DisableSitePermission *string `json:"disableSitePermission,omitempty"`

	// The ID of the drive to use.
	DriveID string `json:"driveId,omitempty"`

	// The type of the drive (personal | business | documentLibrary).
	DriveType string `json:"driveType,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Set to make OneNote files show up in directory listings.
	ExposeOnenoteFiles *string `json:"exposeOnenoteFiles,omitempty"`

	// Specify the hash in use for the backend.
	HashType *string `json:"hashType,omitempty"`

	// Set the password for links created by the link command.
	LinkPassword string `json:"linkPassword,omitempty"`

	// Set the scope of the links created by the link command.
	LinkScope *string `json:"linkScope,omitempty"`

	// Set the type of the links created by the link command.
	LinkType *string `json:"linkType,omitempty"`

	// Size of listing chunk.
	ListChunk *string `json:"listChunk,omitempty"`

	// Remove all versions on modifying operations.
	NoVersions *string `json:"noVersions,omitempty"`

	// Choose national cloud region for OneDrive.
	Region *string `json:"region,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// ID of the root folder.
	RootFolderID string `json:"rootFolderId,omitempty"`

	// Allow server-side operations (e.g. copy) to work across different onedrive configs.
	ServerSideAcrossConfigs *string `json:"serverSideAcrossConfigs,omitempty"`

	// The path of the source to scan items
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// OAuth Access Token as a JSON blob.
	Token string `json:"token,omitempty"`

	// Token server url.
	TokenURL string `json:"tokenUrl,omitempty"`
}

// Validate validates this datasource onedrive request
func (m *DatasourceOnedriveRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleteAfterExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRescanInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceOnedriveRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceOnedriveRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceOnedriveRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this datasource onedrive request based on context it is used
func (m *DatasourceOnedriveRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceOnedriveRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceOnedriveRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceOnedriveRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
