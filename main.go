package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/rombintu/bbgpt3/api"
)

func main() {
	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exitCh
		log.Println("Exit with 0")
		os.Exit(0)
	}()

	godotenv.Load()
	tokenApi := os.Getenv("GPT_TOKEN")
	tokenBot := os.Getenv("BOT_TOKEN")
	secret := os.Getenv("SECRET")
	if tokenApi == "" || secret == "" || tokenBot == "" {
		log.Fatalln("Missing env: GPT_TOKEN/BOT_TOKEN or SECRET")
	}

	log.Println("Service is configured successfully")
	log.Println("Go to http://localhost:9001")
	api.RunApi("9001", tokenApi, tokenBot, secret)
	// if err := api.ListenAndServe(); err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }
}
