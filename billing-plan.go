// Package paypal defines types and operations used to access the Paypal API
//
// The following billing-plan operations are defined for payments:
//
//      POST       /v1/payments/billing-plans
//      GET        /v1/payments/billing-plans
//      GET, PATCH /v1/payments/billing-plans/<Plan-Id>
package paypal

import "fmt"

type (
	CreateBillingPlanResp struct {
		*BillingPlan
		Links []Links `json:"links"`
	}

	ExecuteBillingPlanResp struct {
		Intent       PaymentIntent `json:"intent"`
		Payer        *Payer        `json:"payer"`
		Transactions []Transaction `json:"transactions"`
		Links        []Links       `json:"links"`
	}

	ListBillingPlansResp struct {
		BillingPlans []BillingPlan `json:"plans"`
	}
)

// CreateBillingPlan creates a billingplan in Paypal
func (c *Client) CreateBillingPlan(p *BillingPlan) (*CreateBillingPlanResp, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-plans", c.APIBase), p)
	if err != nil {
		return nil, err
	}

	v := &CreateBillingPlanResp{}

	err = c.SendAndAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// GetBillingPlan fetches a billingplan in Paypal
func (c *Client) GetBillingPlan(id string) (*BillingPlan, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/billing-plans/%s", c.APIBase, id), nil)
	if err != nil {
		return nil, err
	}

	v := &BillingPlan{}

	err = c.SendAndAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ListBillingPlans retrieve billingplans resources from Paypal
func (c *Client) ListBillingPlans(filter map[string]string) ([]BillingPlan, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/billing-plans", c.APIBase), nil)
	if err != nil {
		return nil, err
	}

	if filter != nil {
		q := req.URL.Query()

		for k, v := range filter {
			q.Set(k, v)
		}

		req.URL.RawQuery = q.Encode()
	}

	var v ListBillingPlansResp

	err = c.SendAndAuth(req, &v)
	if err != nil {
		return nil, err
	}

	return v.BillingPlans, nil
}

func (c *Client) UpdateBillingPlan(id string, patch []PatchRequest) error {
	req, err := NewRequest("PATCH", fmt.Sprintf("%s/payments/billing-plans/%s", c.APIBase, id), patch)
	if err != nil {
		return err
	}

	err = c.SendAndAuth(req, nil)
	if err != nil {
		return err
	}
	return nil
}
