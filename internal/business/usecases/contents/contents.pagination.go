package contents

const (
	defaultPage  = 1
	defaultLimit = 20
	maxLimit     = 100
)

func normalizePage(page int) int {
	if page <= 0 {
		return defaultPage
	}
	return page
}

func normalizeLimit(limit int) int {
	if limit <= 0 {
		return defaultLimit
	}
	if limit > maxLimit {
		return maxLimit
	}
	return limit
}
