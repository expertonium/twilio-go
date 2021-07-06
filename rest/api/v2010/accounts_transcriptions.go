/*
 * Twilio - Api
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

// Optional parameters for the method 'DeleteTranscription'
type DeleteTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteTranscriptionParams) SetPathAccountSid(PathAccountSid string) *DeleteTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a transcription from the account used to make the request
func (c *ApiService) DeleteTranscription(Sid string, params *DeleteTranscriptionParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Transcriptions/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

// Optional parameters for the method 'FetchTranscription'
type FetchTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchTranscriptionParams) SetPathAccountSid(PathAccountSid string) *FetchTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a Transcription
func (c *ApiService) FetchTranscription(Sid string, params *FetchTranscriptionParams) (*ApiV2010AccountTranscription, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Transcriptions/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountTranscription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTranscription'
type ListTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListTranscriptionParams) SetPathAccountSid(PathAccountSid string) *ListTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListTranscriptionParams) SetPageSize(PageSize int) *ListTranscriptionParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of transcriptions belonging to the account used to make the request
func (c *ApiService) ListTranscription(params *ListTranscriptionParams) (*ListTranscriptionResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Transcriptions.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTranscriptionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}