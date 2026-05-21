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

// ---------------------------------------------------------------------------
// Fake seed data
// ---------------------------------------------------------------------------

var fakeLessons = []responses.LessonListItemResponse{
	{
		ID:          "starbucks_001",
		Title:       "Starbucks Coffee Vocabulary",
		Subtitle:    "Ariannita la Gringa - Starbucks Coffee Vocabulary",
		Description: "日本觀光指南】來了日本一定要去一次的景點、看一次的風景、吃一次",
		CoverURL:    "https://static.gltjp.com/glt/data/article/21000/20522/20231127_155101_412e7290_w1920.webp",
		DurationMs:  568000,
		Level:       "beginner",
		Category:    "daily-english",
		Tags:        []string{"coffee", "starbucks", "daily-life", "vocabulary"},
		IsFree:      true,
		ViewCount:   0,
		CreatedAt:   "2026-05-20T18:50:00+08:00",
		UpdatedAt:   "2026-05-20T18:50:00+08:00",
	},
	{
		ID:          "starbucks_002",
		Title:       "Starbucks Coffee Vocabulary",
		Subtitle:    "Ariannita la Gringa - Starbucks Coffee Vocabulary",
		Description: "透過星巴克情境學習咖啡點餐與日常英文表達。",
		CoverURL:    "https://cdn.zekkei-japan.jp/images/articles/ef4f61494519692aa5ccb95d213a3e1b.jpg",
		DurationMs:  568000,
		Level:       "beginner",
		Category:    "daily-english",
		Tags:        []string{"coffee", "starbucks", "daily-life", "vocabulary"},
		IsFree:      true,
		ViewCount:   0,
		CreatedAt:   "2026-05-20T18:50:00+08:00",
		UpdatedAt:   "2026-05-20T18:50:00+08:00",
	},
}

