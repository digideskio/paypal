package paypal

import (
	"time"
)

type (
	BillingPlan struct {
		ID          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Type        string `json:"type,omitempty"` // { 'FIXED', 'INFINITE' }
		// optional attributes
		State               string               `json:"state,omitempty"` // { 'CREATED', 'ACTIVE', 'INACTIVE', 'DELETED' }
		PaymentDefinitions  []*PaymentDefinition `json:"payment_definitions,omitempty"`
		Terms               []Terms              `json:"terms,omitempty"`
		MerchantPreferences *MerchantPreferences `json:"merchant_preferences,omitempty"`
		CreateTime          *time.Time           `json:"createTime,omitempty"` // YYYY-MM-DDTimeTimezone, as defined in [ISO8601](http://tools.ietf.org/html/rfc3339#section-5.6).
		UpdateTime          *time.Time           `json:"updateTime,omitempty"` // YYYY-MM-DDTimeTimezone, as defined in [ISO8601](http://tools.ietf.org/html/rfc3339#section-5.6).
		Links               []*Links             `json:"links,omitempty"`
	}

	BillingPlanList struct {
		BillingPlans []*BillingPlan `json:"plans"`
		TotalItems   string         `json:"total_items"`
		TotalPages   string         `json:"total_pages"`
		Links        []*Links       `json:"links"`
	}

	MerchantPreferences struct {
		ID        string `json:"id,omitempty"`
		ReturnURL string `json:"return_url"`
		CancelURL string `json:"cancel_url"`
		// optional attributes
		SetupFee                *Currency `json:"setup_fee,omitempty"`
		NotifyURL               string  `json:"notify_url,omitempty"`
		MaxFailAttempts         string  `json:"max_fail_attempts,omitempty"`
		AutoBillAmount          string  `json:"auto_bill_amount,omitempty"`           // { 'YES', 'NO' }
		InitialFailAmountAction string  `json:"initial_fail_amount_action,omitempty"` // { 'CONTINUE' }
		AcceptedPaymentType     string  `json:"accepted_payment_type,omitempty"`
		CharSet                 string  `json:"charset,omitempty"`
	}

	PaymentDefinition struct {
		ID                string `json:"id,omitempty"`
		Name              string `json:"name"`
		Type              string `json:"type"`      // { 'REGULAR', 'TRIAL' }
		Frequency         string `json:"frequency"` // { 'DAY', 'WEEK', 'MONTH', 'YEAR' }
		Amount            *Currency `json:"amount"`
		Cycles            string `json:"cycles"`
		FrequencyInterval string `json:"frequency_interval"`
		// optional attributes
		ChargeModels []*ChargeModel `json:"charge_models,omitempty"`
	}

	ChargeModel struct {
		ID     string  `json:"id,omitempty"`
		Type   string  `json:"type"` //  {'SHIPPING', 'TAX' }
		Amount *Currency `json:"amount"`
	}

	OverrideChargeModel struct {
		ChargeModelID string `json:"charge_id"`
		Amount        Amount `json:"amount"`
	}

	Terms struct {
		ID               string   `json:"id,omitempty"`
		Type             string   `json:"type"` // { 'WEEKLY', 'MONTHLY', 'YEARLY' }
		MaxBillingAmount Currency `json:"max_billing_amount"`
		Occurrences      string   `json:"occurrences"`
		AmountRange      Currency `json:"amount_range"`
		BuyerEditable    string   `json:"buyer_editable"`
	}
)
