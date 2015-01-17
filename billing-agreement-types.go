package paypal

import (
	"time"
)

type (
	BillingAgreement struct {
		ID          string      `json:"idion_type,omitempty"`
		Name        string      `json:"name"`
		Description string      `json:"description"`
		StartDate   string      `json:"start_date"`
		Payer       Payer       `json:"payer"`
		BillingPlan BillingPlan `json:"plan"`
		// optional attributes
		ShippingAddress             *Address              `json:"shipping_address,omitempty"`
		OverrideMerchantPreferences *MerchantPreferences  `json:"override_merchant_preferences,omitempty"`
		OverrideChargeModels        []OverrideChargeModel `json:"override_charge_models,omitempty"`
		CreateTime                  *time.Time            `json:"create_time,omitempty"`
		UpdateTime                  *time.Time            `json:"update_time,omitempty"`
		Links                       []Links               `json:"links,omitempty"`
	}

	AgreementTransaction struct {
		ID        string `json:"transaction_id"`
		Amount    Amount `json:"amount"`
		FeeAmount Amount `json:"amount"`
		NetAmount Amount `json:"amount"`
		// optional attributes
		Status      string     `json:"status,omitempty"`
		Type        string     `json:"transaction_type,omitempty"`
		PayerEmail  string     `json:"payer_email,omitempty"`
		PayerName   string     `json:"payer_name,omitempty"`
		TimeUpdated *time.Time `json:"time_updated,omitempty"`
		Timezone    string     `json:"time_zone,omitempty"`
	}

	AgreementStateDescriptor struct {
		Note   string `json:"note,omitempty"`
		Amount Amount `json:"amount,omitempty"`
	}

	AgreementTransactionList struct {
		AgreementTransactionList []*AgreementTransaction `json:"agreement_transaction_list,omitempty"`
	}
)
