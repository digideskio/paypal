package paypal

type (

	// Payer maps to payer object
	Payer struct {
		PaymentMethod PaymentMethod `json:"payment_method"`
		// optional attributes
		FundingInstruments []FundingInstrument `json:"funding_instruments,omitempty"`
		PayerInfo          *PayerInfo          `json:"payer_info,omitempty"`
		Status             PayerStatus         `json:"payer_status,omitempty"` // { 'VERIFIED', 'UNVERIFIED' }
	}

	// PayerInfo maps to payer_info object
	PayerInfo struct {
		Email           string           `json:"email,omitempty"`
		FirstName       string           `json:"first_name,omitempty"`
		LastName        string           `json:"last_name,omitempty"`
		PayerID         string           `json:"payer_id,omitempty"`
		Phone           string           `json:"phone,omitempty"`
		ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
		TaxIDType       TaxIDType        `json:"tax_id_type,omitempty"`
		TaxID           string           `json:"tax_id,omitempty"`
	}
)
