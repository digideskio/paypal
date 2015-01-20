// Package paypal defines types and operations used to access the Paypal API
//
// The following capture operations are defined for Payment:
//
// capture: lookup and refund captured payments
//
//      GET            /v1/payments/capture/<Capture-Id>
//      POST           /v1/payments/capture/<Capture-Id>/refund
//
package paypal

import "fmt"

// https://developer.paypal.com/webapps/developer/docs/api/#captures

// GetCapture returns details about a captured payment
func (c *Client) GetCapture(captureID string) (*Capture, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/capture/%s", c.APIBase, captureID), nil)
	if err != nil {
		return nil, err
	}

	v := &Capture{}

	err = c.SendAndAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// RefundCapture refund a captured payment. For partial refunds, a lower
// Amount object can be passed in.
func (c *Client) RefundCapture(captureID string, a *Amount) (*Refund, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/capture/%s/refund", c.APIBase, captureID), struct {
		Amount *Amount `json:"amount"`
	}{a})
	if err != nil {
		return nil, err
	}

	v := &Refund{}

	err = c.SendAndAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
