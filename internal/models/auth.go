// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Auth auth
//
// swagger:model Auth
type Auth struct {

	// auth id
	AuthID int64 `json:"auth_id,omitempty" gorm:"type:int primary key auto_increment"`

	// islogin
	Islogin bool `json:"islogin,omitempty" gorm:"type:tinyint(1);default:0"`

	// token
	Token string `json:"token,omitempty" gorm:"type:varchar(255)"`

	// user id
	UserID int64 `json:"user_id,omitempty" gorm:"foreignKey:UserID;references:UserID"`

	// username
	Username string `json:"username,omitempty" gorm:"type:varchar(255)"`
}

// Validate validates this auth
func (m *Auth) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this auth based on context it is used
func (m *Auth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Auth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Auth) UnmarshalBinary(b []byte) error {
	var res Auth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
