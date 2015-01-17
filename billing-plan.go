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
func (c *Client) CreateBillingPlan(p BillingPlan) (*CreateBillingPlanResp, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/plans/plan", c.APIBase), p)
	if err != nil {
		return nil, err
	}

	v := &CreateBillingPlanResp{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ExecuteBillingPlan completes an approved Paypal billingplan that has been approved by the payer
func (c *Client) ExecuteBillingPlan(planID, payerID string, transactions []Transaction) (*ExecuteBillingPlanResp, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/plans/plan/%s/execute", c.APIBase, planID), struct {
		PayerID      string        `json:"payer_id"`
		Transactions []Transaction `json:"transactions"`
	}{
		payerID,
		transactions,
	})
	if err != nil {
		return nil, err
	}

	v := &ExecuteBillingPlanResp{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// GetBillingPlan fetches a billingplan in Paypal
func (c *Client) GetBillingPlan(id string) (*BillingPlan, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/plans/plan/%s", c.APIBase, id), nil)
	if err != nil {
		return nil, err
	}

	v := &BillingPlan{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ListBillingPlans retrieve billingplans resources from Paypal
func (c *Client) ListBillingPlans(filter map[string]string) ([]BillingPlan, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/plans/plan/", c.APIBase), nil)
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

	err = c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err
	}

	return v.BillingPlans, nil
}
