// Package paypal defines types and operations used to access the Paypal API
//
// The following oreder operations are defined for Payment:
//
// order: manage orders
//
//      GET             /v1/payments/orders/<Order-Id>
//      POST            /v1/payments/orders/<Order-Id>/authorize
//      POST            /v1/payments/orders/<Order-Id>/do-void
//      POST            /v1/payments/orders/<Order-Id>/capture
//      POST            /v1/payments/capture/<Capture-Id>/refund
//
package paypal

import (
	"fmt"
)

type (
	OrderReq struct {
		Amount *Amount `json:"amount"`
	}
)

// GetOrder returns a sale by ID
func (c *Client) GetOrder(saleID string) (*Order, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/sale/%s", c.APIBase, saleID), nil)
	if err != nil {
		return nil, err
	}

	v := &Order{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
