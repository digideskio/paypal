package paypal

import (
	"time"
)

type (

	// Order maps to order object
	Order struct {
		ID                        string        `json:"id,omitempty"`
		PurchaseUnitReferenceID   string        `json:"purchase_unit_reference_id,omitempty"`
		CreateTime                *time.Time    `json:"create_time,omitempty"`
		UpdateTime                *time.Time    `json:"update_time,omitempty"`
		Amount                    []Amount      `json:"amount,omitempty"`
		State                     OrderState    `json:"state,omitempty"`
		PendingReason             PendingReason `json:"pending_reason,omitempty"`
		ReasonCode                ReasonCode    `json:"reason_code,omitempty"`
		ClearingTime              string        `json:"clearing_time,omitempty"`
		ProtectionEligibility     string        `json:"protection_eligibility,omitempty"`
		ProtectionEligibilityType string        `json:"protection_eligiblity_type,omitempty"`
	}
)
