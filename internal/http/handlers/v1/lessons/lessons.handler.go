// Package lessons serves the /lessons/* HTTP endpoints.
// This version uses hard-coded fake data (no DB).
// Next step: swap fakeLessons / fakeDetail for a real Usecase call.
package lessons

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snykk/go-rest-boilerplate/internal/http/datatransfers/responses"
	v1 "github.com/snykk/go-rest-boilerplate/internal/http/handlers/v1"
)

// Handler is the lessons-handler aggregate.
// Per-endpoint methods could be split into separate files later
// (lessons.list.go, lessons.detail.go) following the auth handler pattern.
type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

// Fake data
var fakeLessons = []responses.LessonListItemResponse{
	{
		ID:          "paterson_001",
		Title:       "FUNDAY Cinephile 電影迷 | Paterson 派特森",
		Subtitle:    "FUNDAY英語學院 - 專業線上英文教學",
		Description: "透過電影《Paterson》學習日常英文與詩意表達。",
		CoverURL:    "https://example.com/covers/paterson.jpg",
		DurationMs:  279000,
		Level:       "intermediate",
		Category:    "cinephile",
		Tags:        []string{"movie", "poetry", "daily-life"},
		IsFree:      true,
		ViewCount:   1166,
		CreatedAt:   "2026-05-14T22:00:00+08:00",
		UpdatedAt:   "2026-05-14T22:00:00+08:00",
	},
	{
		ID:          "youtube_reality_001",
		Title:       "不要為流量犧牲全部價值",
		Subtitle:    "FUNDAY Chat",
		Description: "透過訪談影片學習 YouTuber、流量、價值相關英文表達。",
		CoverURL:    "https://example.com/covers/youtube-reality.jpg",
		DurationMs:  360000,
		Level:       "intermediate",
		Category:    "chat",
		Tags:        []string{"interview", "youtube", "creator"},
		IsFree:      false,
		ViewCount:   641,
		CreatedAt:   "2026-05-14T22:00:00+08:00",
		UpdatedAt:   "2026-05-14T22:00:00+08:00",
	},
}

