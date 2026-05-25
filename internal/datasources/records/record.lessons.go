package records

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type Lesson struct {
	ID              string          `db:"id"`
	Title           string          `db:"title"`
	Subtitle        string          `db:"subtitle"`
	Description     string          `db:"description"`
	CoverURL        string          `db:"cover_url"`
	DurationMs      int             `db:"duration_ms"`
	Level           string          `db:"level"`
	Category        string          `db:"category"`
	Tags            pq.StringArray  `db:"tags"`
	IsFree          bool            `db:"is_free"`
	ViewCount       int             `db:"view_count"`
	CaptionsVersion int             `db:"captions_version"`
	Playback        json.RawMessage `db:"playback"`
	Captions        json.RawMessage `db:"captions"`
	VocabularyItems json.RawMessage `db:"vocabulary_items"`
	CreatedAt       time.Time       `db:"created_at"`
	UpdatedAt       time.Time       `db:"updated_at"`
}
