package lessons

import (
	"encoding/json"

	"github.com/snykk/go-rest-boilerplate/internal/datasources/records"
	"github.com/snykk/go-rest-boilerplate/internal/http/datatransfers/responses"
)

func toListItem(r records.Lesson) responses.LessonListItemResponse {
	return responses.LessonListItemResponse{
		ID:          r.ID,
		Title:       r.Title,
		Subtitle:    r.Subtitle,
		Description: r.Description,
		CoverURL:    r.CoverURL,
		DurationMs:  r.DurationMs,
		Level:       r.Level,
		Category:    r.Category,
		Tags:        []string(r.Tags),
		IsFree:      r.IsFree,
		ViewCount:   r.ViewCount,
		CreatedAt:   r.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   r.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func toDetail(r records.Lesson) (responses.LessonDetailResponse, error) {
	var playback responses.LessonPlaybackResponse
	if err := json.Unmarshal(r.Playback, &playback); err != nil {
		return responses.LessonDetailResponse{}, err
	}
	var captions []responses.CaptionResponse
	if err := json.Unmarshal(r.Captions, &captions); err != nil {
		return responses.LessonDetailResponse{}, err
	}
	var vocabularyItems []responses.VocabularyItemResponse
	if err := json.Unmarshal(r.VocabularyItems, &vocabularyItems); err != nil {
		return responses.LessonDetailResponse{}, err
	}
	return responses.LessonDetailResponse{
		Lesson:          toListItem(r),
		Playback:        playback,
		CaptionsVersion: r.CaptionsVersion,
		Captions:        captions,
		VocabularyItems: vocabularyItems,
	}, nil
}
