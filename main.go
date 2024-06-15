package main

import (
	"context"

	"entrlcom.dev/mongox/docker"
)

func main() {
	ctx := context.Background()

	mongo, err := docker.NewMongo(ctx)
	if err != nil {
		// TODO: Handle error.
		return
	}

	defer func() {
		// TODO: Handle error.
		_ = mongo.Shutdown(ctx) //nolint:errcheck // OK.
	}()

	client := mongo.GetClient()

	// ...
}
