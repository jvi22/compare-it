package api

import (
	"errors"
	"os"
	"sync"

	twilio "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

// ---- Helper functions to safely fetch env vars ----

func envACCOUNTSID() string {
	val := os.Getenv("TWILIO_ACCOUNT_SID")
	if val == "" {
		panic("TWILIO_ACCOUNT_SID not set in environment")
	}
	return val
}

func envAUTHTOKEN() string {
	val := os.Getenv("TWILIO_AUTH_TOKEN")
	if val == "" {
		panic("TWILIO_AUTH_TOKEN not set in environment")
	}
	return val
}

func envSERVICESID() string {
	val := os.Getenv("TWILIO_SERVICES_ID")
	if val == "" {
		panic("TWILIO_SERVICES_ID not set in environment")
	}
	return val
}

// ---- Lazy-initialized Twilio client ----

var (
	client     *twilio.RestClient
	clientOnce sync.Once
)

func getTwilioClient() *twilio.RestClient {
	clientOnce.Do(func() {
		client = twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: envACCOUNTSID(),
			Password: envAUTHTOKEN(),
		})
	})
	return client
}

// ---- OTP functions ----

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := getTwilioClient().VerifyV2.CreateVerification(envSERVICESID(), params)
	if err != nil {
		return "", err
	}
	if resp.Sid == nil {
		return "", errors.New("failed to get verification SID from Twilio")
	}
	return *resp.Sid, nil
}

func (app *Config) twilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := getTwilioClient().VerifyV2.CreateVerificationCheck(envSERVICESID(), params)
	if err != nil {
		return err
	}
	if resp.Status == nil || *resp.Status != "approved" {
		return errors.New("not a valid code")
	}
	return nil
}
