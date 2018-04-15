package main

import (
	"net/http"
	
	"github.com/skillkit/go-ghome"
)

func main() {
	app := ghome.NewApp(nil)

	app.OnIntent(func(w ghome.ResponseWriter, r *ghome.Request) error {
		w.WriteSpeech("Hello, world!")
		return nil
	})

	http.Handle("/", app.Handler())
	http.ListenAndServe(":3000", nil)
}
