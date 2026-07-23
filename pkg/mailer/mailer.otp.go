package mailer

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"time"

	gomail "gopkg.in/mail.v2"
)

// Embedded so the binary is self-contained — distroless runtime has
// no shell or filesystem to load templates from at deploy time.
//
//go:embed templates/*.html
var templatesFS embed.FS

// otpTpl is parsed once at package init. html/template (not text)
// auto-escapes the OTP code, defending against an attacker
// somehow injecting markup into the OTP path.
var otpTpl = template.Must(template.ParseFS(templatesFS, "templates/otp.html"))
var passwordResetTpl = template.Must(template.ParseFS(templatesFS, "templates/password_reset.html"))

// otpTemplateData feeds the template. AppName / Region are constants
// today; pulling them out makes white-labeling and i18n a config
// change rather than a code edit.
type otpTemplateData struct {
	AppName      string
	Region       string
	Code         string
	Year         int
	ValidMinutes int
}

const (
	defaultAppName      = "Go Rest boilerplate"
	defaultRegion       = "East Java, Indonesia"
	defaultValidMinutes = 5
)

type OTPMailer interface {
	SendOTP(otpCode string, receiver string) (err error)
	// SendPasswordReset delivers the six-digit password-reset code
	// to the receiver's inbox, using a dedicated template so the
	// user can tell it apart from an account-activation OTP.
	SendPasswordReset(code string, receiver string) error
}

type otpMailer struct {
	email    string
	password string
}

func NewOTPMailer(email, password string) OTPMailer {
	return &otpMailer{
		email:    email,
		password: password,
	}
}

func (mailer *otpMailer) SendOTP(otpCode, receiver string) (err error) {
	body, err := renderOTPBody(otpCode)
	if err != nil {
		return fmt.Errorf("render otp template: %w", err)
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", mailer.email)
	msg.SetHeader("To", receiver)
	msg.SetHeader("Subject", "Verification Email")
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, mailer.email, mailer.password)
	dialer.Timeout = 10 * time.Second

	return dialer.DialAndSend(msg)
}

func (mailer *otpMailer) SendPasswordReset(code, receiver string) error {
	body, err := renderPasswordResetBody(code)
	if err != nil {
		return fmt.Errorf("render password reset template: %w", err)
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", mailer.email)
	msg.SetHeader("To", receiver)
	msg.SetHeader("Subject", "Password Reset Code")
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, mailer.email, mailer.password)
	dialer.Timeout = 10 * time.Second
	return dialer.DialAndSend(msg)
}

func renderPasswordResetBody(code string) (string, error) {
	var buf bytes.Buffer
	data := otpTemplateData{
		AppName: defaultAppName, Region: defaultRegion,
		Code: code, Year: time.Now().Year(), ValidMinutes: defaultValidMinutes,
	}
	if err := passwordResetTpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// renderOTPBody is exported as a helper for tests so they can assert
// on the rendered HTML without spinning up an SMTP dialer.
func renderOTPBody(code string) (string, error) {
	var buf bytes.Buffer
	data := otpTemplateData{
		AppName:      defaultAppName,
		Region:       defaultRegion,
		Code:         code,
		Year:         time.Now().Year(),
		ValidMinutes: defaultValidMinutes,
	}
	if err := otpTpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
