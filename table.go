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
	".git":        true,
	".vscode":     true,
	"__pycache__": true,
	".gitignore":  true,
	"README.md":   true,
	".env":        true,
	"venv":        true,
	"dist":        true,
	"build":       true,
}

// shouldIgnore checks if a path should be ignored
func shouldIgnore(path string) bool {
	base := filepath.Base(path)
	if ignoredPaths[base] {
		return true
	}
	if strings.HasPrefix(base, ".") {
		return true
	}
	return false
}

// fileExtension returns the programming language based on file extension
func fileExtension(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	ext = strings.TrimPrefix(ext, ".")
	return ext
}

// getBaseName returns the filename without extension
func getBaseName(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

// wrapText wraps text to specified width maintaining word boundaries
func wrapText(text string, width int) []string {
	if len(text) <= width {
		return []string{text}
	}
	truncated := text[:width-3] + "..."
	return []string{truncated}
}

// generateMarkdown creates the markdown content with tables
func generateMarkdown(sections []SectionInfo) string {
	var md strings.Builder
	md.WriteString("# Algorithm\n\n")

	const fileNameWidth = 80

	for _, section := range sections {
		if section.Name == "." {
			continue
		}

		md.WriteString(fmt.Sprintf("## %s\n\n", section.Name))

		// Write table header
		md.WriteString("|  name | go | py | ts |\n")
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
				goLink = fmt.Sprintf("[✓](%s)", group.GoFile.Path)
			}
			if group.PyFile != nil {
				pyLink = fmt.Sprintf("[✓](%s)", group.PyFile.Path)
			}
			if group.TsFile != nil {
				tsLink = fmt.Sprintf("[✓](%s)", group.TsFile.Path)
			}

			// Wrap the basename if it's too long
			lines := wrapText(baseName, fileNameWidth)
			for i, line := range lines {
				if i == 0 {
					// First line includes the links
					md.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
						line,
						goLink,
						pyLink,
						tsLink))
				} else {
					// Continuation lines only include the wrapped text
					md.WriteString(fmt.Sprintf("| %s |  |  |  |\n",
						line))
				}
			}
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
			language := fileExtension(info.Name())
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
			case "go":
				group.GoFile = fileInfo
			case "py":
				group.PyFile = fileInfo
			case "ts":
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
