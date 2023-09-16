// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// TimeDuration time duration
//
// swagger:model time.Duration
type TimeDuration int64

// for schema
var timeDurationEnum []interface{}

func init() {
	var res []TimeDuration
	if err := json.Unmarshal([]byte(`[-9223372036854776000,9223372036854776000,1,1000,1000000,1000000000,60000000000,3600000000000]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		timeDurationEnum = append(timeDurationEnum, v)
	}
}

func (m TimeDuration) validateTimeDurationEnum(path, location string, value TimeDuration) error {
	if err := validate.EnumCase(path, location, value, timeDurationEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this time duration
func (m TimeDuration) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTimeDurationEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this time duration based on context it is used
func (m TimeDuration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}