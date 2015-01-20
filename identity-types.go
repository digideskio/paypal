package paypal

import "fmt"

type (
	IdentityAddress struct {
		StreetAddress string `json:"street_address"`
		Locality      string `json:"locality"`
		Region        string `json:"region"`
		PostalCode    string `json:"postal_code"`
		Country       string `json:"country"`
	}

	IdentityError struct {
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
		ErrorURI         string `json:"error_uri"`
	}

	TokenInfo struct {
		Scope        string `json:"scope,omitempty"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    string `json:"expires_in"`
	}

	UserInfo struct {
		AccountType     string          `json:"account_type"` // either personal or business
		UserID          string          `json:"user_id"`
		Email           string          `json:"email"`
		VerifiedEmail   bool            `json:"verified"`
		Name            string          `json:"name"`
		FamilyName      string          `json:"family_name"`
		GivenName       string          `json:"given_name"`
		Gender          string          `json:"gender"`
		BirthDate       string          `json:"birth_date"` // YYYY-MM-DD or YYYY - 0000 ommited
		AgeRange        string          `json:"age_range"`
		Address         IdentityAddress `json:"address"`
		PhoneNumber     string          `json:"phone_number"`
		VerifiedAccount bool            `json:"verified_account"`
		ZoneInfo        string          `json:"zoneinfo"`
		Locale          string          `json:"locale"`
		CreateDate      string          `json:"account_creation_date"`
	}

)

func (t *TokenInfo) String() string {
	return fmt.Sprintf("Type: %s, Scope: %s, AccessToken: %s, RefreshToken: %s, ExpiresIn: %s", t.TokenType, t.Scope, t.AccessToken, t.RefreshToken, t.ExpiresIn)
}

func (u *UserInfo) String() string {
	return fmt.Sprintf("Email: %s verified: %t", u.Email, u.VerifiedEmail)
}

