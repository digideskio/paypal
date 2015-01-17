package paypal

type (

	// Address maps to address object
	Address struct {
		Line1       string `json:"line1"`
		Line2       string `json:"line2,omitempty"`
		City        string `json:"city"`
		CountryCode string `json:"country_code"`
		PostalCode  string `json:"postal_code,omitempty"`
		State       string `json:"state,omitempty"`
		Phone       string `json:"phone,omitempty"`
	}

	// ShippingAddress maps to shipping_address object
	ShippingAddress struct {
		RecipientName string      `json:"recipient_name,omitempty"`
		Type          AddressType `json:"type,omitempty"`
		Line1         string      `json:"line1"`
		Line2         string      `json:"line2,omitempty"`
		City          string      `json:"city"`
		CountryCode   string      `json:"country_code"`
		PostalCode    string      `json:"postal_code,omitempty"`
		State         string      `json:"state,omitempty"`
		Phone         string      `json:"phone,omitempty"`
	}

	Phone struct {
		CountryCode    string `json:"country_code"` // E.164 format
		NationalNumber string `json:"national_number"`
	}
)
