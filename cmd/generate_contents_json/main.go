package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func main() {
	raw, err := os.ReadFile("data/lessons.json")
	if err != nil {
		panic(fmt.Errorf("read lessons.json: %w", err))
	}
	var lessons []oldItem
	if err := json.Unmarshal(raw, &lessons); err != nil {
		panic(fmt.Errorf("parse lessons.json: %w", err))
	}
	result := make([]newItem, 0, len(lessons))
	for _, lesson := range lessons {
		content := copyMap(lesson.Lesson)
		// Read lesson.type to determine content.type
		// If no type is set, the default is "video"
		contentType, _ := content["type"].(string)
		if contentType == "" {
			contentType = "video"
		}
		content["type"] = contentType
		result = append(result, newItem{
			Content:         content,
			Playback:        lesson.Playback,
			CaptionsVersion: lesson.CaptionsVersion,
			Captions:        lesson.Captions,
			VocabularyItems: lesson.VocabularyItems,
		})
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
