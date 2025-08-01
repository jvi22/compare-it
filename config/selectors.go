package config

type ProviderSelector struct {
    LoginButton      string
    PhoneInput       string
    SubmitPhone      string
    OTPInput         string
    SubmitOTP        string
}

var Selectors = map[string]ProviderSelector{
    "blinkit": {
        LoginButton:   "#loginBtn",
        PhoneInput:    "#phoneInput",
        SubmitPhone:   "#submitPhone",
        OTPInput:      "#otpInput",
        SubmitOTP:     "#submitOtpBtn",
    },
    // populate for zepto, instamart similarly
}
