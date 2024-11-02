package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(a *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := a.Listen(getPort()); err != nil {
		log.Printf("Oops... Server is not running! Reason port: %v", err)
	}

	<-idleConnsClosed
}

func StartServer(a *fiber.App) {
	if err := a.Listen(getPort()); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

func getPort() string {
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))

	return fmt.Sprintf("%v", port)
}
