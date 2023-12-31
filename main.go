package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mattiasbonte/gobuzzer"
)

type Config struct {
	Telegram struct {
		AuthToken string `toml:"auth_token"`
		ChatID    string `toml:"chat_id"`
	} `toml:"telegram"`
}

func main() {
	// CONFIG
	config := loadConfigData()

	// FLAGS
	var (
		phone               bool
		system              bool
		notificationType    string
		notificationMessage string
	)
	flag.BoolVar(&phone, "p", false, "When passed sends phone notification")
	flag.BoolVar(&system, "s", false, "When passed sends system notification")
	flag.StringVar(&notificationType, "t", "notify", "System Notification type (beep, notify, alert)")
	flag.StringVar(&notificationMessage, "m", "Buzzzzr 🛸", "Notification message")
	flag.Parse()

	// SCRIPT
	if !phone && !system {
		log.Fatalf("Please specify at least one notification method, see --help")
	}

	if system {
		if err := gobuzzer.SystemBuzz(notificationMessage, notificationType); err != nil {
			log.Fatalf("error sending system notification: %v", err)
		}
		log.Println("🔊 System buzzed!")
	}

	if phone {
		if err := gobuzzer.TelegramNotification(config.Telegram.AuthToken, config.Telegram.ChatID, notificationMessage); err != nil {
			log.Fatalf("error sending telegram message: %v", err)
		}
		log.Println("📣 Phone buzzed!")
	}
}

func loadConfigData() Config {
	var userConfigPath = os.Getenv("HOME") + "/.config/notifyme/config.toml"
	if _, err := os.Stat(userConfigPath); os.IsNotExist(err) {
		log.Fatalf("missing user config file: %v", err)
	}

	var config = Config{}
	_, err := toml.DecodeFile(userConfigPath, &config)
	if err != nil {
		log.Fatalf("error decoding user config: %v", err)
	}

	return config
}
