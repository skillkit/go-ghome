package ghome

import "testing"

func TestResponseCard(t *testing.T) {
	resp := NewResponse()

	resp.WriteCard(&Card{
		Title:    "Hello Gopher",
		Subtitle: "",
		ImageUri: "https://golang.org/doc/gopher/frontpage.png",
		Buttons: CardButtons(&CardButton{
			Text:     "Learn more about Go",
			Postback: "https://golang.org",
		}),
	})

	if resp.FulfillmentMessages[0].Card.Title != "Hello Gopher" {
		t.Errorf("Expected 'Hello Gopher' got %s", resp.FulfillmentMessages[0].Card.Title)
	}

	if resp.FulfillmentMessages[0].Card.ImageUri != "https://golang.org/doc/gopher/frontpage.png" {
		t.Errorf("Expected 'https://golang.org/doc/gopher/frontpage.png' got %s", resp.FulfillmentMessages[0].Card.ImageUri)
	}
}

func TestResponseWriteFollowupEventInput(t *testing.T) {
	resp := NewResponse()

	err := resp.WriteFollowupEventInput("Test", "en", map[string]interface{}{
		"hello": "world",
	})

	if err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	if resp.FollowupEventInput.Name != "Test" {
		t.Errorf("Expected: 'Test' got %s", resp.FollowupEventInput.Name)
	}

	if string(resp.FollowupEventInput.Parameters) != `{"hello":"world"}` {
		t.Errorf("Expected: '{\"hello\":\"world\"}' got %s", resp.FollowupEventInput.Parameters)
	}
}

func TestResponseWriteOutputContext(t *testing.T) {
	resp := NewResponse()

	err := resp.WriteOutputContext("Test", 0, map[string]interface{}{
		"hello": "world",
	})

	if err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	if resp.OutputContexts[0].Name != "Test" {
		t.Errorf("Expected: 'Test' got %s", resp.OutputContexts[0].Name)
	}

	if string(resp.OutputContexts[0].Parameters) != `{"hello":"world"}` {
		t.Errorf("Expected: '{\"hello\":\"world\"}' got %s", resp.OutputContexts[0].Parameters)
	}
}