// fakeDetailMap lets GetLessonDetail look up any seeded lesson by ID.
var fakeDetailMap = map[string]responses.LessonDetailResponse{
	"paterson_001": {
		Lesson:          fakeLessons[0],
		CaptionsVersion: 1,
		Playback: responses.LessonPlaybackResponse{
			VideoProvider:      "youtube",
			YoutubeVideoID:     "abc123xyz",
			DurationMs:         279000,
			StartAtMs:          0,
			AllowSeek:          true,
			AllowPlaybackSpeed: true,
		},
		Captions: []responses.CaptionResponse{
			{ID: "cap_001", SortOrder: 1, StartMs: 0, EndMs: 8000,
				TextEn: "The movie teaches us:", TextZhTw: "這部電影告訴我們："},
			{ID: "cap_002", SortOrder: 2, StartMs: 8000, EndMs: 20000,
				TextEn: "poetry is not in grand moments —", TextZhTw: "詩意不在那些宏大的時刻"},
			{ID: "cap_003", SortOrder: 3, StartMs: 20000, EndMs: 35000,
				TextEn: "it's in noticing ordinary ones.", TextZhTw: "而是在留心那些平凡的瞬間"},
			{ID: "cap_004", SortOrder: 4, StartMs: 35000, EndMs: 50000,
				TextEn: "A key line in the film's spirit is:", TextZhTw: "電影精神裡有一句很重要的話："},
			{ID: "cap_005", SortOrder: 5, StartMs: 50000, EndMs: 65000,
				TextEn: "We're all just walking each other home.", TextZhTw: "我們都只是在陪彼此回家"},
		},
		VocabularyItems: []responses.VocabularyItemResponse{
			{
				ID:             "voc_001",
				CaptionID:      "cap_002",
				StartMs:        8000,
				EndMs:          20000,
				Phrase:         "grand moments",
				DefinitionEn:   "important, impressive, or dramatic moments",
				DefinitionZhTw: "宏大、壯觀或戲劇性的時刻",
				NoteZhTw:       "grand 在這裡不是單純指大，而是帶有壯觀、重要、正式的感覺。",
				Level:          "B1",
				Examples: []responses.VocabularyExampleResponse{
					{En: "Life is not only about grand moments.", ZhTw: "人生不只是那些宏大的時刻。"},
				},
			},
			{
				ID:             "voc_002",
				CaptionID:      "cap_003",
				StartMs:        20000,
				EndMs:          35000,
				Phrase:         "ordinary ones",
				DefinitionEn:   "normal or everyday moments",
				DefinitionZhTw: "平凡的瞬間",
				NoteZhTw:       "ones 代替前面的 moments，避免重複。",
				Level:          "A2",
				Examples: []responses.VocabularyExampleResponse{
					{En: "Small ordinary moments can be meaningful.", ZhTw: "小小的平凡瞬間也可以很有意義。"},
				},
			},
		},
	},
	"youtube_reality_001": {
		Lesson:          fakeLessons[1],
		CaptionsVersion: 1,
		Playback: responses.LessonPlaybackResponse{
			VideoProvider:      "youtube",
			YoutubeVideoID:     "def456uvw",
			DurationMs:         360000,
			StartAtMs:          0,
			AllowSeek:          true,
			AllowPlaybackSpeed: true,
		},
		Captions: []responses.CaptionResponse{
			{ID: "cap_101", SortOrder: 1, StartMs: 10000, EndMs: 18000,
				TextEn: "Don't sacrifice your values for views.", TextZhTw: "不要為了流量犧牲你的價值觀"},
			{ID: "cap_102", SortOrder: 2, StartMs: 18000, EndMs: 27000,
				TextEn: "The algorithm rewards consistency.", TextZhTw: "演算法獎勵的是持續性"},
			{ID: "cap_103", SortOrder: 3, StartMs: 27000, EndMs: 38000,
				TextEn: "You have to be in it for the long game.", TextZhTw: "你必須做好長期抗戰的準備"},
			{ID: "cap_104", SortOrder: 4, StartMs: 38000, EndMs: 50000,
				TextEn: "Burnout is real, and it sneaks up on you.", TextZhTw: "過勞是真實存在的，而且會悄悄找上你"},
		},
		VocabularyItems: []responses.VocabularyItemResponse{
			{
				ID:             "voc_101",
				CaptionID:      "cap_103",
				StartMs:        27000,
				EndMs:          38000,
				Phrase:         "be in it for the long game",
				DefinitionEn:   "to focus on long-term success rather than quick results",
				DefinitionZhTw: "著眼於長期目標，而非短期成果",
				NoteZhTw:       "long game 指的是需要耐心和策略的長期計畫。",
				Level:          "B2",
				Examples: []responses.VocabularyExampleResponse{
					{En: "Building an audience takes time — you have to be in it for the long game.",
						ZhTw: "建立觀眾群需要時間，你必須做好長期抗戰的準備。"},
				},
			},
			{
				ID:             "voc_102",
				CaptionID:      "cap_104",
				StartMs:        38000,
				EndMs:          50000,
				Phrase:         "sneak up on you",
				DefinitionEn:   "to happen gradually without you noticing until it's too late",
				DefinitionZhTw: "在你不知不覺中悄悄逼近",
				NoteZhTw:       "常用來描述問題或感受慢慢積累，直到難以忽視。",
				Level:          "B1",
				Examples: []responses.VocabularyExampleResponse{
					{En: "Stress can really sneak up on you if you don't take breaks.",
						ZhTw: "如果你不休息，壓力真的會悄悄找上你。"},
				},
			},
		},
	},
}

// Handlers
// ListLessons godoc
// @Summary      影片列表 / Get lesson list
// @Description  Returns all published video lessons for the home/list page.
// @Tags         lessons
// @Produce      json
// @Success      200  {object}  v1.BaseResponse{data=map[string][]responses.LessonListItemResponse}
// @Router       /lessons [get]
func (h Handler) ListLessons(ctx *gin.Context) {
	v1.NewSuccessResponse(ctx, http.StatusOK, "lessons fetched successfully", map[string]interface{}{
		"lessons": fakeLessons,
	})
}

// GetLessonDetail godoc
// @Summary      影片詳情 / Get lesson detail
// @Description  Returns full lesson data: lesson info, playback config, time-synced captions (startMs/endMs), and vocabulary teaching items. Flutter calls this once when entering the playback screen.
// @Tags         lessons
// @Produce      json
// @Param        lessonId  path      string  true  "Lesson ID (e.g. paterson_001)"
// @Success      200       {object}  v1.BaseResponse{data=responses.LessonDetailResponse}
// @Failure      404       {object}  v1.BaseResponse
// @Router       /lessons/{lessonId} [get]
func (h Handler) GetLessonDetail(ctx *gin.Context) {
	lessonID := ctx.Param("lessonId")
	detail, ok := fakeDetailMap[lessonID]
	if !ok {
		v1.NewErrorResponse(ctx, http.StatusNotFound, "lesson not found")
		return
	}
	v1.NewSuccessResponse(ctx, http.StatusOK, "lesson detail fetched successfully", detail)
}
