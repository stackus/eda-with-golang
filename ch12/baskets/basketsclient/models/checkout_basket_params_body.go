// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CheckoutBasketParamsBody checkout basket params body
//
// swagger:model checkoutBasketParamsBody
type CheckoutBasketParamsBody struct {

	// payment Id
	PaymentID string `json:"paymentId,omitempty"`
}

// Validate validates this checkout basket params body
func (m *CheckoutBasketParamsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checkout basket params body based on context it is used
func (m *CheckoutBasketParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CheckoutBasketParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckoutBasketParamsBody) UnmarshalBinary(b []byte) error {
	var res CheckoutBasketParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
