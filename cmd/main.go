package main

import (
	"API-email-confirmation/configs"
	"API-email-confirmation/internal/verify"
	"fmt"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		Config: config,
	})

	server := http.Server{
		Addr:    ":" + config.Project.PORT,
		Handler: router,
	}

	fmt.Println("Server is running on port", config.Project.PORT)
	server.ListenAndServe()
}