// fakeDetailMap lets GetLessonDetail look up any seeded lesson by ID.
var fakeDetailMap = map[string]responses.LessonDetailResponse{
	"starbucks_001": {
		Lesson:          fakeLessons[0],
		CaptionsVersion: 1,
		Playback: responses.LessonPlaybackResponse{
			VideoProvider:      "youtube",
			YoutubeVideoID:     "jhEtBuuYNj4",
			VideoURL:           "https://www.youtube.com/watch?v=jhEtBuuYNj4",
			HlsURL:             "",
			DurationMs:         568000, // TODO: 確認 YouTube 實際片長
			StartAtMs:          0,
			AllowSeek:          true,
			AllowPlaybackSpeed: true,
		},
		Captions: []responses.CaptionResponse{
			{
				ID:        "cap_001",
				SortOrder: 1,
				StartMs:   0,
				EndMs:     6000,
				TextEn:    "Hey guys! It's Ariannita la Gringa and welcome back to my YouTube channel. Can you guys guess",
				TextZhTw:  "嗨大家！我是 Ariannita la Gringa，歡迎回到我的 YouTube 頻道。你們猜得到嗎",
			},
			{
				ID:        "cap_002",
				SortOrder: 2,
				StartMs:   6000,
				EndMs:     12000,
				TextEn:    "where I'm at today? Today I'm at Starbucks and as you can see behind me you can see",
				TextZhTw:  "我今天在哪裡嗎？今天我在星巴克，而且你們可以看到我後面有",
			},
			{
				ID:        "cap_003",
				SortOrder: 3,
				StartMs:   12000,
				EndMs:     18000,
				TextEn:    "the beautiful Starbucks logo that they have. And this logo is famous worldwide. You all",
				TextZhTw:  "他們漂亮的星巴克標誌。這個標誌在全世界都很有名。你們大家",
			},
			{
				ID:        "cap_004",
				SortOrder: 4,
				StartMs:   18000,
				EndMs:     25000,
				TextEn:    "might be wondering... Ariannita why are you at Starbucks? Well today I'm at Starbucks because I",
				TextZhTw:  "可能會想說……Ariannita，妳為什麼在星巴克？其實今天我在星巴克是因為我",
			},
			{
				ID:        "cap_005",
				SortOrder: 5,
				StartMs:   25000,
				EndMs:     31000,
				TextEn:    "want to teach you guys some coffee vocabulary. And you guys might be wondering... Wait it's",
				TextZhTw:  "想教你們一些咖啡相關的單字。你們可能會想說……等等，這不是",
			},
			{
				ID:        "cap_006",
				SortOrder: 6,
				StartMs:   31000,
				EndMs:     35000,
				TextEn:    "really easy to order coffee... Actually, it can be pretty difficult especially at",
				TextZhTw:  "點咖啡很簡單嗎……其實，點咖啡可能滿難的，尤其是在",
			},
			{
				ID:        "cap_007",
				SortOrder: 7,
				StartMs:   35000,
				EndMs:     41000,
				TextEn:    "Starbucks because they have different sizes, different coffee, and different drinks all",
				TextZhTw:  "星巴克，因為他們有不同尺寸、不同咖啡，還有各種不同飲品",
			},
			{
				ID:        "cap_008",
				SortOrder: 8,
				StartMs:   41000,
				EndMs:     49000,
				TextEn:    "together. So come along with me and let's go learn some English vocab inside Starbucks.",
				TextZhTw:  "全部混在一起。所以跟我一起來吧，我們進星巴克學一些英文單字。",
			},
			{
				ID:        "cap_009",
				SortOrder: 9,
				StartMs:   49000,
				EndMs:     55000,
				TextEn:    "Before walking inside this Starbucks store I want to teach you guys some facts about Starbucks.",
				TextZhTw:  "在走進這間星巴克之前，我想先教你們一些關於星巴克的小知識。",
			},
			{
				ID:        "cap_010",
				SortOrder: 10,
				StartMs:   55000,
				EndMs:     66000,
				TextEn:    "The first original Starbucks opened in 1971 in Seattle, Washington. And there are over 33,833",
				TextZhTw:  "第一間星巴克在 1971 年於華盛頓州西雅圖開幕。而全球目前有超過 33,833",
			},
			{
				ID:        "cap_011",
				SortOrder: 11,
				StartMs:   66000,
				EndMs:     72000,
				TextEn:    "stores worldwide. 15,000 of those stores are located right in the United States of America.",
				TextZhTw:  "間門市。其中有 15,000 間門市就位在美國。",
			},
		},
		VocabularyItems: []responses.VocabularyItemResponse{
			{
				ID:             "voc_001",
				CaptionID:      "cap_005",
				StartMs:        25000,
				EndMs:          31000,
				Phrase:         "coffee vocabulary",
				DefinitionEn:   "words and phrases related to coffee",
				DefinitionZhTw: "與咖啡相關的單字和片語",
				NoteZhTw:       "vocabulary 是「單字、詞彙」的意思，可以用在各種主題，例如 coffee vocabulary、business vocabulary。",
				Level:          "A2",
				Examples: []responses.VocabularyExampleResponse{
					{En: "Today I want to teach you some coffee vocabulary.", ZhTw: "今天我想教你一些咖啡相關的單字。"},
				},
			},
			{
				ID:             "voc_002",
				CaptionID:      "cap_006",
				StartMs:        31000,
				EndMs:          35000,
				Phrase:         "pretty difficult",
				DefinitionEn:   "quite difficult or a little difficult",
				DefinitionZhTw: "相當困難、有點難",
				NoteZhTw:       "pretty 在口語裡常用來表示「滿、相當」，不是只有「漂亮」的意思。",
				Level:          "A2",
				Examples: []responses.VocabularyExampleResponse{
					{En: "Ordering coffee can be pretty difficult.", ZhTw: "點咖啡可能會滿難的。"},
				},
			},
			{
				ID:             "voc_003",
				CaptionID:      "cap_008",
				StartMs:        41000,
				EndMs:          49000,
				Phrase:         "come along with me",
				DefinitionEn:   "come with me",
				DefinitionZhTw: "跟我一起來",
				NoteZhTw:       "come along with me 是影片、旅遊、Vlog 裡很常見的自然說法。",
				Level:          "A2",
				Examples: []responses.VocabularyExampleResponse{
					{En: "Come along with me and let's learn some English.", ZhTw: "跟我一起來，我們來學一些英文。"},
				},
			},
			{
				ID:             "voc_004",
				CaptionID:      "cap_009",
				StartMs:        49000,
				EndMs:          55000,
				Phrase:         "facts about Starbucks",
				DefinitionEn:   "information or true details about something",
				DefinitionZhTw: "關於某件事的事實或資訊",
				NoteZhTw:       "facts about something 表示「關於某事的事實／小知識」。",
				Level:          "A2",
				Examples: []responses.VocabularyExampleResponse{
					{En: "I want to teach you some facts about Starbucks.", ZhTw: "我想教你一些關於星巴克的小知識。"},
				},
			},
			{
				ID:             "voc_005",
				CaptionID:      "cap_011",
				StartMs:        66000,
				EndMs:          72000,
				Phrase:         "worldwide",
				DefinitionEn:   "around the world",
				DefinitionZhTw: "全世界、全球",
				NoteZhTw:       "worldwide 可以當副詞或形容詞，表示「全球的／在全世界」。",
				Level:          "A2",
				Examples: []responses.VocabularyExampleResponse{
					{En: "There are Starbucks stores worldwide.", ZhTw: "全球都有星巴克門市。"},
				},
			},
		},
	},
}

// ---------------------------------------------------------------------------
// Handlers
// ---------------------------------------------------------------------------

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
// @Param        lessonId  path      string  true  "Lesson ID (e.g. starbucks_001)"
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
