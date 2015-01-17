package paypal

type (
	Invoice struct {
		Number       string        `json:"number,omitempty"` // auto-inremented from last number if empty
		MerchantInfo *MerchantInfo `json:"merchant_info"`
		// optional attributes
		BillingInfo                []BillingInfo `json:"billing_info,omitempty"`
		ShippingInfo               *ShippingInfo `json:"shipping_info,omitempty"`
		InvoiceItems               []InvoiceItem `json:"items,omitempty"`
		InvoiceDate                string        `json:"invoice_date,omitempty"`
		PaymentTerm                *PaymentTerm  `json:"payment_term,omitempty"`
		Discount                   *Cost         `json:"discount,omitempty"`
		ShippingCost               *ShippingCost `json:"shipping_cost,omitempty"`
		Custom                     *CustomAmount `json:"custom_amount,omitempty"`
		TaxCalculatedAfterDiscount bool          `json:"tax_calculated_after_discount,omitempty"`
		TaxInclusive               bool          `json:"tax_inclusive,omitempty"`
		Terms                      string        `json:"terms,omitempty"`
		MerchantMemo               string        `json:"merchant_memo,omitempty"`
		LogoURL                    string        `json:"logo_url,omitempty"`
	}

	BillingInfo struct {
		Email          string   `json:"email"`
		FirstName      string   `json:"first_name,omitempty"`
		LastName       string   `json:"last_name,omitempty"`
		BusinessName   string   `json:"business_name,omitempty"`
		Address        *Address `json:"address,omitempty"`
		Language       string   `json:"language,omitempty"` // { da_DK, de_DE, en_AU, en_GB, en_US, es_ES, es_XC, fr_CA, fr_FR, fr_XC, he_IL, id_ID, it_IT, ja_JP, nl_NL, no_NO, pl_PL, pt_BR, pt_PT, ru_RU, sv_SE, th_TH, tr_TR, zh_CN, zh_HK, zh_TW, zh_XC }
		AdditionalInfo string   `json:"additional_info,omitempty"`
	}

	Cost struct {
		Percent int     `json:"cost,omitempty"`
		Amount  *Amount `json:"amount,omitempty"`
	}

	CustomAmount struct {
		Label  string  `json:"label,omitempty"`
		Amount *Amount `json:"amount,omitempty"`
	}

	InvoiceItem struct {
		Name      string `json:"name"`
		Quantity  int    `json:"quantity"`
		UnitPrice Amount `json:"unit_price"`
		// optional attributes
		Description string `json:"description,omitempty"`
		Tax         *Tax   `json:"tax,omitempty"`
		Date        string `json:"date,omitempty"`
		Discount    *Cost  `json:"cost,omitempty"`
	}

	MerchantInfo struct {
		Email          string   `json:"email"`
		FirstName      string   `json:"first_name,omitempty"`
		LastName       string   `json:"last_name,omitempty"`
		BusinessName   string   `json:"business_name,omitempty"`
		Address        *Address `json:"address,omitempty"`
		Phone          *Phone   `json:"phone,omitempty"`
		Fax            *Phone   `json:"fax,omitempty"`
		Website        string   `json:"website,omitempty"`
		TaxID          string   `json:"tax_id,omitempty"`
		AdditionalInfo string   `json:"additional_info,omitempty"`
	}

	PaymentTerm struct {
		TermType string `json:"term_type,omitempty"` // { DUE_ON_RECEIPT, NET_10, NET_15, NET_30, NET_45 }
		DueDate  string `json:"due_date,omitempty"`
	}

	ShippingCost struct {
		Amount *Amount `json:"amount,omitempty"`
		Tax    *Tax    `json:"tax,omitempty"`
	}

	ShippingInfo struct {
		FirstName    string `json:"first_name,omitempty"`
		LastName     string `json:"last_name,omitempty"`
		BusinessName string `json:"business_name,omitempty"`
	}

	Tax struct {
		ID      string  `json:"id,omitempty"`
		Name    string  `json:"name"`
		Percent float64 `json:"percent"` // 0.001 .. 99.999
		Amount  *Amount `json:"amount,omitempty"`
	}
)
