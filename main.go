package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mattiasbonte/gobuzzer"
)

func main() {
	// Flags
	var (
		notificationMessage string
		notificationType    string
	)
	flag.StringVar(&notificationMessage, "m", "", "Notification message")
	flag.StringVar(&notificationType, "t", "notify", "Notification type (beep, notify, alert)")
	flag.Parse()

	// Env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		telegramAuthToken = os.Getenv("TELEGRAM_AUTH_TOKEN")
		telegramChatID    = os.Getenv("TELEGRAM_CHAT_ID")
	)

	// Script
	if notificationMessage == "" {
		log.Fatal("error: pass a message with -m flag")
	}

	if err := gobuzzer.TelegramNotification(telegramAuthToken, telegramChatID, notificationMessage); err != nil {
		log.Fatalf("error sending telegram message: %v", err)
	}

	if err := gobuzzer.SystemBuzz(notificationMessage, notificationType); err != nil {
		log.Fatalf("error sending system notification: %v", err)
	}

	log.Println("🛸 notification sent")
}
