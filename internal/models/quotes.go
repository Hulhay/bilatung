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

// Quotes quotes
//
// swagger:model Quotes
type Quotes struct {

	// author
	// Required: true
	Author *string `json:"author" gorm:"type:varchar(255)"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamp;autoCreateTime"`

	// downvote
	Downvote int64 `json:"downvote,omitempty" gorm:"type:int(255);default:0"`

	// quote
	// Required: true
	// Max Length: 100
	Quote *string `json:"quote" gorm:"type:varchar(255)"`

	// quote id
	QuoteID int64 `json:"quote_id,omitempty" gorm:"type:int primary key auto_increment"`

	// tags
	// Required: true
	Tags []string `json:"tags" gorm:"type:varchar(255)"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamp;autoUpdateTime"`

	// upvote
	Upvote int64 `json:"upvote,omitempty" gorm:"type:int(255);default:0"`

	// user id
	UserID int64 `json:"user_id,omitempty" gorm:"foreignKey:UserID;references:UserID"`
}

// Validate validates this quotes
func (m *Quotes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Quotes) validateAuthor(formats strfmt.Registry) error {

	if err := validate.Required("author", "body", m.Author); err != nil {
		return err
	}

	return nil
}

func (m *Quotes) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quotes) validateQuote(formats strfmt.Registry) error {

	if err := validate.Required("quote", "body", m.Quote); err != nil {
		return err
	}

	if err := validate.MaxLength("quote", "body", *m.Quote, 100); err != nil {
		return err
	}

	return nil
}

func (m *Quotes) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *Quotes) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this quotes based on context it is used
func (m *Quotes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Quotes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Quotes) UnmarshalBinary(b []byte) error {
	var res Quotes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}