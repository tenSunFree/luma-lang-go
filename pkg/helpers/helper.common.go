package helpers

import "strings"

func IsArrayContains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

// MaskEmail redacts the local part of an email address for logging,
// e.g. "patrick@example.com" -> "p***@example.com". Keeps the first
// character and the domain intact so logs stay useful for debugging
// without exposing the full address. Malformed input (no "@") is
// masked entirely to avoid leaking anything unexpected.
func MaskEmail(email string) string {
	at := strings.LastIndex(email, "@")
	if at <= 0 {
		return "***"
	}
	local := email[:at]
	domain := email[at:]
	if len(local) <= 1 {
		return "*" + domain
	}
	return local[:1] + strings.Repeat("*", len(local)-1) + domain
}
