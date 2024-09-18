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

// AssistantsV1AssistantWithToolsAndKnowledge struct for AssistantsV1AssistantWithToolsAndKnowledge
type AssistantsV1AssistantWithToolsAndKnowledge struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Assistant resource.
	AccountSid string `json:"account_sid"`
	// The Personalization and Perception Engine settings.
	CustomerAi map[string]interface{} `json:"customer_ai"`
	// The Assistant ID.
	Id string `json:"id"`
	// The default model used by the assistant.
	Model string `json:"model"`
	// The name of the assistant.
	Name string `json:"name"`
	// The owner/company of the assistant.
	Owner string `json:"owner"`
	// The personality prompt to be used for assistant.
	PersonalityPrompt string `json:"personality_prompt"`
	// The date and time in GMT when the Assistant was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated time.Time `json:"date_created"`
	// The date and time in GMT when the Assistant was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated time.Time `json:"date_updated"`
	// The list of knowledge sources associated with the assistant.
	Knowledge []AssistantsV1Knowledge `json:"knowledge"`
	// The list of tools associated with the assistant.
	Tools []AssistantsV1Tool `json:"tools"`
}
