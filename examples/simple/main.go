package main

import (
	"net/http"
)

func main() {
	app := ghome.NewApp(&ghome.Options{
		Source: "my-app",
	})

	app.OnIntent(func(w ghome.ResponseWriter, r *ghome.Request) error {
		w.WriteSpeech("Hello, world!")
		return nil
	})

	http.Handle("/", app.Handler())
	http.ListenAndServe(":3000", nil)
}
