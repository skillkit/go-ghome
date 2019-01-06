package ghome

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	dialogflow "google.golang.org/api/dialogflow/v2"
)

// HandleFunc represents the handler function.
type HandleFunc func(ResponseWriter, *Request) error

// VerifyRequestFunc represents the verify request function.
type VerifyRequestFunc func(*http.Request) error

// App represents the Google Home app.
type App struct {
	verifyRequest VerifyRequestFunc
	onIntent      HandleFunc
	source        string
}

// NewApp creates a new Home app.
func NewApp() *App {
	return &App{}
}

// OnIntent sets the intent handler.
func (a *App) OnIntent(h HandleFunc) {
	a.onIntent = h
}

// VerifyRequest sets the verify request function.
// This should be used to set a custom function to verify the request from Dialogflow.
func (a *App) VerifyRequest(h VerifyRequestFunc) {
	a.verifyRequest = h
}

// Process handles a request passed from Google Home.
func (a *App) Process(r *Request) (*Response, error) {
	w := NewResponse()

	// Run intent function if it exists.
	if a.onIntent != nil {
		if err := a.onIntent(w, r); err != nil {
			return nil, err
		}
	}

	return w, nil
}

// Handler returns a http handler to hook Google Home app into a http server.
func (a *App) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req *dialogflow.GoogleCloudDialogflowV2WebhookRequest

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Bail if POST method.
		if strings.ToUpper(r.Method) != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Bail if JSON decode failes.
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Verify dialogflow request if verify request function is set.
		if a.verifyRequest != nil {
			if err := a.verifyRequest(r); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		// Create our own request based on dialogflow request.
		req2, err := NewRequest(req)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := a.Process(req2)

		// Bail if process request failes.
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(resp)
	})
}
