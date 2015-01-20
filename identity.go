// Package paypal defines types and operations used to access the Paypal API
//
// The following identity operations are defined
//
//      POST       /v1/identity/openidconnect/tokenservice
//      GET        /v1/identity/openidconnect/userinfo/?schema=<Schema>
package paypal

import (
	"fmt"
	"log"
	//	"strconv"
	//	"time"
)


// GetUserInfo retrieves user profile data using the AccessToken obtained using GrantToken()
func (c *Client) GetUserInfo(code string) (*UserInfo, error) {
        token, err := c.GrantToken(code)
        if err != nil {
                log.Fatalf("Error in GrantToken: %s\n", err)
        }
        log.Printf("Access Token: %v\n", token)

	req, err := NewRequest("GET", fmt.Sprintf("%s/identity/openidconnect/userinfo", c.APIBase), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.ClientID, c.Secret)

	// set scope of request to openid
	q := req.URL.Query()
	q.Set("schema", "openid")
	req.URL.RawQuery = q.Encode()

	v := &UserInfo{}

	// use the supplied AccessToken to get user info
	err = c.SendWithAuth(req, token.AccessToken, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
