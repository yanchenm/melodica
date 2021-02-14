package main

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"

	server "github.com/yanchenm/musemoods"
)

func main() {
	ctx := context.Background()

	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/login", server.LoginHTTP); err != nil {
		log.Fatalf("/login: %v\n", err)
	}
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/user", server.UserHTTP); err != nil {
		log.Fatalf("/user: %v\n", err)
	}

	port := "8080"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("start: %v\n", err)
	}
}
