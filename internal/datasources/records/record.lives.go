package records

import "time"

// LiveCourse corresponds to the data table fields, excluding calculated fields (viewer_count, is_reminder_enabled).
// Calculated fields are assembled using subqueries at the repository layer and then converted into usecase domain structs.
type LiveCourse struct {
	ID               string     `db:"id"`
	CourseID         string     `db:"course_id"`
	Title            string     `db:"title"`
	Category         string     `db:"category"`
	Level            string     `db:"level"`
	CourseType       string     `db:"course_type"`
	Status           string     `db:"status"`
	ScheduledStartAt time.Time  `db:"scheduled_start_at"`
	StartedAt        *time.Time `db:"started_at"`
	EndedAt          *time.Time `db:"ended_at"`
	TeacherID        string     `db:"teacher_id"`
	TeacherName      string     `db:"teacher_name"`
	TeacherAvatarURL *string    `db:"teacher_avatar_url"`
	ThumbnailURL     *string    `db:"thumbnail_url"`
	TextbookURL      *string    `db:"textbook_url"`
	AgoraChannelName string     `db:"agora_channel_name"`
	TeacherCameraUID int        `db:"teacher_camera_uid"`
	TeacherScreenUID int        `db:"teacher_screen_uid"`
	// The following two are calculated fields added during querying and need to be provided as subqueries in the SELECT statement.
	ViewerCount       int  `db:"viewer_count"`
	IsReminderEnabled bool `db:"is_reminder_enabled"`
}

// LiveParticipant corresponds to the data table field, IsMuted/IsCameraOn are the CASE calculation fields in the query.
type LiveParticipant struct {
	UserID      string  `db:"user_id"`
	DisplayName string  `db:"display_name"`
	AvatarURL   *string `db:"avatar_url"`
	Role        string  `db:"role"`
	AgoraUID    int     `db:"agora_uid"`
	IsMuted     bool    `db:"is_muted"`
	IsCameraOn  bool    `db:"is_camera_on"`
}
