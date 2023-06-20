package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mattiasbonte/gobuzzer"
)

func main() {
	// Env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		telegramAuthToken = os.Getenv("TELEGRAM_AUTH_TOKEN")
		telegramChatID    = os.Getenv("TELEGRAM_CHAT_ID")
	)

	// Flags
	var (
		noPhone             bool
		notificationType    string
		notificationMessage string
	)
	flag.BoolVar(&noPhone, "n", false, "When passed skips sending phone notification")
	flag.StringVar(&notificationType, "t", "notify", "System Notification type (beep, notify, alert)")
	flag.StringVar(&notificationMessage, "m", "Buzzzzr 🛸", "Notification message")
	flag.Parse()

	// Script
	if err := gobuzzer.SystemBuzz(notificationMessage, notificationType); err != nil {
		log.Fatalf("error sending system notification: %v", err)
	}
	log.Println("🔊 System buzzed!")

	if !noPhone {
		if err := gobuzzer.TelegramNotification(telegramAuthToken, telegramChatID, notificationMessage); err != nil {
			log.Fatalf("error sending telegram message: %v", err)
		}
		log.Println("📣 Phone buzzed!")
	}

}
