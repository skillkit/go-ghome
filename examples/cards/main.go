package main

import (
	"net/http"

	"github.com/skillkit/go-ghome"
	"google.golang.org/api/dialogflow/v2"
)

func main() {
	app := ghome.NewApp(nil)

	app.OnIntent(func(w ghome.ResponseWriter, r *ghome.Request) error {
		w.WriteCard(&dialogflow.GoogleCloudDialogflowV2IntentMessageCard{
			Title:    "Hello Gopher",
			Subtitle: "Nice day",
			ImageUri: "https://golang.org/doc/gopher/frontpage.png",
			Buttons: []*dialogflow.GoogleCloudDialogflowV2IntentMessageCardButton{
				{
					Text:     "Learn more about Go",
					Postback: "https://golang.org",
				},
			},
		})

		w.WriteCard(&dialogflow.GoogleCloudDialogflowV2IntentMessageBasicCard{
			Title:    "Hello Gopher",
			Subtitle: "Nice day",
			Image: &dialogflow.GoogleCloudDialogflowV2IntentMessageImage{
				AccessibilityText: "Gopher",
				ImageUri:          "https://golang.org/doc/gopher/frontpage.png",
			},
			Buttons: []*dialogflow.GoogleCloudDialogflowV2IntentMessageBasicCardButton{
				{
					Title: "Learn more about Go",
					OpenUriAction: &dialogflow.GoogleCloudDialogflowV2IntentMessageBasicCardButtonOpenUriAction{
						Uri: "https://golang.org",
					},
				},
			},
		})

		return nil
	})

	http.Handle("/", app.Handler())
	http.ListenAndServe(":3000", nil)
}
