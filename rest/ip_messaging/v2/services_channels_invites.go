/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Optional parameters for the method 'CreateInvite'
type CreateInviteParams struct {
	//
	Identity *string `json:"Identity,omitempty"`
	//
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *CreateInviteParams) SetIdentity(Identity string) *CreateInviteParams {
	params.Identity = &Identity
	return params
}
func (params *CreateInviteParams) SetRoleSid(RoleSid string) *CreateInviteParams {
	params.RoleSid = &RoleSid
	return params
}

func (c *ApiService) CreateInvite(ServiceSid string, ChannelSid string, params *CreateInviteParams) (*IpMessagingV2ServiceChannelInvite, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2ServiceChannelInvite{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteInvite(ServiceSid string, ChannelSid string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchInvite(ServiceSid string, ChannelSid string, Sid string) (*IpMessagingV2ServiceChannelInvite, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2ServiceChannelInvite{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListInvite'
type ListInviteParams struct {
	//
	Identity *[]string `json:"Identity,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListInviteParams) SetIdentity(Identity []string) *ListInviteParams {
	params.Identity = &Identity
	return params
}
func (params *ListInviteParams) SetPageSize(PageSize int) *ListInviteParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListInvite(ServiceSid string, ChannelSid string, params *ListInviteParams) (*ListInviteResponse, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Identity != nil {
		for _, item := range *params.Identity {
			data.Add("Identity", item)
		}
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInviteResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}