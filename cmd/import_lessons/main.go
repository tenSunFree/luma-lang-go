// go run cmd/import_lessons/main.go -file data/lessons.json
// make import-lessons
// Each execution is an upsert (ON CONFLICT DO UPDATE), so rerunning is safe.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"time"

	"github.com/lib/pq"
	"github.com/snykk/go-rest-boilerplate/internal/config"
	"github.com/snykk/go-rest-boilerplate/internal/constants"
	"github.com/snykk/go-rest-boilerplate/internal/datasources/drivers"
	"github.com/snykk/go-rest-boilerplate/internal/datasources/records"
	lessonspostgres "github.com/snykk/go-rest-boilerplate/internal/datasources/repositories/postgres/lessons"
	"github.com/snykk/go-rest-boilerplate/internal/http/datatransfers/responses"
	"github.com/snykk/go-rest-boilerplate/pkg/logger"
)

type LessonImportItem struct {
	Lesson          responses.LessonListItemResponse   `json:"lesson"`
	Playback        responses.LessonPlaybackResponse   `json:"playback"`
	CaptionsVersion int                                `json:"captionsVersion"`
	Captions        []responses.CaptionResponse        `json:"captions"`
	VocabularyItems []responses.VocabularyItemResponse `json:"vocabularyItems"`
}

func init() {
	if err := config.InitializeAppConfig(); err != nil {
		logger.Fatal(err.Error(), logger.Fields{
			constants.LoggerCategory: constants.LoggerCategoryConfig,
		})
	}
}

func main() {
	filePath := flag.String("file", "data/lessons.json", "path to lessons JSON file")
	flag.Parse()
	// Read JSON
	raw, err := os.ReadFile(*filePath)
	if err != nil {
		logger.Fatal("cannot read lessons file: "+err.Error(), logger.Fields{
			constants.LoggerCategory: "import_lessons",
			"path":                   *filePath,
		})
	}
	var items []LessonImportItem
	if err := json.Unmarshal(raw, &items); err != nil {
		logger.Fatal("cannot parse lessons JSON: "+err.Error(), logger.Fields{
			constants.LoggerCategory: "import_lessons",
		})
	}
	logger.Info("lessons file loaded", logger.Fields{
		constants.LoggerCategory: "import_lessons",
		"count":                  len(items),
	})
	// Connect to DB
	db, err := drivers.SetupSQLXPostgres()
	if err != nil {
		logger.Fatal("db connect failed: "+err.Error(), logger.Fields{
			constants.LoggerCategory: "import_lessons",
		})
	}
	defer func() { _ = db.Close() }()
	repo := lessonspostgres.NewLessonRepository(db)
	ctx := context.Background()
	// Upsert Entries
	for _, item := range items {
		rec, err := toRecord(item)
		if err != nil {
			logger.Fatal("cannot convert lesson to record: "+err.Error(), logger.Fields{
				constants.LoggerCategory: "import_lessons",
				"lesson_id":              item.Lesson.ID,
			})
		}
		if err := repo.Upsert(ctx, rec); err != nil {
			logger.Fatal("upsert failed: "+err.Error(), logger.Fields{
				constants.LoggerCategory: "import_lessons",
				"lesson_id":              item.Lesson.ID,
			})
		}
		logger.Info("upserted lesson", logger.Fields{
			constants.LoggerCategory: "import_lessons",
			"lesson_id":              item.Lesson.ID,
		})
	}
	logger.Info("import complete", logger.Fields{
		constants.LoggerCategory: "import_lessons",
		"total":                  len(items),
	})
}

// `toRecord` converts the JSON item into a DB record.
// JSONB fields (playback / captions / vocabularyItems) are directly serialized into
// json.RawMessage, stored in the DB exactly as is, and do not need to be converted again when read.
func toRecord(item LessonImportItem) (records.Lesson, error) {
	playback, err := json.Marshal(item.Playback)
	if err != nil {
		return records.Lesson{}, err
	}
	captions, err := json.Marshal(item.Captions)
	if err != nil {
		return records.Lesson{}, err
	}
	vocabularyItems, err := json.Marshal(item.VocabularyItems)
	if err != nil {
		return records.Lesson{}, err
	}
	createdAt, err := time.Parse(time.RFC3339, item.Lesson.CreatedAt)
	if err != nil {
		return records.Lesson{}, err
	}
	updatedAt, err := time.Parse(time.RFC3339, item.Lesson.UpdatedAt)
	if err != nil {
		return records.Lesson{}, err
	}
	return records.Lesson{
		ID:              item.Lesson.ID,
		Title:           item.Lesson.Title,
		Subtitle:        item.Lesson.Subtitle,
		Description:     item.Lesson.Description,
		CoverURL:        item.Lesson.CoverURL,
		DurationMs:      item.Lesson.DurationMs,
		Level:           item.Lesson.Level,
		Category:        item.Lesson.Category,
		Tags:            pq.StringArray(item.Lesson.Tags),
		IsFree:          item.Lesson.IsFree,
		ViewCount:       item.Lesson.ViewCount,
		CaptionsVersion: item.CaptionsVersion,
		Playback:        json.RawMessage(playback),
		Captions:        json.RawMessage(captions),
		VocabularyItems: json.RawMessage(vocabularyItems),
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}, nil
}
