package ghome

import (
	"encoding/json"
	"time"

	"google.golang.org/api/dialogflow/v2"
)

// Request represents the request.
type Request struct {
	parameters map[string]interface{}
	payload    *RequestPayload
	*dialogflow.GoogleCloudDialogflowV2WebhookRequest
}

// RequestPayloadInput represents the request payload input object.
type RequestPayloadInput struct {
	RawInputs []struct {
		Query     string `json:"query"`
		InputType string `json:"inputType"`
	} `json:"rawInputs"`
	Intent string `json:"intent"`
}

// RequestPayload represents the request payload object.
type RequestPayload struct {
	IsInSandbox bool `json:"isInSandbox"`
	Surface     struct {
		Capabilities []struct {
			Name string `json:"name"`
		} `json:"capabilities"`
	} `json:"surface"`
	Inputs []RequestPayloadInput `json:"inputs"`
	User   struct {
		LastSeen time.Time `json:"lastSeen"`
		Locale   string    `json:"locale"`
		UserID   string    `json:"userId"`
	} `json:"user"`
	Conversation struct {
		ConversationID string `json:"conversationId"`
		Type           string `json:"type"`
	} `json:"conversation"`
	AvailableSurfaces []struct {
		Capabilities []struct {
			Name string `json:"name"`
		} `json:"capabilities"`
	} `json:"availableSurfaces"`
}

// NewRequest creates a new request out of dialogflow webhook request.
func NewRequest(d *dialogflow.GoogleCloudDialogflowV2WebhookRequest) (*Request, error) {
	r := &Request{
		GoogleCloudDialogflowV2WebhookRequest: d,
	}

	if len(r.OriginalDetectIntentRequest.Payload) > 0 {
		if err := json.Unmarshal(r.OriginalDetectIntentRequest.Payload, &r.payload); err != nil {
			return nil, err
		}
	}

	if len(r.QueryResult.Parameters) > 0 {
		if err := json.Unmarshal(r.QueryResult.Parameters, &r.parameters); err != nil {
			return nil, err
		}
	}

	return r, nil
}

// Action returns the action value.
func (r *Request) Action() string {
	return r.QueryResult.Action
}

// IntentName returns the intent name.
func (r *Request) IntentName() string {
	return r.QueryResult.Intent.DisplayName
}

// Inputs returns all the inputs.
func (r *Request) Inputs() []RequestPayloadInput {
	return r.payload.Inputs
}

// Parameters returns collection of extracted parameters.
func (r *Request) Parameters() map[string]interface{} {
	return r.parameters
}

// SessionID returns the session id.
func (r *Request) SessionID() string {
	return r.Session
}

// UserID returns the user id.
func (r *Request) UserID() string {
	return r.payload.User.UserID
}
