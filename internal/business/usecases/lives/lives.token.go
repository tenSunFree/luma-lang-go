package lives

import (
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

func (s *agoraTokenService) BuildRTCToken(channelName string, uid int, role string) (TokenResult, error) {
	expireAt := time.Now().Add(time.Duration(s.ttlSeconds) * time.Second)
	ttl := uint32(s.ttlSeconds)

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
		uint32(uid),
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
