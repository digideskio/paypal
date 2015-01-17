package paypal

type (

	// Amount maps to the amount object
	Amount struct {
		Currency string   `json:"currency"`
		Total    string   `json:"total"`
		Details  *Details `json:"details,omitempty"`
	}

	// Currency maps to Currency object
	Currency struct {
		Currency string `json:"currency"` // 3 letter currency code as defined in ISO 4217
		Value    string `json:"value"`
	}

	// Links maps to links object
	Links struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
		// TODO: Support HyperSchema with its multiple types per field
		// TargetSchema HyperSchema `json:"targetSchema"`
		Method  string `json:"method"`
		Enctype string `json:"enctype"`
		// Schema HyperSchema `json:"schema"`
	}
)
