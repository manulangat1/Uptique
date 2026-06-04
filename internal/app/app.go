package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Run() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	log.Println("Starting Uptique...")

	ticker := time.NewTicker(5 * time.Second)

	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down...")
			return
		case <-ticker.C:
			log.Println("Heartbeat")
		}
	}
}
