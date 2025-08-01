package automation

import (
	"compare-it/config"
	"compare-it/internal/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
)

func blockNonEssentialRequests(_ context.Context) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		return network.SetBlockedURLs([]string{".ttf", ".css", ".png", ".jpg", ".gif"}).Do(ctx)
	})
}

func AutomateLoginPreOTP(ctx context.Context, provider, phone string) (*models.Session, error) {
	fmt.Println("Simulating successful response...")
	return &models.Session{
		ID:          uuid.NewString(),
		Provider:    provider,
		PhoneNumber: phone,
		Status:      "otp_sent",
	}, nil
}

func AutomateSubmitOTP(ctx context.Context, session *models.Session, otp string) (*models.Session, error) {
	sel, ok := config.Selectors[session.Provider]
	if !ok {
		return nil, errors.New("invalid provider")
	}

	// Restore cookies
	err := SetCookiesFromBytes(ctx, session.Cookies)
	if err != nil {
		return nil, err
	}

	err = chromedp.Run(ctx,
		chromedp.Navigate(session.CurrentURL),
		chromedp.WaitVisible(sel.OTPInput),
		chromedp.SendKeys(sel.OTPInput, otp),
		chromedp.Click(sel.SubmitOTP),
	)
	if err != nil {
		return nil, err
	}

	// Update session status
	session.Status = "logged_in"

	// Optional: Fetch new cookies and DOM snapshot if needed
	return session, nil
}

func getProviderURL(provider string) string {
	switch provider {
	case "blinkit":
		return "https://blinkit.com/"
	// Add more providers as needed
	default:
		return ""
	}
}

// SetCookiesFromBytes restores cookies from a byte array
func SetCookiesFromBytes(ctx context.Context, cookies []byte) error {
	var cookieParams []*network.CookieParam
	if err := json.Unmarshal(cookies, &cookieParams); err != nil {
		return err
	}

	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			for _, cookie := range cookieParams {
				err := network.SetCookie(cookie.Name, cookie.Value).
					WithDomain(cookie.Domain).
					WithPath(cookie.Path).
					WithHTTPOnly(cookie.HTTPOnly).
					WithSecure(cookie.Secure).
					WithSameSite(network.CookieSameSite(cookie.SameSite)).
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
	)
}

func NewContext(parentCtx context.Context) (context.Context, context.CancelFunc) {
	return chromedp.NewContext(parentCtx, chromedp.WithDebugf(log.Printf))
}

func AutomateLogin(ctx context.Context, provider, phone string) (*models.Session, error) {
	sel, ok := config.Selectors[provider]
	if !ok {
		return nil, errors.New("invalid provider")
	}

	// Set up a timeout for the entire login process
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second) // 1-minute timeout
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(getProviderURL(provider)),
		chromedp.Click(sel.LoginButton),
		chromedp.WaitVisible(sel.PhoneInput),
		chromedp.SendKeys(sel.PhoneInput, phone),
		chromedp.Click(sel.SubmitPhone),
		chromedp.WaitVisible(sel.OTPInput),
	)
	if err != nil {
		return nil, err
	}

	// Simulate sending OTP (in a real scenario, you would wait for the user to provide the OTP)
	session, err := AutomateLoginPreOTP(ctx, provider, phone)
	if err != nil {
		return nil, err
	}

	return session, nil
}
