package auth

import "time"

// Config is the slice of configuration the auth use case needs.
type Config struct {
	OTPMaxAttempts int
	OTPTTL         time.Duration

	// PwdResetCodeTTL is how long a password-reset 6-digit code
	// (and its attempt counter) stay live in Redis.
	PwdResetCodeTTL time.Duration
	// PwdResetMaxAttempts is the lockout threshold for ResetPassword.
	PwdResetMaxAttempts int

	BcryptCost       int
	LoginMaxAttempts int
	LoginLockoutTTL  time.Duration

	ForgotMaxAttempts int
	ForgotLockoutTTL  time.Duration
}
