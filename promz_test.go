package promz

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	// Create a temporary .promz file for testing
	fileContent := `{
		"version": "1.0",
		"metadata": {
			"title": "My Example Prompt",
			"description": "A prompt for generating creative text.",
			"tags": ["writing", "creativity", "ai"],
			"author": "Your Name",
			"created_at": "2024-11-23T10:00:00Z",
			"updated_at": "2024-11-23T10:00:00Z"
		},
		"content": {
			"prompt": "Write a short story about a robot who learns to love.",
			"examples": [
				{"input": "", "output": "In a world of steel and circuits..."}
			],
			"variables": [
				{"name": "genre", "description": "The genre of the story"}
			],
			"apps": [
				"app1",
				"app2"
			]
		},
		"history": [
			{
				"version": "0.9",
				"updated_at": "2024-11-22T09:00:00Z",
				"changes": "Initial draft."
			}
		]
	}`

	tmpFile, err := os.CreateTemp("", "*.promz")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(fileContent)); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	// Test the Read function
	promz, err := Read(tmpFile.Name())
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}

	// Validate the parsed data
	if promz.Version != "1.0" {
		t.Errorf("expected version 1.0, got %s", promz.Version)
	}
	if promz.Metadata.Title != "My Example Prompt" {
		t.Errorf("expected title 'My Example Prompt', got %s", promz.Metadata.Title)
	}
	if len(promz.Content.Examples) != 1 || promz.Content.Examples[0].Output != "In a world of steel and circuits..." {
		t.Errorf("unexpected examples: %+v", promz.Content.Examples)
	}
	if len(promz.Content.Apps) != 2 || promz.Content.Apps[0] != "app1" {
		t.Errorf("unexpected apps: %+v", promz.Content.Apps)
	}
	if len(promz.History) != 1 || promz.History[0].Version != "0.9" {
		t.Errorf("unexpected history: %+v", promz.History)
	}
}
