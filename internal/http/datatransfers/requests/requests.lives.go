package requests

// JoinLiveRequest is the body for POST /lives/:liveId/join.
type JoinLiveRequest struct {
	ClientType string `json:"clientType" example:"android"`
}

// LeaveLiveRequest is the body for POST /lives/:liveId/leave.
type LeaveLiveRequest struct {
	UID int `json:"uid" binding:"required" example:"10042"`
}

// RenewTokenRequest is the body for POST /lives/:liveId/renew-token.
type RenewTokenRequest struct {
	UID int `json:"uid" binding:"required" example:"10042"`
}

// TeacherRenewTokenRequest is the body for POST /teacher/lives/:liveId/renew-token.
type TeacherRenewTokenRequest struct {
	UID        int    `json:"uid"        example:"1000"`
	StreamType string `json:"streamType" example:"camera"`
}
