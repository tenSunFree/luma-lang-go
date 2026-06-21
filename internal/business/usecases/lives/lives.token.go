package lives

import (
	"fmt"
	"math"
	"time"

	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
)

type agoraTokenService struct {
	appID          string
	appCertificate string
	ttlSeconds     int
}

func NewAgoraTokenService(appID, appCertificate string, ttlSeconds int) TokenService {
	return &agoraTokenService{
		appID:          appID,
		appCertificate: appCertificate,
		ttlSeconds:     ttlSeconds,
	}
}

// safeUint32 converts an int to uint32 with explicit bounds checking,
// satisfying gosec G115 and preventing silent overflow/underflow.
func safeUint32(name string, value int) (uint32, error) {
	if value < 0 {
		return 0, fmt.Errorf("%s must not be negative, got %d", name, value)
	}
	if int64(value) > int64(math.MaxUint32) {
		return 0, fmt.Errorf("%s exceeds uint32 max, got %d", name, value)
	}
	return uint32(value), nil
}

func (s *agoraTokenService) BuildRTCToken(channelName string, uid int, role string) (TokenResult, error) {
	ttl, err := safeUint32("ttlSeconds", s.ttlSeconds)
	if err != nil {
		return TokenResult{}, fmt.Errorf("build rtc token: %w", err)
	}

	agoraUID, err := safeUint32("uid", uid)
	if err != nil {
		return TokenResult{}, fmt.Errorf("build rtc token: %w", err)
	}

	expireAt := time.Now().Add(time.Duration(s.ttlSeconds) * time.Second)

	// Use the named constants of the package directly, without using int as an intermediate variable
	// rtctokenbuilder2.Role is uint16, so you can't assign int values to it.
	var agoraRole rtctokenbuilder.Role
	if role == RoleBroadcaster {
		agoraRole = rtctokenbuilder.RolePublisher
	} else {
		agoraRole = rtctokenbuilder.RoleSubscriber
	}

	token, err := rtctokenbuilder.BuildTokenWithUid(
		s.appID,
		s.appCertificate,
		channelName,
		agoraUID,
		agoraRole,
		ttl,
		ttl,
	)
	if err != nil {
		return TokenResult{}, err
	}

	return TokenResult{
		Token:    token,
		ExpireAt: expireAt,
	}, nil
}
