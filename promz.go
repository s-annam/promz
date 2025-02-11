package promz

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xeipuuv/gojsonschema"
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

// Write writes a Promz struct to a .promz file
func Write(promz *Promz, filePath string) error {
	data, err := json.MarshalIndent(promz, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

// Validate validates a .promz file against the JSON schema
func Validate(filePath string) error {
	schemaLoader := gojsonschema.NewStringLoader(promzSchema)
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}
	fileLoader := gojsonschema.NewReferenceLoader("file:///" + filepath.ToSlash(absPath))

	result, err := gojsonschema.Validate(schemaLoader, fileLoader)
	if err != nil {
		return fmt.Errorf("failed to validate file: %w", err)
	}

	if !result.Valid() {
		var errs []string
		for _, desc := range result.Errors() {
			errs = append(errs, desc.String())
		}
		return fmt.Errorf("invalid .promz file: %s", strings.Join(errs, ", "))
	}

	return nil
}

const promzSchema = `{
	"$schema": "http://json-schema.org/draft-07/schema#",
	"type": "object",
	"properties": {
		"version": {"type": "string"},
		"metadata": {
			"type": "object",
			"properties": {
				"title": {"type": "string"},
				"description": {"type": "string"},
				"tags": {"type": "array", "items": {"type": "string"}},
				"author": {"type": "string"},
				"created_at": {"type": "string", "format": "date-time"},
				"updated_at": {"type": "string", "format": "date-time"}
			},
			"required": ["title", "description", "tags", "author", "created_at", "updated_at"]
		},
		"content": {
			"type": "object",
			"properties": {
				"prompt": {"type": "string"},
				"examples": {
					"type": "array",
					"items": {
						"type": "object",
						"properties": {
							"input": {"type": "string"},
							"output": {"type": "string"}
						},
						"required": ["input", "output"]
					}
				},
				"variables": {
					"type": "array",
					"items": {
						"type": "object",
						"properties": {
							"name": {"type": "string"},
							"description": {"type": "string"}
						},
						"required": ["name", "description"]
					}
				},
				"apps": {"type": "array", "items": {"type": "string"}}
			},
			"required": ["prompt", "examples", "variables", "apps"]
		},
		"history": {
			"type": "array",
			"items": {
				"type": "object",
				"properties": {
					"version": {"type": "string"},
					"updated_at": {"type": "string", "format": "date-time"},
					"changes": {"type": "string"}
				},
				"required": ["version", "updated_at", "changes"]
			}
		}
	},
	"required": ["version", "metadata", "content", "history"]
}`
