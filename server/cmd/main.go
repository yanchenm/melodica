package main

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/joho/godotenv"

	server "github.com/yanchenm/musemoods"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load("../.env.yaml"); err != nil {
		log.Fatalf("couldn't read env: %v\n", err)
	}

	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/login", server.Login); err != nil {
		log.Fatalf("/login: %v\n", err)
	}
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/user", server.User); err != nil {
		log.Fatalf("/user: %v\n", err)
	}
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/getsongs", server.GetRecentlyPlayed); err != nil {
		log.Fatalf("/getsongs: %v\n", err)
	}

	port := "8080"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("start: %v\n", err)
	}
}