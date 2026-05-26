package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type oldItem struct {
	Lesson          map[string]interface{}   `json:"lesson"`
	Playback        map[string]interface{}   `json:"playback"`
	CaptionsVersion int                      `json:"captionsVersion"`
	Captions        []map[string]interface{} `json:"captions"`
	VocabularyItems []map[string]interface{} `json:"vocabularyItems"`
}

type newItem struct {
	Content         map[string]interface{}   `json:"content"`
	Playback        map[string]interface{}   `json:"playback"`
	CaptionsVersion int                      `json:"captionsVersion"`
	Captions        []map[string]interface{} `json:"captions"`
	VocabularyItems []map[string]interface{} `json:"vocabularyItems"`
}

var contentTypes = []string{"video", "music", "fairy_tale", "column", "supplement"}

func main() {
	raw, err := os.ReadFile("data/lessons.json")
	if err != nil {
		panic(fmt.Errorf("read lessons.json: %w", err))
	}
	var lessons []oldItem
	if err := json.Unmarshal(raw, &lessons); err != nil {
		panic(fmt.Errorf("parse lessons.json: %w", err))
	}
	result := make([]newItem, 0, len(lessons)*len(contentTypes))
	for _, lesson := range lessons {
		oldID, _ := lesson.Lesson["id"].(string)
		for _, ct := range contentTypes {
			content := copyMap(lesson.Lesson)
			content["type"] = ct
			// Add a type prefix to the ID to avoid conflicts between 5 types.
			content["id"] = fmt.Sprintf("%s_%s", ct, oldID)
			// Adding a prefix to the title makes it easier to test and identify
			if title, ok := content["title"].(string); ok {
				content["title"] = fmt.Sprintf("[%s] %s", strings.ToUpper(ct), title)
			}
			result = append(result, newItem{
				Content:         content,
				Playback:        lesson.Playback,
				CaptionsVersion: lesson.CaptionsVersion,
				Captions:        lesson.Captions,
				VocabularyItems: lesson.VocabularyItems,
			})
		}
	}
	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(fmt.Errorf("marshal: %w", err))
	}
	if err := os.WriteFile("data/contents.json", out, 0o600); err != nil {
		panic(fmt.Errorf("write contents.json: %w", err))
	}
	fmt.Printf("✅  generated data/contents.json  (%d entries)\n", len(result))
}

func copyMap(m map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}
