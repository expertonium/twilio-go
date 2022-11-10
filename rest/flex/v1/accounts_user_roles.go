/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
)

// Optional parameters for the method 'FetchUserRoles'
type FetchUserRolesParams struct {
	// The Token HTTP request header
	Token *string `json:"Token,omitempty"`
}

func (params *FetchUserRolesParams) SetToken(Token string) *FetchUserRolesParams {
	params.Token = &Token
	return params
}

// This is used by Flex UI and Quality Management to fetch the Flex Insights roles for the user
func (c *ApiService) FetchUserRoles(params *FetchUserRolesParams) (*FlexV1UserRoles, error) {
	path := "/v1/Accounts/UserRoles"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Token != nil {
		headers["Token"] = *params.Token
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1UserRoles{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
