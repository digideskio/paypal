// Package paypal defines types and operations used to access the Paypal API
//
// The following billing-agreement operations are defined for payments:
//
//      POST       /v1/payments/billing-agreements
//      POST       /v1/payments/billing-agreements/<Payment-Token>/agreement-execute
//      GET, PATCH /v1/payments/billing-agreements/<Agreement-Id>
//      POST       /v1/payments/billing-agreements/<Agreement-Id>/suspend
//      POST       /v1/payments/billing-agreements/<Agreement-Id>/re-activate
//      POST       /v1/payments/billing-agreements/<Agreement-Id>/cancel
//      POST       /v1/payments/billing-agreements/<Agreement-Id>/bill-balance
//      POST       /v1/payments/billing-agreements/<Agreement-Id>/set-balance
//      GET        /v1/payments/billing-agreements/<Agreement-Id>/transactions
package paypal

import "fmt"

type (
	CreateBillingAgreementResp struct {
		*BillingAgreement
		Links []Links `json:"links"`
	}

	ExecuteBillingAgreementResp struct {
		ID string 	`json:"id"`
		Links        []Links       `json:"links"`
	}

	ListBillingAgreementsResp struct {
		BillingAgreements []BillingAgreement `json:"plans"`
	}
)

// CreateBillingAgreement creates a billingagreement in Paypal
func (c *Client) CreateBillingAgreement(p *BillingAgreement) (*CreateBillingAgreementResp, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-agreements", c.APIBase), p)
	if err != nil {
		return nil, err
	}

	v := &CreateBillingAgreementResp{}

	err = c.SendAndAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ExecuteBillingAgreement completes an approved Paypal billingagreement that has been approved by the payer
func (c *Client) ExecuteBillingAgreement(planID, payerID string, transactions []Transaction) (*ExecuteBillingAgreementResp, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-agreements/%s/execute", c.APIBase, planID), struct {
		PayerID      string        `json:"payer_id"`
		Transactions []Transaction `json:"transactions"`
	}{
		payerID,
		transactions,
	})
	if err != nil {
		return nil, err
	}

	v := &ExecuteBillingAgreementResp{}

	err = c.SendAndAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// GetBillingAgreement fetches a billingagreement in Paypal
func (c *Client) GetBillingAgreement(id string) (*BillingAgreement, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/billing-agreements/%s", c.APIBase, id), nil)
	if err != nil {
		return nil, err
	}

	v := &BillingAgreement{}

	err = c.SendAndAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ListBillingAgreements retrieve billingagreements resources from Paypal
func (c *Client) ListBillingAgreements(filter map[string]string) ([]BillingAgreement, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/payments/billing-agreements/", c.APIBase), nil)
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

	var v ListBillingAgreementsResp

	err = c.SendAndAuth(req, &v)
	if err != nil {
		return nil, err
	}

	return v.BillingAgreements, nil
}
