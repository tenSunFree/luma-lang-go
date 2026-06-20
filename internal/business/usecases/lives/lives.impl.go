package lives

import (
	"context"
	"fmt"
	"time"

	"github.com/snykk/go-rest-boilerplate/internal/apperror"
)

type usecase struct {
	repo         Repository
	tokenService TokenService
	appID        string
	wsBaseURL    string
}

func NewUsecase(repo Repository, tokenService TokenService, appID, wsBaseURL string) Usecase {
	return &usecase{
		repo:         repo,
		tokenService: tokenService,
		appID:        appID,
		wsBaseURL:    wsBaseURL,
	}
}

func (uc *usecase) ListLiveCourses(ctx context.Context, req ListLiveCoursesRequest) (ListLiveCoursesResponse, error) {
	items, err := uc.repo.ListLiveCourses(ctx, req.Statuses, req.UserID)
	if err != nil {
		return ListLiveCoursesResponse{}, err
	}
	return ListLiveCoursesResponse{Items: items}, nil
}

func (uc *usecase) JoinLive(ctx context.Context, req JoinLiveRequest) (JoinLiveResponse, error) {
	live, err := uc.repo.GetLiveByID(ctx, req.LiveID)
	if err != nil {
		return JoinLiveResponse{}, err
	}
	if live.Status != StatusLive {
		return JoinLiveResponse{}, apperror.Forbidden("live is not joinable")
	}

	uid, err := uc.repo.GetNextStudentUID(ctx, req.LiveID)
	if err != nil {
		return JoinLiveResponse{}, apperror.InternalCause(fmt.Errorf("allocate uid: %w", err))
	}

	token, err := uc.tokenService.BuildRTCToken(live.AgoraChannelName, uid, RoleAudience)
	if err != nil {
		return JoinLiveResponse{}, apperror.InternalCause(fmt.Errorf("build student rtc token: %w", err))
	}

	if err := uc.repo.UpsertParticipant(ctx, UpsertParticipantInput{
		LiveID:      req.LiveID,
		UserID:      req.UserID,
		DisplayName: req.DisplayName,
		AvatarURL:   req.AvatarURL,
		AgoraUID:    uid,
		Role:        RoleAudience,
	}); err != nil {
		return JoinLiveResponse{}, err
	}

	return JoinLiveResponse{
		LiveID:   live.LiveID,
		CourseID: live.CourseID,
		Agora: AgoraJoinConfig{
			AppID:         uc.appID,
			ChannelName:   live.AgoraChannelName,
			UID:           uid,
			RTCToken:      token.Token,
			Role:          RoleAudience,
			TokenExpireAt: token.ExpireAt,
		},
		Teacher: TeacherConfig{
			TeacherID: live.TeacherID,
			Name:      live.TeacherName,
			AvatarURL: live.TeacherAvatarURL,
			CameraUID: live.TeacherCameraUID,
			ScreenUID: live.TeacherScreenUID,
		},
		StreamLayout: StreamLayout{
			MainUID:  live.TeacherScreenUID,
			MainType: "teacher_screen",
			PipUID:   live.TeacherCameraUID,
			PipType:  "teacher_camera",
		},
		Chat: ChatConfig{
			Enabled:  true,
			Provider: "websocket",
			RoomID:   "chat_" + live.LiveID,
			WSUrl:    fmt.Sprintf("%s/ws/lives/%s", uc.wsBaseURL, live.LiveID),
		},
		Features: LiveFeatures{
			CanSendMessage:  true,
			CanRaiseHand:    true,
			CanSendReaction: true,
			CanPublishAudio: false,
			CanPublishVideo: false,
		},
	}, nil
}

func (uc *usecase) LeaveLive(ctx context.Context, req LeaveLiveRequest) error {
	return uc.repo.LeaveParticipant(ctx, req.LiveID, req.UserID, req.UID)
}

func (uc *usecase) RenewToken(ctx context.Context, req RenewTokenRequest) (RenewTokenResponse, error) {
	live, err := uc.repo.GetLiveByID(ctx, req.LiveID)
	if err != nil {
		return RenewTokenResponse{}, err
	}
	role := req.Role
	if role == "" {
		role = RoleAudience
	}
	token, err := uc.tokenService.BuildRTCToken(live.AgoraChannelName, req.UID, role)
	if err != nil {
		return RenewTokenResponse{}, apperror.InternalCause(fmt.Errorf("renew rtc token: %w", err))
	}
	return RenewTokenResponse{
		UID:           req.UID,
		RTCToken:      token.Token,
		TokenExpireAt: token.ExpireAt,
	}, nil
}

