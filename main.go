package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define the root path and ignore list
	rootPath := "test_directory"
	ignoreList := []string{
		"ignore_this_file.txt",
		"ignore_this_directory",
		"accept_this_directory/ignore_this_thing.txt",
	}

	// Call the function to generate the Markdown snapshot
	err := generateMarkdownSnapshot(rootPath, ignoreList)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// generateMarkdownSnapshot creates a Markdown file documenting the directory contents
func generateMarkdownSnapshot(rootPath string, ignoreList []string) error {
	outputFile, err := os.Create("output.md")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	// Function to determine if the path should be ignored
	shouldIgnore := func(path string) bool {
		// Normalize the path to use forward slashes for consistent handling
		normalizedPath := filepath.ToSlash(path)

		for _, ignore := range ignoreList {
			// Normalize the ignore pattern to use forward slashes
			normalizedIgnore := filepath.ToSlash(ignore)

			// Match ignore patterns exactly from the root relative path
			trimmedPath := strings.TrimPrefix(normalizedPath, filepath.ToSlash(rootPath)+"/")
			if trimmedPath == normalizedIgnore || strings.HasPrefix(trimmedPath, normalizedIgnore+"/") {
				return true
			}
		}
		return false
	}

	// Walk the directory tree
	// func WalkDirAndWrite(dir string, info os.FileInfo, err error) error
	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relativePath, err := filepath.Rel(rootPath, path)
		if err != nil {
			return err
		}
		if shouldIgnore(relativePath) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if !info.IsDir() {
			fileContent, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			// Write the path and file content to the Markdown file
			fmt.Fprintf(writer, "### %s\n```\n%s\n```\n\n", relativePath, fileContent)
		}
		return nil
	})

	return err
}
