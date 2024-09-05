/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Marketplace
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateInstalledAddOn'
type CreateInstalledAddOnParams struct {
	// The SID of the AvaliableAddOn to install.
	AvailableAddOnSid *string `json:"AvailableAddOnSid,omitempty"`
	// Whether the Terms of Service were accepted.
	AcceptTermsOfService *bool `json:"AcceptTermsOfService,omitempty"`
	// The JSON object that represents the configuration of the new Add-on being installed.
	Configuration *interface{} `json:"Configuration,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be unique within the Account.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateInstalledAddOnParams) SetAvailableAddOnSid(AvailableAddOnSid string) *CreateInstalledAddOnParams {
	params.AvailableAddOnSid = &AvailableAddOnSid
	return params
}
func (params *CreateInstalledAddOnParams) SetAcceptTermsOfService(AcceptTermsOfService bool) *CreateInstalledAddOnParams {
	params.AcceptTermsOfService = &AcceptTermsOfService
	return params
}
func (params *CreateInstalledAddOnParams) SetConfiguration(Configuration interface{}) *CreateInstalledAddOnParams {
	params.Configuration = &Configuration
	return params
}
func (params *CreateInstalledAddOnParams) SetUniqueName(UniqueName string) *CreateInstalledAddOnParams {
	params.UniqueName = &UniqueName
	return params
}

// Install an Add-on for the Account specified.
func (c *ApiService) CreateInstalledAddOn(params *CreateInstalledAddOnParams) (*MarketplaceV1InstalledAddOn, error) {
	path := "/v1/InstalledAddOns"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.AvailableAddOnSid != nil {
		data.Set("AvailableAddOnSid", *params.AvailableAddOnSid)
	}
	if params != nil && params.AcceptTermsOfService != nil {
		data.Set("AcceptTermsOfService", fmt.Sprint(*params.AcceptTermsOfService))
	}
	if params != nil && params.Configuration != nil {
		v, err := json.Marshal(params.Configuration)

		if err != nil {
			return nil, err
		}

		data.Set("Configuration", string(v))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MarketplaceV1InstalledAddOn{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Add-on installation from your account
func (c *ApiService) DeleteInstalledAddOn(Sid string) error {
	path := "/v1/InstalledAddOns/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch an instance of an Add-on currently installed on this Account.
func (c *ApiService) FetchInstalledAddOn(Sid string) (*MarketplaceV1InstalledAddOn, error) {
	path := "/v1/InstalledAddOns/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MarketplaceV1InstalledAddOn{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListInstalledAddOn'
type ListInstalledAddOnParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListInstalledAddOnParams) SetPageSize(PageSize int) *ListInstalledAddOnParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListInstalledAddOnParams) SetLimit(Limit int) *ListInstalledAddOnParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of InstalledAddOn records from the API. Request is executed immediately.
func (c *ApiService) PageInstalledAddOn(params *ListInstalledAddOnParams, pageToken, pageNumber string) (*ListInstalledAddOnResponse, error) {
	path := "/v1/InstalledAddOns"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInstalledAddOnResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists InstalledAddOn records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInstalledAddOn(params *ListInstalledAddOnParams) ([]MarketplaceV1InstalledAddOn, error) {
	response, errors := c.StreamInstalledAddOn(params)

	records := make([]MarketplaceV1InstalledAddOn, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams InstalledAddOn records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInstalledAddOn(params *ListInstalledAddOnParams) (chan MarketplaceV1InstalledAddOn, chan error) {
	if params == nil {
		params = &ListInstalledAddOnParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MarketplaceV1InstalledAddOn, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageInstalledAddOn(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamInstalledAddOn(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamInstalledAddOn(response *ListInstalledAddOnResponse, params *ListInstalledAddOnParams, recordChannel chan MarketplaceV1InstalledAddOn, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.InstalledAddOns
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListInstalledAddOnResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListInstalledAddOnResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListInstalledAddOnResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInstalledAddOnResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateInstalledAddOn'
type UpdateInstalledAddOnParams struct {
	// Valid JSON object that conform to the configuration schema exposed by the associated AvailableAddOn resource. This is only required by Add-ons that need to be configured
	Configuration *interface{} `json:"Configuration,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be unique within the Account.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateInstalledAddOnParams) SetConfiguration(Configuration interface{}) *UpdateInstalledAddOnParams {
	params.Configuration = &Configuration
	return params
}
func (params *UpdateInstalledAddOnParams) SetUniqueName(UniqueName string) *UpdateInstalledAddOnParams {
	params.UniqueName = &UniqueName
	return params
}

// Update an Add-on installation for the Account specified.
func (c *ApiService) UpdateInstalledAddOn(Sid string, params *UpdateInstalledAddOnParams) (*MarketplaceV1InstalledAddOn, error) {
	path := "/v1/InstalledAddOns/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Configuration != nil {
		v, err := json.Marshal(params.Configuration)

		if err != nil {
			return nil, err
		}

		data.Set("Configuration", string(v))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MarketplaceV1InstalledAddOn{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}