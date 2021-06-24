// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RetentionRule retention rule
//
// swagger:model RetentionRule
type RetentionRule struct {

	// action
	Action string `json:"action,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// params
	Params map[string]interface{} `json:"params,omitempty"`

	// priority
	Priority int64 `json:"priority,omitempty"`

	// scope selectors
	ScopeSelectors map[string][]RetentionSelector `json:"scope_selectors,omitempty"`

	// tag selectors
	TagSelectors []*RetentionSelector `json:"tag_selectors"`

	// template
	Template string `json:"template,omitempty"`
}

// Validate validates this retention rule
func (m *RetentionRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScopeSelectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagSelectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetentionRule) validateScopeSelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.ScopeSelectors) { // not required
		return nil
	}

	for k := range m.ScopeSelectors {

		if err := validate.Required("scope_selectors"+"."+k, "body", m.ScopeSelectors[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.ScopeSelectors[k]); i++ {

			if err := m.ScopeSelectors[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scope_selectors" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *RetentionRule) validateTagSelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.TagSelectors) { // not required
		return nil
	}

	for i := 0; i < len(m.TagSelectors); i++ {
		if swag.IsZero(m.TagSelectors[i]) { // not required
			continue
		}

		if m.TagSelectors[i] != nil {
			if err := m.TagSelectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tag_selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RetentionRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionRule) UnmarshalBinary(b []byte) error {
	var res RetentionRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
