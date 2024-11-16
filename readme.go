package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// FileInfo stores information about each file
type FileInfo struct {
	Name     string
	Path     string
	Language string
}

// FileGroup represents a group of files with the same base name
type FileGroup struct {
	BaseName string
	GoFile   *FileInfo
	PyFile   *FileInfo
	TsFile   *FileInfo
}

// SectionInfo stores information about each directory section
type SectionInfo struct {
	Name       string
	FileGroups map[string]*FileGroup
}

// ignoredPaths contains paths and files to ignore
var ignoredPaths = map[string]bool{
	".git":              true,
	"node_modules":      true,
	".idea":             true,
	".vscode":           true,
	"vendor":            true,
	"__pycache__":       true,
	".gitignore":        true,
	"README.md":         true,
	".env":              true,
	"venv":              true,
	"dist":              true,
	"build":             true,
}

// shouldIgnore checks if a path should be ignored
func shouldIgnore(path string) bool {
	// Get the base name of the path
	base := filepath.Base(path)
	
	// Check if it's in the ignored paths
	if ignoredPaths[base] {
		return true
	}
	
	// Check if it's a hidden file/directory
	if strings.HasPrefix(base, ".") {
		return true
	}
	
	return false
}

// determineLanguage returns the programming language based on file extension
func determineLanguage(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".go":
		return "Go"
	case ".py":
		return "Python"
	case ".ts":
		return "TypeScript"
	default:
		return "Unknown"
	}
}

// getBaseName returns the filename without extension
func getBaseName(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

// generateMarkdown creates the markdown content with tables
func generateMarkdown(sections []SectionInfo) string {
	var md strings.Builder
	md.WriteString("# Project Structure\n\n")

	for _, section := range sections {
		if section.Name == "." {
			md.WriteString("## Root Directory\n\n")
		} else {
			md.WriteString(fmt.Sprintf("## %s\n\n", section.Name))
		}

		// Create table header
		md.WriteString("| File Name | Go | Python | TypeScript |\n")
		md.WriteString("|-----------|----|---------|-----------|\n")

		// Get all base names and sort them
		var baseNames []string
		for baseName := range section.FileGroups {
			baseNames = append(baseNames, baseName)
		}
		sort.Strings(baseNames)

		// Add table rows
		for _, baseName := range baseNames {
			group := section.FileGroups[baseName]
			
			goLink := ""
			pyLink := ""
			tsLink := ""

			if group.GoFile != nil {
				goLink = fmt.Sprintf("[%s](%s)", "✓", group.GoFile.Path)
			}
			if group.PyFile != nil {
				pyLink = fmt.Sprintf("[%s](%s)", "✓", group.PyFile.Path)
			}
			if group.TsFile != nil {
				tsLink = fmt.Sprintf("[%s](%s)", "✓", group.TsFile.Path)
			}

			md.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
				baseName,
				goLink,
				pyLink,
				tsLink))
		}
		md.WriteString("\n")
	}

	return md.String()
}

func main() {
	sections := make([]SectionInfo, 0)
	sectionMap := make(map[string]*SectionInfo)
	
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip ignored directories and files
		if shouldIgnore(path) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		relPath, err := filepath.Rel(currentDir, path)
		if err != nil {
			return err
		}

		dir := filepath.Dir(relPath)
		if !info.IsDir() {
			language := determineLanguage(info.Name())
			if language == "Unknown" {
				return nil
			}

			// Get or create section
			section, exists := sectionMap[dir]
			if !exists {
				section = &SectionInfo{
					Name:       dir,
					FileGroups: make(map[string]*FileGroup),
				}
				sectionMap[dir] = section
				sections = append(sections, *section)
			}

			// Create FileInfo
			fileInfo := &FileInfo{
				Name:     info.Name(),
				Path:     relPath,
				Language: language,
			}

			// Get or create FileGroup
			baseName := getBaseName(info.Name())
			group, exists := section.FileGroups[baseName]
			if !exists {
				group = &FileGroup{BaseName: baseName}
				section.FileGroups[baseName] = group
			}

			// Add file to appropriate language slot
			switch language {
			case "Go":
				group.GoFile = fileInfo
			case "Python":
				group.PyFile = fileInfo
			case "TypeScript":
				group.TsFile = fileInfo
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// Generate and write markdown content
	markdownContent := generateMarkdown(sections)
	err = os.WriteFile("README.md", []byte(markdownContent), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("README.md has been generated successfully!")
}
