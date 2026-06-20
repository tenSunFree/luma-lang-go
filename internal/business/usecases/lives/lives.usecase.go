package lives

import (
	"context"
	"time"
)

// UID Planning
const (
	StatusScheduled = "scheduled"
	StatusLive      = "live"
	StatusEnded     = "ended"
	StatusCancelled = "cancelled"

	RoleAudience    = "audience"
	RoleBroadcaster = "broadcaster"
	RoleTeacher     = "teacher"

	TeacherCameraUID = 1000
	TeacherScreenUID = 2000
	StudentUIDBase   = 10000
)

// Usecase is the business layer interface of lives, and the handler layer only depends on this.
type Usecase interface {
	ListLiveCourses(ctx context.Context, req ListLiveCoursesRequest) (ListLiveCoursesResponse, error)
	JoinLive(ctx context.Context, req JoinLiveRequest) (JoinLiveResponse, error)
	LeaveLive(ctx context.Context, req LeaveLiveRequest) error
	RenewToken(ctx context.Context, req RenewTokenRequest) (RenewTokenResponse, error)
	GetParticipants(ctx context.Context, req GetParticipantsRequest) (GetParticipantsResponse, error)
	StartTeacherLive(ctx context.Context, req StartTeacherLiveRequest) (StartTeacherLiveResponse, error)
	EndTeacherLive(ctx context.Context, req EndTeacherLiveRequest) error
	SetReminder(ctx context.Context, req SetReminderRequest) (ReminderResponse, error)
	DeleteReminder(ctx context.Context, req DeleteReminderRequest) (ReminderResponse, error)
}

// Repository is a storage abstraction of lives, implemented by postgres.
// It's placed in the usecase layer (not the interface layer) because this interface is business-specific.
// This aligns with existing contents usecase implementations without requiring additional interface packages.
type Repository interface {
	ListLiveCourses(ctx context.Context, statuses []string, userID string) ([]LiveCourse, error)
	GetLiveByID(ctx context.Context, liveID string) (LiveCourse, error)
	GetNextStudentUID(ctx context.Context, liveID string) (int, error)
	UpsertParticipant(ctx context.Context, in UpsertParticipantInput) error
	LeaveParticipant(ctx context.Context, liveID, userID string, uid int) error
	ListParticipants(ctx context.Context, liveID string) ([]Participant, error)
	MarkLiveStarted(ctx context.Context, liveID string) error
	MarkLiveEnded(ctx context.Context, liveID string) error
	SetReminder(ctx context.Context, liveID, userID string, remindAt time.Time) error
	DeleteReminder(ctx context.Context, liveID, userID string) error
}

// TokenService abstracts the generation of Agora tokens, making it convenient to mock during testing.
type TokenService interface {
	BuildRTCToken(channelName string, uid int, role string) (TokenResult, error)
}

// TokenResult contains the generated token and its expiration time.
type TokenResult struct {
	Token    string
	ExpireAt time.Time
}

// --- domain structs ---
type LiveCourse struct {
	LiveID            string
	CourseID          string
	Title             string
	Category          string
	Level             string
	CourseType        string
	Status            string
	ScheduledStartAt  time.Time
	StartedAt         *time.Time
	EndedAt           *time.Time
	TeacherID         string
	TeacherName       string
	TeacherAvatarURL  *string
	ThumbnailURL      *string
	TextbookURL       *string
	AgoraChannelName  string
	TeacherCameraUID  int
	TeacherScreenUID  int
	ViewerCount       int
	IsReminderEnabled bool
}

type Participant struct {
	UserID      string
	DisplayName string
	AvatarURL   *string
	Role        string
	AgoraUID    int
	IsMuted     bool
	IsCameraOn  bool
}

type UpsertParticipantInput struct {
	LiveID      string
	UserID      string
	DisplayName string
	AvatarURL   *string
	AgoraUID    int
	Role        string
}

// --- request / response types ---
type ListLiveCoursesRequest struct {
	UserID   string
	Statuses []string
}

type ListLiveCoursesResponse struct {
	Items []LiveCourse
}

type JoinLiveRequest struct {
	LiveID      string
	UserID      string
	DisplayName string
	AvatarURL   *string
	ClientType  string
}

type JoinLiveResponse struct {
	LiveID       string
	CourseID     string
	Agora        AgoraJoinConfig
	Teacher      TeacherConfig
	StreamLayout StreamLayout
	Chat         ChatConfig
	Features     LiveFeatures
}

type AgoraJoinConfig struct {
	AppID         string
	ChannelName   string
	UID           int
	RTCToken      string
	Role          string
	TokenExpireAt time.Time
}

type TeacherConfig struct {
	TeacherID string
	Name      string
	AvatarURL *string
	CameraUID int
	ScreenUID int
}

type StreamLayout struct {
	MainUID  int
	MainType string
	PipUID   int
	PipType  string
}

type ChatConfig struct {
	Enabled  bool
	Provider string
	RoomID   string
	WSUrl    string
}

type LiveFeatures struct {
	CanSendMessage  bool
	CanRaiseHand    bool
	CanSendReaction bool
	CanPublishAudio bool
	CanPublishVideo bool
}

type LeaveLiveRequest struct {
	LiveID string
	UserID string
	UID    int
}

type RenewTokenRequest struct {
	LiveID string
	UID    int
	Role   string
}

type RenewTokenResponse struct {
	UID           int
	RTCToken      string
	TokenExpireAt time.Time
}

type GetParticipantsRequest struct {
	LiveID string
}

type GetParticipantsResponse struct {
	TotalCount     int
	TeacherStreams TeacherStreams
	Participants   []Participant
}

type TeacherStreams struct {
	TeacherID string
	CameraUID int
	ScreenUID int
}

type StartTeacherLiveRequest struct {
	LiveID    string
	TeacherID string
}

type StartTeacherLiveResponse struct {
	LiveID   string
	CourseID string
	Agora    TeacherAgoraConfig
	Streams  TeacherStreamsToken
}

type TeacherAgoraConfig struct {
	AppID         string
	ChannelName   string
	Role          string
	TokenExpireAt time.Time
}

type TeacherStreamsToken struct {
	Camera TeacherStreamToken
	Screen TeacherStreamToken
}

type TeacherStreamToken struct {
	UID      int
	RTCToken string
}

type EndTeacherLiveRequest struct {
	LiveID    string
	TeacherID string
}

type SetReminderRequest struct {
	LiveID string
	UserID string
}

type DeleteReminderRequest struct {
	LiveID string
	UserID string
}

type ReminderResponse struct {
	LiveID            string
	IsReminderEnabled bool
}
