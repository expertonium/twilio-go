/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Assistants
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// AssistantsV1Session struct for AssistantsV1Session
type AssistantsV1Session struct {
	// The Session ID.
	Id string `json:"id,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Session resource.
	AccountSid string `json:"account_sid,omitempty"`
	// The Assistant ID.
	AssistantId string `json:"assistant_id,omitempty"`
	// True if the session is verified.
	Verified bool `json:"verified,omitempty"`
	// The unique identity of user for the session.
	Identity string `json:"identity,omitempty"`
	// The date and time in GMT when the Session was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the Session was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated time.Time `json:"date_updated,omitempty"`
}
