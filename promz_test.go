package promz

import (
	"os"
	"testing"
	"time"
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

func TestWrite(t *testing.T) {
	// Create a Promz struct for testing
	promz := &Promz{
		Version: "1.0",
		Metadata: Metadata{
			Title:       "My Example Prompt",
			Description: "A prompt for generating creative text.",
			Tags:        []string{"writing", "creativity", "ai"},
			Author:      "Your Name",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		Content: Content{
			Prompt: "Write a short story about a robot who learns to love.",
			Examples: []Example{
				{Input: "", Output: "In a world of steel and circuits..."},
			},
			Variables: []Variable{
				{Name: "genre", Description: "The genre of the story"},
			},
			Apps: []string{"app1", "app2"},
		},
		History: []History{
			{Version: "0.9", UpdatedAt: time.Now(), Changes: "Initial draft."},
		},
	}

	// Create a temporary file for testing
	tmpFile, err := os.CreateTemp("", "*.promz")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Test the Write function
	if err := Write(promz, tmpFile.Name()); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Read the file back and validate the content
	readPromz, err := Read(tmpFile.Name())
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}

	if readPromz.Version != promz.Version {
		t.Errorf("expected version %s, got %s", promz.Version, readPromz.Version)
	}
	if readPromz.Metadata.Title != promz.Metadata.Title {
		t.Errorf("expected title %s, got %s", promz.Metadata.Title, readPromz.Metadata.Title)
	}
	if len(readPromz.Content.Examples) != len(promz.Content.Examples) || readPromz.Content.Examples[0].Output != promz.Content.Examples[0].Output {
		t.Errorf("unexpected examples: %+v", readPromz.Content.Examples)
	}
	if len(readPromz.Content.Apps) != len(promz.Content.Apps) || readPromz.Content.Apps[0] != promz.Content.Apps[0] {
		t.Errorf("unexpected apps: %+v", readPromz.Content.Apps)
	}
	if len(readPromz.History) != len(promz.History) || readPromz.History[0].Version != promz.History[0].Version {
		t.Errorf("unexpected history: %+v", readPromz.History)
	}
}

func TestValidate(t *testing.T) {
	validFileContent := `{
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

	invalidFileContent := `{
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
		}
	}` // Missing history

	tests := []struct {
		name    string
		content string
		valid   bool
	}{
		{"valid file", validFileContent, true},
		{"invalid file", invalidFileContent, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "*.promz")
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}
			defer os.Remove(tmpFile.Name())

			if _, err := tmpFile.Write([]byte(tt.content)); err != nil {
				t.Fatalf("failed to write to temp file: %v", err)
			}
			if err := tmpFile.Close(); err != nil {
				t.Fatalf("failed to close temp file: %v", err)
			}

			err = Validate(tmpFile.Name())
			if (err == nil) != tt.valid {
				t.Errorf("Validate() error = %v, valid = %v", err, tt.valid)
			}
		})
	}
}

func TestGetPromptText(t *testing.T) {
	promz := Promz{
		Content: Content{
			Prompt: "Write a short story about a robot who learns to love.",
		},
	}

	prompt := GetPromptText(promz)
	if prompt != "Write a short story about a robot who learns to love." {
		t.Errorf("expected prompt 'Write a short story about a robot who learns to love.', got %s", prompt)
	}
}

func TestSetPromptText(t *testing.T) {
	promz := Promz{
		Content: Content{
			Prompt: "",
		},
	}

	SetPromptText(&promz, "Write a short story about a robot who learns to love.")
	if promz.Content.Prompt != "Write a short story about a robot who learns to love." {
		t.Errorf("expected prompt 'Write a short story about a robot who learns to love.', got %s", promz.Content.Prompt)
	}
}

func TestSetPromptTextEmpty(t *testing.T) {
	promz := Promz{
		Content: Content{
			Prompt: "Initial prompt",
		},
	}

	SetPromptText(&promz, "")
	if promz.Content.Prompt != "" {
		t.Errorf("expected empty prompt, got '%s'", promz.Content.Prompt)
	}
}

func TestSetPromptTextMultiline(t *testing.T) {
	promz := Promz{
		Content: Content{
			Prompt: "",
		},
	}

	multilinePrompt := "Line 1\nLine 2\nLine 3"
	SetPromptText(&promz, multilinePrompt)
	if promz.Content.Prompt != multilinePrompt {
		t.Errorf("expected multiline prompt, got %s", promz.Content.Prompt)
	}
}

func TestAddExample(t *testing.T) {
	promz := Promz{
		Content: Content{
			Examples: []Example{},
		},
	}

	example := Example{Input: "input1", Output: "output1"}
	AddExample(&promz, example)

	if len(promz.Content.Examples) != 1 {
		t.Errorf("expected 1 example, got %d", len(promz.Content.Examples))
	}

	if promz.Content.Examples[0] != example {
		t.Errorf("expected example %+v, got %+v", example, promz.Content.Examples[0])
	}
}

func TestGetExamples(t *testing.T) {
	examples := []Example{
		{Input: "input1", Output: "output1"},
		{Input: "input2", Output: "output2"},
	}
	promz := Promz{
		Content: Content{
			Examples: examples,
		},
	}

	result := GetExamples(promz)

	if len(result) != len(examples) {
		t.Errorf("expected %d examples, got %d", len(examples), len(result))
	}

	for i, example := range examples {
		if result[i] != example {
			t.Errorf("expected example %+v, got %+v", example, result[i])
		}
	}
}

func TestAddTag(t *testing.T) {
	promz := Promz{
		Metadata: Metadata{
			Tags: []string{"tag1", "tag2"},
		},
	}

	AddTag(&promz, "tag3")

	if len(promz.Metadata.Tags) != 3 {
		t.Errorf("expected 3 tags, got %d", len(promz.Metadata.Tags))
	}

	AddTag(&promz, "tag2") // Adding duplicate tag

	if len(promz.Metadata.Tags) != 3 {
		t.Errorf("expected 3 tags after adding duplicate, got %d", len(promz.Metadata.Tags))
	}
}

func TestRemoveTag(t *testing.T) {
	promz := Promz{
		Metadata: Metadata{
			Tags: []string{"tag1", "tag2", "tag3"},
		},
	}

	RemoveTag(&promz, "tag2")

	if len(promz.Metadata.Tags) != 2 {
		t.Errorf("expected 2 tags, got %d", len(promz.Metadata.Tags))
	}

	for _, tag := range promz.Metadata.Tags {
		if tag == "tag2" {
			t.Errorf("tag2 should have been removed")
		}
	}

	RemoveTag(&promz, "tag4") // Removing non-existent tag

	if len(promz.Metadata.Tags) != 2 {
		t.Errorf("expected 2 tags after removing non-existent tag, got %d", len(promz.Metadata.Tags))
	}
}
