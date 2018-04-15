package ghome

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"google.golang.org/api/dialogflow/v2"
)

const (
	requestJSON = `{"responseId":"CLZizlVv7g","queryResult":{"queryText":"GOOGLE_ASSISTANT_WELCOME","action":"location","parameters":{"hello":"world"},"allRequiredParamsPresent":true,"outputContexts":[{"name":"projects/fake-CLZizlVv7g/agent/sessions/CLZizlVv7g/contexts/google_assistant_welcome"},{"name":"projects/fake-CLZizlVv7g/agent/sessions/CLZizlVv7g/contexts/actions_capability_screen_output"},{"name":"projects/fake-CLZizlVv7g/agent/sessions/CLZizlVv7g/contexts/actions_capability_audio_output"},{"name":"projects/fake-CLZizlVv7g/agent/sessions/CLZizlVv7g/contexts/google_assistant_input_type_keyboard"},{"name":"projects/fake-CLZizlVv7g/agent/sessions/CLZizlVv7g/contexts/actions_capability_media_response_audio"},{"name":"projects/fake-CLZizlVv7g/agent/sessions/CLZizlVv7g/contexts/actions_capability_web_browser"}],"intent":{"name":"projects/fake-CLZizlVv7g/agent/intents/CLZizlVv7g","displayName":"Welcome"},"intentDetectionConfidence":1,"diagnosticInfo":{},"languageCode":"en-us"},"originalDetectIntentRequest":{"source":"google","version":"2","payload":{"isInSandbox":true,"surface":{"capabilities":[{"name":"actions.capability.MEDIA_RESPONSE_AUDIO"},{"name":"actions.capability.SCREEN_OUTPUT"},{"name":"actions.capability.AUDIO_OUTPUT"},{"name":"actions.capability.WEB_BROWSER"}]},"inputs":[{"rawInputs":[{"query":"Talk to my test app","inputType":"KEYBOARD"}],"intent":"actions.intent.MAIN"}],"user":{"lastSeen":"2018-04-14T19:23:26Z","locale":"en-US","userId":"fakeCLZizlVv7g--CLZizlVv7g_CLZizlVv7g"},"conversation":{"conversationId":"1111111111","type":"NEW"},"availableSurfaces":[{"capabilities":[{"name":"actions.capability.SCREEN_OUTPUT"},{"name":"actions.capability.AUDIO_OUTPUT"}]}]}},"session":"projects/fake-CLZizlVv7g/agent/sessions/CLZizlVv7g"}`
)

func createRequest(t *testing.T) *dialogflow.GoogleCloudDialogflowV2WebhookRequest {
	var req *dialogflow.GoogleCloudDialogflowV2WebhookRequest

	if err := json.Unmarshal([]byte(requestJSON), &req); err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	return req
}

func TestApp(t *testing.T) {
	app := NewApp()

	app.OnIntent(func(w ResponseWriter, r *Request) error {
		w.WriteSpeech("Hello, world!")
		return nil
	})

	req, err := NewRequest(createRequest(t))
	if err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	w, err := app.Process(req)
	if err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	if w.FulfillmentText != "Hello, world!" {
		t.Errorf("Expected 'Hello, world!' got %s", w.FulfillmentText)
	}
}

func TestAppVerifyRequest(t *testing.T) {
	app := NewApp()

	app.VerifyRequest(func(r *http.Request) error {
		return errors.New("Bad request")
	})

	if err := app.verifyRequest(nil); err == nil {
		t.Errorf("Expected error got: %v", err)
	}
}

func TestAppRequest(t *testing.T) {
	app := NewApp()

	app.OnIntent(func(w ResponseWriter, r *Request) error {
		if r.IntentName() != "Welcome" {
			t.Errorf("Expected 'Welcome' got %s", r.IntentName())
		}

		if r.Action() != "location" {
			t.Errorf("Expected 'location' got %s", r.Action())
		}

		if r.Parameters()["hello"] != "world" {
			t.Errorf("Expected 'world' got %s", r.Parameters()["hello"])
		}

		if r.Inputs()[0].RawInputs[0].Query != "Talk to my test app" {
			t.Errorf("Expected 'Talk to my test app' got %s", r.Inputs()[0].RawInputs[0].Query)
		}

		return nil
	})

	req, err := NewRequest(createRequest(t))
	if err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	if _, err := app.Process(req); err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}
}
