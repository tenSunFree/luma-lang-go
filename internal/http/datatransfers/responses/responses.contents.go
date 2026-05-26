package responses

// ContentListItemResponse contains information about each card on the list page.
type ContentListItemResponse struct {
	ID          string   `json:"id"          example:"video_starbucks_001"`
	Type        string   `json:"type"        example:"video"`
	Title       string   `json:"title"       example:"FUNDAY Cinephile 電影迷 | The Social Network 社群網戰"`
	Subtitle    string   `json:"subtitle"    example:"FUNDAY英語學院"`
	Description string   `json:"description" example:"透過電影片段學習自然英文表達。"`
	CoverURL    string   `json:"coverUrl"    example:"https://example.com/cover.jpg"`
	DurationMs  int      `json:"durationMs"  example:"568000"`
	Level       string   `json:"level"       example:"beginner"`
	Category    string   `json:"category"    example:"daily-english"`
	Tags        []string `json:"tags"        example:"movie,daily-life,vocabulary"`
	IsFree      bool     `json:"isFree"      example:"true"`
	ViewCount   int      `json:"viewCount"   example:"236"`
	CreatedAt   string   `json:"createdAt"   example:"2026-05-20T18:50:00+08:00"`
	UpdatedAt   string   `json:"updatedAt"   example:"2026-05-20T18:50:00+08:00"`
}

// ContentPlaybackResponse tells the player how to initialize.
type ContentPlaybackResponse struct {
	VideoProvider      string `json:"videoProvider"           example:"youtube"`
	YoutubeVideoID     string `json:"youtubeVideoId,omitempty" example:"abc123xyz"`
	VideoURL           string `json:"videoUrl,omitempty"       example:"https://cdn.example.com/video.mp4"`
	HlsURL             string `json:"hlsUrl,omitempty"         example:"https://cdn.example.com/video.m3u8"`
	DurationMs         int    `json:"durationMs"              example:"568000"`
	StartAtMs          int    `json:"startAtMs"               example:"0"`
	AllowSeek          bool   `json:"allowSeek"               example:"true"`
	AllowPlaybackSpeed bool   `json:"allowPlaybackSpeed"      example:"true"`
}

// CaptionResponse is a single caption entry.
type CaptionResponse struct {
	ID        string `json:"id"        example:"cap_001"`
	SortOrder int    `json:"sortOrder" example:"1"`
	StartMs   int    `json:"startMs"   example:"0"`
	EndMs     int    `json:"endMs"     example:"6000"`
	TextEn    string `json:"textEn"    example:"In 2003,"`
	TextZhTw  string `json:"textZhTw"  example:"2003年"`
}

// VocabularyExampleResponse is a single-word usage example.
type VocabularyExampleResponse struct {
	En   string `json:"en"   example:"This is a useful phrase."`
	ZhTw string `json:"zhTw" example:"這是一個實用片語。"`
}

// VocabularyItemResponse is the keyword corresponding to the subtitle timeline.
type VocabularyItemResponse struct {
	ID             string                      `json:"id"             example:"voc_001"`
	CaptionID      string                      `json:"captionId"      example:"cap_001"`
	StartMs        int                         `json:"startMs"        example:"0"`
	EndMs          int                         `json:"endMs"          example:"6000"`
	Phrase         string                      `json:"phrase"         example:"within two decades"`
	DefinitionEn   string                      `json:"definitionEn"   example:"within twenty years"`
	DefinitionZhTw string                      `json:"definitionZhTw" example:"二十年內"`
	NoteZhTw       string                      `json:"noteZhTw"       example:"常用於描述某段時間內的變化。"`
	Level          string                      `json:"level"          example:"A2"`
	Examples       []VocabularyExampleResponse `json:"examples"`
}

// ContentDetailResponse is the complete information returned by GET /contents/:contentId.
type ContentDetailResponse struct {
	Content         ContentListItemResponse  `json:"content"`
	Playback        ContentPlaybackResponse  `json:"playback"`
	CaptionsVersion int                      `json:"captionsVersion" example:"1"`
	Captions        []CaptionResponse        `json:"captions"`
	VocabularyItems []VocabularyItemResponse `json:"vocabularyItems"`
}
