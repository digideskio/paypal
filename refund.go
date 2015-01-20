// Package paypal defines types and operations used to access the Paypal API
//
// The following refund operations are defined for Payment:
//
// refund: lookup refund details
//
//      GET            /v1/payments/refund/<Refund-Id>
//
package paypal

import "fmt"

// GetRefund returns a refund by ID
func (c *Client) GetRefund(refundID string) (*Refund, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/refund/%s", c.APIBase, refundID), nil)
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
