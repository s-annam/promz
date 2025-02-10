# promz

A Go library for working with the `.promz` file format for prompt data.

## Overview

This library provides functionality for reading, writing, and validating `.promz` files, which are JSON-based files used to store prompt data for the promz.ai platform.  The `.promz` format is designed to be flexible and extensible, allowing for the storage of various types of prompt information, including the prompt text, metadata, examples, and history.

## Installation

```bash
go get [github.com/s-annam/promz](https://github.com/s-annam/promz)
```

## Usage

```go
import (
        "fmt"
        "[github.com/s-annam/promz](https://github.com/s-annam/promz)"
        "encoding/json"
        "os"
)

func main() {
        // Example: Reading a.promz file
        file, err:= os.ReadFile("my_prompt.promz")
        if err!= nil {
                panic(err)
        }

    var p promz.Promz // Define a Promz struct
    err = json.Unmarshal(file, &p) // Unmarshal to Promz struct
    if err!= nil {
        panic(err)
    }

        fmt.Println(p.Metadata.Title) // Accessing the title

    // Example: Creating a new Promz struct
    newPromz:= promz.Promz{
        Version: "1.0",
        Metadata: promz.Metadata{
            Title:       "My New Prompt",
            Description: "A description of my prompt.",
            Tags:       string{"ai", "writing"},
            Author:      "Your Name",
        },
        Content: promz.Content{
            Prompt: "This is the actual prompt text.",
            Examples:promz.Example{
                {Input: "Input example", Output: "Expected output"},
            },
        },
    }

    // Example: Marshaling to JSON and writing to a file
    jsonData, err:= json.MarshalIndent(newPromz, "", "  ") // Use MarshalIndent for pretty printing
    if err!= nil {
        panic(err)
    }

    err = os.WriteFile("new_prompt.promz", jsonData, 0644) // Write the file
    if err!= nil {
        panic(err)
    }
}
```

## .promz File Format

The .promz file format is a JSON-based structure.  A .promz file contains a JSON object with the following structure:

```json
{
  "version": "1.0",
  "metadata": {
    "title": "Prompt Title",
    "description": "A brief description of the prompt.",
    "tags": ["tag1", "tag2"],
    "author": "User ID or Name",
    "created_at": "2024-11-22T12:00:00Z",
    "updated_at": "2024-11-22T12:00:00Z"
  },
  "content": {
    "prompt": "The actual prompt text goes here.",
    "examples": [
      {"input": "Input example 1", "output": "Expected output 1"}
    ],
    "variables": [
      {"name": "topic", "description": "The topic of the prompt"}
    ],
        "apps": [
            "app1",
            "app2"
        ]
  },
  "history": [
    {
      "version": "0.9",
      "updated_at": "2024-11-21T10:00:00Z",
      "changes": "Initial draft of the prompt."
    }
  ]
}
```

## MIME Type

The MIME type for .promz files is application/vnd.promz+json.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This library is licensed under the MIT License.
