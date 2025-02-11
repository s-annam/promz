package main

import (
	"fmt"
	"log"

	"github.com/s-annam/promz"
)

func main() {
	// Example: Reading a .promz file
	filePath := "example.promz"
	promzData, err := promz.Read(filePath)
	if err != nil {
		log.Fatalf("failed to read .promz file: %v", err)
	}
	fmt.Printf("Read .promz file: %+v\n", promzData)

	// Example: Writing a .promz file
	newFilePath := "new_example.promz"
	if err := promz.Write(promzData, newFilePath); err != nil {
		log.Fatalf("failed to write .promz file: %v", err)
	}
	fmt.Printf("Wrote .promz file to %s\n", newFilePath)

	// Example: Validating a .promz file
	if err := promz.Validate(newFilePath); err != nil {
		log.Fatalf("failed to validate .promz file: %v", err)
	}
	fmt.Printf("Validated .promz file: %s\n", newFilePath)

	// Example: Manipulating prompt text
	promptText := promz.GetPromptText(*promzData)
	fmt.Printf("Prompt text: %s\n", promptText)

	newPromptText := "Write a poem about the sea."
	promz.SetPromptText(promzData, newPromptText)
	fmt.Printf("Updated prompt text: %s\n", promz.GetPromptText(*promzData))

	// Example: Adding and getting examples
	example := promz.Example{Input: "input2", Output: "output2"}
	promz.AddExample(promzData, example)
	examples := promz.GetExamples(*promzData)
	fmt.Printf("Examples: %+v\n", examples)

	// Example: Adding and removing tags
	tag := "new_tag"
	promz.AddTag(promzData, tag)
	fmt.Printf("Tags after adding: %+v\n", promzData.Metadata.Tags)

	promz.RemoveTag(promzData, tag)
	fmt.Printf("Tags after removing: %+v\n", promzData.Metadata.Tags)
}
