package promz

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// Promz struct definition
type Promz struct {
	Version  string    `json:"version"`
	Metadata Metadata  `json:"metadata"`
	Content  Content   `json:"content"`
	History  []History `json:"history"`
}

type Metadata struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Content struct {
	Prompt    string     `json:"prompt"`
	Examples  []Example  `json:"examples"`
	Variables []Variable `json:"variables"`
	Apps      []string   `json:"apps"`
}

type Example struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type Variable struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type History struct {
	Version   string    `json:"version"`
	UpdatedAt time.Time `json:"updated_at"`
	Changes   string    `json:"changes"`
}

// Read reads a .promz file and returns a Promz struct
func Read(filePath string) (*Promz, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var promz Promz
	if err := json.Unmarshal(data, &promz); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &promz, nil
}