func (uc *usecase) GetParticipants(ctx context.Context, req GetParticipantsRequest) (GetParticipantsResponse, error) {
	live, err := uc.repo.GetLiveByID(ctx, req.LiveID)
	if err != nil {
		return GetParticipantsResponse{}, err
	}
	items, err := uc.repo.ListParticipants(ctx, req.LiveID)
	if err != nil {
		return GetParticipantsResponse{}, err
	}
	return GetParticipantsResponse{
		TotalCount: len(items),
		TeacherStreams: TeacherStreams{
			TeacherID: live.TeacherID,
			CameraUID: live.TeacherCameraUID,
			ScreenUID: live.TeacherScreenUID,
		},
		Participants: items,
	}, nil
}

func (uc *usecase) StartTeacherLive(ctx context.Context, req StartTeacherLiveRequest) (StartTeacherLiveResponse, error) {
	// 先驗身份再改狀態（修正方案二的邏輯順序錯誤）
	live, err := uc.repo.GetLiveByID(ctx, req.LiveID)
	if err != nil {
		return StartTeacherLiveResponse{}, err
	}
	if live.TeacherID != req.TeacherID {
		return StartTeacherLiveResponse{}, apperror.Forbidden("teacher does not own this live")
	}

	if err := uc.repo.MarkLiveStarted(ctx, req.LiveID); err != nil {
		return StartTeacherLiveResponse{}, err
	}

	cameraToken, err := uc.tokenService.BuildRTCToken(live.AgoraChannelName, live.TeacherCameraUID, RoleBroadcaster)
	if err != nil {
		return StartTeacherLiveResponse{}, apperror.InternalCause(fmt.Errorf("build camera token: %w", err))
	}
	screenToken, err := uc.tokenService.BuildRTCToken(live.AgoraChannelName, live.TeacherScreenUID, RoleBroadcaster)
	if err != nil {
		return StartTeacherLiveResponse{}, apperror.InternalCause(fmt.Errorf("build screen token: %w", err))
	}

	// 把老師自己也記錄進 participants
	_ = uc.repo.UpsertParticipant(ctx, UpsertParticipantInput{
		LiveID: live.LiveID, UserID: live.TeacherID,
		DisplayName: live.TeacherName, AvatarURL: live.TeacherAvatarURL,
		AgoraUID: live.TeacherCameraUID, Role: RoleTeacher,
	})

	return StartTeacherLiveResponse{
		LiveID:   live.LiveID,
		CourseID: live.CourseID,
		Agora: TeacherAgoraConfig{
			AppID:         uc.appID,
			ChannelName:   live.AgoraChannelName,
			Role:          RoleBroadcaster,
			TokenExpireAt: cameraToken.ExpireAt,
		},
		Streams: TeacherStreamsToken{
			Camera: TeacherStreamToken{UID: live.TeacherCameraUID, RTCToken: cameraToken.Token},
			Screen: TeacherStreamToken{UID: live.TeacherScreenUID, RTCToken: screenToken.Token},
		},
	}, nil
}

func (uc *usecase) EndTeacherLive(ctx context.Context, req EndTeacherLiveRequest) error {
	live, err := uc.repo.GetLiveByID(ctx, req.LiveID)
	if err != nil {
		return err
	}
	if live.TeacherID != req.TeacherID {
		return apperror.Forbidden("teacher does not own this live")
	}
	return uc.repo.MarkLiveEnded(ctx, req.LiveID)
}

func (uc *usecase) SetReminder(ctx context.Context, req SetReminderRequest) (ReminderResponse, error) {
	live, err := uc.repo.GetLiveByID(ctx, req.LiveID)
	if err != nil {
		return ReminderResponse{}, err
	}
	remindAt := live.ScheduledStartAt.Add(-10 * time.Minute)
	if err := uc.repo.SetReminder(ctx, req.LiveID, req.UserID, remindAt); err != nil {
		return ReminderResponse{}, err
	}
	return ReminderResponse{LiveID: req.LiveID, IsReminderEnabled: true}, nil
}

func (uc *usecase) DeleteReminder(ctx context.Context, req DeleteReminderRequest) (ReminderResponse, error) {
	if err := uc.repo.DeleteReminder(ctx, req.LiveID, req.UserID); err != nil {
		return ReminderResponse{}, err
	}
	return ReminderResponse{LiveID: req.LiveID, IsReminderEnabled: false}, nil
}
