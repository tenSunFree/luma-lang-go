package responses

// Lesson List
// LessonListItemResponse is one card in the lesson list page.
type LessonListItemResponse struct {
	ID          string   `json:"id" example:"paterson_001"`
	Title       string   `json:"title" example:"FUNDAY Cinephile 電影迷 | Paterson 派特森"`
	Subtitle    string   `json:"subtitle" example:"FUNDAY英語學院 - 專業線上英文教學"`
	Description string   `json:"description" example:"透過電影片段學習自然英文表達。"`
	CoverURL    string   `json:"coverUrl" example:"https://example.com/covers/paterson.jpg"`
	DurationMs  int      `json:"durationMs" example:"279000"`
	Level       string   `json:"level" example:"intermediate"`
	Category    string   `json:"category" example:"cinephile"`
	Tags        []string `json:"tags" example:"movie,poetry,daily-life"`
	IsFree      bool     `json:"isFree" example:"true"`
	ViewCount   int      `json:"viewCount" example:"1166"`
	CreatedAt   string   `json:"createdAt" example:"2026-05-14T22:00:00+08:00"`
	UpdatedAt   string   `json:"updatedAt" example:"2026-05-14T22:00:00+08:00"`
}

// Playback
// LessonPlaybackResponse tells the Flutter player how to initialise.
// For YouTube lessons only YoutubeVideoID is filled;
// for CDN/HLS lessons VideoURL and HlsURL are filled instead.
type LessonPlaybackResponse struct {
	VideoProvider      string `json:"videoProvider" example:"youtube"`
	YoutubeVideoID     string `json:"youtubeVideoId,omitempty" example:"abc123xyz"`
	VideoURL           string `json:"videoUrl,omitempty" example:"https://cdn.example.com/videos/paterson.mp4"`
	HlsURL             string `json:"hlsUrl,omitempty" example:"https://cdn.example.com/videos/paterson.m3u8"`
	DurationMs         int    `json:"durationMs" example:"279000"`
	StartAtMs          int    `json:"startAtMs" example:"0"`
	AllowSeek          bool   `json:"allowSeek" example:"true"`
	AllowPlaybackSpeed bool   `json:"allowPlaybackSpeed" example:"true"`
}

// Captions
// CaptionResponse is a single subtitle entry with its time window.
// Flutter uses startMs/endMs to match the current player position.
type CaptionResponse struct {
	ID        string `json:"id" example:"cap_001"`
	SortOrder int    `json:"sortOrder" example:"1"`
	StartMs   int    `json:"startMs" example:"65000"`
	EndMs     int    `json:"endMs" example:"70000"`
	TextEn    string `json:"textEn" example:"poetry is not in grand moments"`
	TextZhTw  string `json:"textZhTw" example:"詩意不在那些宏大的時刻"`
}

// Vocabulary
// VocabularyExampleResponse is one usage example for a vocabulary item.
type VocabularyExampleResponse struct {
	En   string `json:"en" example:"That loud crash scared the daylights out of me."`
	ZhTw string `json:"zhTw" example:"那一聲巨響把我嚇壞了。"`
}

// VocabularyItemResponse is a key phrase/word tied to a caption time window.
type VocabularyItemResponse struct {
	ID             string                      `json:"id" example:"voc_001"`
	CaptionID      string                      `json:"captionId" example:"cap_003"`
	StartMs        int                         `json:"startMs" example:"20000"`
	EndMs          int                         `json:"endMs" example:"35000"`
	Phrase         string                      `json:"phrase" example:"ordinary ones"`
	DefinitionEn   string                      `json:"definitionEn" example:"normal or everyday moments"`
	DefinitionZhTw string                      `json:"definitionZhTw" example:"平凡的瞬間"`
	NoteZhTw       string                      `json:"noteZhTw" example:"ones 代替前面的 moments，避免重複。"`
	Level          string                      `json:"level" example:"A2"`
	Examples       []VocabularyExampleResponse `json:"examples"`
}

// Detail (full lesson)
// LessonDetailResponse is returned by GET /lessons/:lessonId.
// Flutter enters the playback page with a single API call.
type LessonDetailResponse struct {
	Lesson          LessonListItemResponse   `json:"lesson"`
	Playback        LessonPlaybackResponse   `json:"playback"`
	CaptionsVersion int                      `json:"captionsVersion" example:"1"`
	Captions        []CaptionResponse        `json:"captions"`
	VocabularyItems []VocabularyItemResponse `json:"vocabularyItems"`
}
