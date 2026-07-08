package lives

import (
	"context"
	"time"
)

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

// Usecase is the business layer interface; handler layer only depends on this.
type Usecase interface {
	ListLiveCourses(ctx context.Context, req ListLiveCoursesRequest) (ListLiveCoursesResponse, error)
	JoinLive(ctx context.Context, req JoinLiveRequest) (JoinLiveResponse, error)
	LeaveLive(ctx context.Context, req LeaveLiveRequest) error
	RenewToken(ctx context.Context, req RenewTokenRequest) (RenewTokenResponse, error)
	GetParticipants(ctx context.Context, req GetParticipantsRequest) (GetParticipantsResponse, error)
	// Teachers can directly create and start broadcasting without needing a liveId first.
	StartTeacherLive(ctx context.Context, req StartTeacherLiveRequest) (StartTeacherLiveResponse, error)
	// When the teacher ends the live stream, the backend automatically locates the currently running live stream and sends back the liveId of the stream that was ended.
	EndTeacherLive(ctx context.Context, req EndTeacherLiveRequest) (string, error)
	SetReminder(ctx context.Context, req SetReminderRequest) (ReminderResponse, error)
	DeleteReminder(ctx context.Context, req DeleteReminderRequest) (ReminderResponse, error)
}

// Repository is the storage abstraction; implemented by postgres layer.
type Repository interface {
	ListLiveCourses(ctx context.Context, statuses []string, userID string) ([]LiveCourse, error)
	GetLiveByID(ctx context.Context, liveID string) (LiveCourse, error)
	// Teacher sets up a live streaming room
	CreateTeacherLive(ctx context.Context, in CreateTeacherLiveInput) (LiveCourse, error)
	// I tried to find the teacher's current live stream (status=live), but received no response. NotFound error.
	GetActiveLiveByTeacherID(ctx context.Context, teacherID string) (LiveCourse, error)
	// End the teacher's current live stream and send back the ended liveId.
	MarkTeacherActiveLiveEnded(ctx context.Context, teacherID string) (string, error)
	GetNextStudentUID(ctx context.Context, liveID string) (int, error)
	UpsertParticipant(ctx context.Context, in UpsertParticipantInput) error
	LeaveParticipant(ctx context.Context, liveID, userID string, uid int) error
	ListParticipants(ctx context.Context, liveID string) ([]Participant, error)
	MarkLiveStarted(ctx context.Context, liveID string) error
	MarkLiveEnded(ctx context.Context, liveID string) error
	SetReminder(ctx context.Context, liveID, userID string, remindAt time.Time) error
	DeleteReminder(ctx context.Context, liveID, userID string) error
}

// TokenService abstracts Agora token generation.
type TokenService interface {
	BuildRTCToken(channelName string, uid int, role string) (TokenResult, error)
}

type TokenResult struct {
	Token    string
	ExpireAt time.Time
}

// domain structs
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

type CreateTeacherLiveInput struct {
	LiveID           string
	CourseID         string
	Title            string
	Category         string
	Level            string
	CourseType       string
	TeacherID        string
	TeacherName      string
	AvatarURL        *string
	ThumbnailURL     *string
	TextbookURL      *string
	AgoraChannelName string
	TeacherCameraUID int
	TeacherScreenUID int
}

// request / response types
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

// StartTeacherLiveRequest Information that the teacher includes when starting the live stream (no liveId required, automatically generated by the backend)
type StartTeacherLiveRequest struct {
	TeacherID    string
	TeacherName  string
	AvatarURL    *string
	Title        string
	Category     string
	Level        string
	CourseType   string
	ThumbnailURL *string
	TextbookURL  *string
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

// EndTeacherLiveRequest When a teacher ends a live stream, only the TeacherID is needed; the backend will automatically find ongoing live streams.
type EndTeacherLiveRequest struct {
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
