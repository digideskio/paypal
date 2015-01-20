package paypal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	// APIBaseSandBox points to the sandbox (for testing) version of the API
	APIBaseSandBox = "https://api.sandbox.paypal.com/v1"

	// APIBaseLive points to the live version of the API
	APIBaseLive = "https://api.paypal.com/v1"
)

type (

	// Client represents a Paypal REST API Client
	Client struct {
		client   *http.Client
		ClientID string
		Secret   string
		APIBase  string
		Token    *TokenResp
	}

	// ErrorResponse is used when a response contains errors
	// maps to error object
	ErrorResponse struct {
		// HTTP response that caused this error
		Response *http.Response `json:"-"`

		Name            string       `json:"name"`
		DebugID         string       `json:"debug_id"`
		Message         string       `json:"message"`
		InformationLink string       `json:"information_link"`
		Details         ErrorDetails `json:"details"`
	}

	// ErrorDetails map to error_details object
	ErrorDetails struct {
		Field string `json:"field"`
		Issue string `json:"issue"`
	}

	// TokenResp maps to the API response for the /oauth2/token endpoint
	TokenResp struct {
		Scope        string    `json:"scope"`        // "https://api.paypal.com/v1/payments/.* https://api.paypal.com/v1/vault/credit-card https://api.paypal.com/v1/vault/credit-card/.*"
		AccessToken  string    `json:"access_token"` // "EEwJ6tF9x5WCIZDYzyZGaz6Khbw7raYRIBV_WxVvgmsG"
		TokenType    string    `json:"token_type"`   // "Bearer"
		AppID        string    `json:"app_id"`       // "APP-6XR95014BA15863X"
		ExpiresIn    int       `json:"expires_in"`   // 28800
		ExpiresAt    time.Time `json:"expires_at"`
		RefreshToken string    `json:"refresh_token"`
	}

	Grant map[string]string
)

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v\nDetails: %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message, r.Details)
}

func (d *ErrorDetails) String() string {
	return fmt.Sprintf("Field: %s, Issue: %s", d.Field, d.Issue)
}

// NewClient returns a new Client struct
func NewClient(clientID, secret, APIBase string) *Client {
	return &Client{
		&http.Client{},
		clientID,
		secret,
		APIBase,
		nil,
	}
}

// NewRequest constructs a request. If payload is not empty, it will be
// marshalled into JSON
func NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		var b []byte
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequest(method, url, buf)
}

func grantFromAuthCode(code string) Grant {
	return map[string]string{
		"grant_type": "authorization_code",
		"code":       code,
	}
}

func grantFromRefreshToken(token string) Grant {
	return map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": token,
	}
}

// GetAccessToken requests a new access token from Paypal
func (c *Client) GetAccessToken() (*TokenResp, error) {
	buf := bytes.NewBuffer([]byte("grant_type=client_credentials"))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/oauth2/token"), buf)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.ClientID, c.Secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	t := TokenResp{}
	err = c.Send(req, &t)
	if err == nil {
		t.ExpiresAt = time.Now().Add(time.Duration(t.ExpiresIn/2) * time.Second)
	}

	return &t, err
}

// GrantToken grants a new access token from an authorization code
func (c *Client) GrantToken(code string) (*TokenInfo, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/identity/openidconnect/tokenservice", c.APIBase), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.ClientID, c.Secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	q := req.URL.Query()

	grant := grantFromAuthCode(code)
	for k, v := range grant {
		q.Set(k, v)
	}

	req.URL.RawQuery = q.Encode()

	v := &TokenInfo{}

	err = c.Send(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// RefreshToken grants a new access token from a refresh token
func (c *Client) RefreshToken(token string) (*TokenInfo, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/identity/openidconnect/tokenservice", c.APIBase), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.ClientID, c.Secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	q := req.URL.Query()

	grant := grantFromRefreshToken(token)
	for k, v := range grant {
		q.Set(k, v)
	}

	req.URL.RawQuery = q.Encode()

	v := &TokenInfo{}

	err = c.Send(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Send makes a request to the API, the response body will be
// unmarshaled into v, or if v is an io.Writer, the response will
// be written to it without decoding
func (c *Client) Send(req *http.Request, v interface{}) error {
	// Set default headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")

	// Default values for headers
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	log.Println(req.Method, ": ", req.URL)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	log.Printf("Send response: %v with status %d\n", resp.Body, resp.StatusCode)

	if httpStatus := resp.StatusCode; httpStatus < 200 || httpStatus > 299 {
		errResp := &ErrorResponse{Response: resp}
		data, err := ioutil.ReadAll(resp.Body)

		log.Printf("Error Response: %c\n", string(data))

		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}

		return errResp
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// SendAndAuth makes a request to the API and applies OAuth2 header automatically.
// If the access token expires soon, it will try to get a new one before
// making the main request
func (c *Client) SendAndAuth(req *http.Request, v interface{}) error {
	if (c.Token == nil) || (c.Token.ExpiresAt.Before(time.Now())) {
		resp, err := c.GetAccessToken()
		if err != nil {
			return err
		}

		c.Token = resp
	}
	req.Header.Set("Authorization", "Bearer "+c.Token.AccessToken)

	return c.Send(req, v)
}

// SendWithAuth makes a request to the API using the supplied Bearer accessToken
func (c *Client) SendWithAuth(req *http.Request, accessToken string, v interface{}) error {
	req.Header.Set("Authorization", "Bearer "+accessToken)

	log.Printf("Sending with Bearer token: %s\n", accessToken)

	return c.Send(req, v)
}
