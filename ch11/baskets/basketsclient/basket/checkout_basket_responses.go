// Code generated by go-swagger; DO NOT EDIT.

package basket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"eda-in-golang/baskets/basketsclient/models"
)

// CheckoutBasketReader is a Reader for the CheckoutBasket structure.
type CheckoutBasketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckoutBasketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckoutBasketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCheckoutBasketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckoutBasketOK creates a CheckoutBasketOK with default headers values
func NewCheckoutBasketOK() *CheckoutBasketOK {
	return &CheckoutBasketOK{}
}

/* CheckoutBasketOK describes a response with status code 200, with default header values.

A successful response.
*/
type CheckoutBasketOK struct {
	Payload models.BasketspbCheckoutBasketResponse
}

func (o *CheckoutBasketOK) Error() string {
	return fmt.Sprintf("[PUT /api/baskets/{id}/checkout][%d] checkoutBasketOK  %+v", 200, o.Payload)
}
func (o *CheckoutBasketOK) GetPayload() models.BasketspbCheckoutBasketResponse {
	return o.Payload
}

func (o *CheckoutBasketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckoutBasketDefault creates a CheckoutBasketDefault with default headers values
func NewCheckoutBasketDefault(code int) *CheckoutBasketDefault {
	return &CheckoutBasketDefault{
		_statusCode: code,
	}
}

/* CheckoutBasketDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CheckoutBasketDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the checkout basket default response
func (o *CheckoutBasketDefault) Code() int {
	return o._statusCode
}

func (o *CheckoutBasketDefault) Error() string {
	return fmt.Sprintf("[PUT /api/baskets/{id}/checkout][%d] checkoutBasket default  %+v", o._statusCode, o.Payload)
}
func (o *CheckoutBasketDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CheckoutBasketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
