package api

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	TwilioAccountSID  string
	TwilioAuthToken    string
	TwilioServicesID   string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TwilioAccountSID = os.Getenv("TWILIO_ACCOUNT_SID")
	TwilioAuthToken = os.Getenv("TWILIO_AUTH_TOKEN")
	TwilioServicesID = os.Getenv("TWILIO_SERVICES_ID")
}
