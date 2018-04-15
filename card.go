package ghome

import "google.golang.org/api/dialogflow/v2"

// BasicCard represents a basic card.
// Just a alias for dialogflow.GoogleCloudDialogflowV2IntentMessageBasicCard.
type BasicCard = dialogflow.GoogleCloudDialogflowV2IntentMessageBasicCard

// CardButton represents a card button.
// Just a alias for dialogflow.GoogleCloudDialogflowV2IntentMessageCardButton.
type CardButton = dialogflow.GoogleCloudDialogflowV2IntentMessageCardButton

// Card represents a card.
// Just a alias for dialogflow.GoogleCloudDialogflowV2IntentMessageCard.
type Card = dialogflow.GoogleCloudDialogflowV2IntentMessageCard

// NewCard creates a new card.
func NewCard() *Card {
	return &Card{}
}

// NewBasicCard creates a new basic card.
func NewBasicCard() *BasicCard {
	return &BasicCard{}
}

// CardButtons returns buttons.
func CardButtons(args ...*CardButton) []*dialogflow.GoogleCloudDialogflowV2IntentMessageCardButton {
	return args
}
