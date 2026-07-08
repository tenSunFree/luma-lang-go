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

// StartTeacherLiveRequest is the body for POST /teacher/lives/start.
// title is required; other fields are optional.
type StartTeacherLiveRequest struct {
	Title        string  `json:"title"        binding:"required" example:"雅思閱讀 - 即時直播課"`
	Category     string  `json:"category"                        example:"語言學習"`
	Level        string  `json:"level"                           example:"A2"`
	CourseType   string  `json:"courseType"                      example:"live"`
	ThumbnailURL *string `json:"thumbnailUrl"                    example:"https://cdn.example.com/thumb.jpg"`
	TextbookURL  *string `json:"textbookUrl"                     example:"https://cdn.example.com/book.pdf"`
}
