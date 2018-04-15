# ghome [![Build Status](https://travis-ci.org/skillkit/go-ghome.svg?branch=master)](https://travis-ci.org/skillkit/go-ghome) [![GoDoc](https://godoc.org/github.com/skillkit/go-ghome?status.svg)](https://godoc.org/github.com/skillkit/go-ghome) [![Go Report Card](https://goreportcard.com/badge/github.com/skillkit/go-ghome)](https://goreportcard.com/report/github.com/skillkit/go-ghome)

> Work in progress

Go package for creating Google Home actions (fulfillment) via [Dialogflow](https://dialogflow.com). Not a hundred percent package for all Dialogflow features. The most basic features exists.

This package does only support V2 of Dialogflow API (not v1). More about how to use it [here](https://dialogflow.com/docs/reference/v2-agent-setup).

## Installation

```
go get -u github.com/skillkit/go-ghome
```

## Example with HTTP

```go
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
```

## License

MIT © [Fredrik Forsmo](https://github.com/frozzare)