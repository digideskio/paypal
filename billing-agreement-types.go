package paypal

import (
	"time"
)

type (
	BillingAgreement struct {
		ID          string       `json:"id,omitempty"`
		Name        string       `json:"name"`
		Description string       `json:"description"`
		StartDate   string       `json:"start_date"`
		Payer       *Payer       `json:"payer"`
		BillingPlan *BillingPlan `json:"plan"`
		// optional attributes
		ShippingAddress             *Address              `json:"shipping_address,omitempty"`
		OverrideMerchantPreferences *MerchantPreferences  `json:"override_merchant_preferences,omitempty"`
		OverrideChargeModels        []OverrideChargeModel `json:"override_charge_models,omitempty"`
		CreateTime                  *time.Time            `json:"create_time,omitempty"`
		UpdateTime                  *time.Time            `json:"update_time,omitempty"`
		Links                       []Links               `json:"links,omitempty"`
	}

	AgreementDetails struct {
		OutstandingBalance *Currency `json:"outstanding_balance" bson:"outstanding_balance"`
		CyclesRemaining    string    `json:"cycles_remaining" bson:"cycles_remaining"`
		CyclesCompleted    string    `json:"cycles_completed" bson:"cycles_completed"`
		NextBillingDate    string    `json:"next_billing_date" bson:"next_billing_date"`
		LastBilingDate     string    `json:"last_billing_date" bson:"last_billing_date"`
		LastPaymentAmount  *Currency `json:"last_payment_amount" bson:"last_payment_amount"`
		FinalPaymentDate   string    `json:"final_payment_date" bson:"final_payment_date"`
		FailedPaymentCount string    `json:"failed_payment_count" bson:"failed_payment_count"`
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
