package ghome

import (
	"encoding/json"

	"google.golang.org/api/dialogflow/v2"
)

// ResponseWriter represents the interface that handle a response to Google Home.
type ResponseWriter interface {
	WriteCard(interface{})
	WriteFollowupEventInput(string, string, map[string]interface{}) error
	WriteOutputContext(string, int64, map[string]interface{}) error
	WritePayload(interface{}) error
	WriteSpeech(string)
	String() (string, error)
}

// Response represents the top level response object to Google Home.
type Response struct {
	*dialogflow.GoogleCloudDialogflowV2WebhookResponse
}

// NewResponse creates a new default response.
func NewResponse() *Response {
	return &Response{
		GoogleCloudDialogflowV2WebhookResponse: &dialogflow.GoogleCloudDialogflowV2WebhookResponse{
			FulfillmentMessages: []*dialogflow.GoogleCloudDialogflowV2IntentMessage{},
			OutputContexts:      []*dialogflow.GoogleCloudDialogflowV2Context{},
		},
	}
}

// WriteCard adds a new card to the JSON output.
func (r *Response) WriteCard(arg interface{}) {
	if basic, ok := arg.(*dialogflow.GoogleCloudDialogflowV2IntentMessageBasicCard); ok {
		r.FulfillmentMessages = append(r.FulfillmentMessages, &dialogflow.GoogleCloudDialogflowV2IntentMessage{
			BasicCard: basic,
		})
	}

	if card, ok := arg.(*dialogflow.GoogleCloudDialogflowV2IntentMessageCard); ok {
		r.FulfillmentMessages = append(r.FulfillmentMessages, &dialogflow.GoogleCloudDialogflowV2IntentMessage{
			Card: card,
		})
	}
}

// WriteFollowupEventInput adds a new followup event input to the JSON output.
func (r *Response) WriteFollowupEventInput(name, lang string, params map[string]interface{}) error {
	r.FollowupEventInput = &dialogflow.GoogleCloudDialogflowV2EventInput{
		Name:         name,
		LanguageCode: lang,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		return err
	}

	r.FollowupEventInput.Parameters = buf

	return nil
}

// WriteOutputContext adds a new output context to the JSON output.
func (r *Response) WriteOutputContext(name string, count int64, params map[string]interface{}) error {
	ctx := &dialogflow.GoogleCloudDialogflowV2Context{
		Name:          name,
		LifespanCount: count,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		return err
	}

	ctx.Parameters = buf
	r.OutputContexts = append(r.OutputContexts, ctx)

	return nil
}

// WritePayload adds the payload value to the JSON output.
func (r *Response) WritePayload(arg interface{}) error {
	switch v := arg.(type) {
	case string:
		r.Payload = []byte(v)
	case []byte:
		r.Payload = v
	default:
		buf, err := json.Marshal(arg)

		if err != nil {
			return err
		}

		r.Payload = buf
	}

	return nil
}

// WriteSpeech sets the JSON output speech value.
// You can also write SSML as a speech text and Dialogflow will handle it.
func (r *Response) WriteSpeech(content string) {
	r.FulfillmentText = content
}

// String will return the response as JSON.
func (r *Response) String() (string, error) {
	buf, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
