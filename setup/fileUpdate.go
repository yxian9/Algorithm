package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func listFilesInParent(parentDir string) (map[string]bool, error) {
	filesMap := make(map[string]bool)

	entries, err := os.ReadDir(parentDir)
	if err != nil {
		fmt.Println("failed to list parentDir")
		return nil, err
	}

	for _, fileInfo := range entries {
		if !fileInfo.IsDir() { // only add files
			filesMap[fileInfo.Name()] = true
		}
	}

	return filesMap, nil
}

func listFilesInCurrent(currentDir string) (map[string]bool, error) {
	filesMap := make(map[string]bool)

	err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && filepath.Base(path) == ".git" {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			filesMap[filepath.Base(path)] = true
		}
		return nil
	})

	return filesMap, err
}

// Function to copy a file from source to destination
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	parentDir := filepath.Dir(currentDir)

	parentFiles, err := listFilesInParent(parentDir)
	if err != nil {
		fmt.Println("Error listing files in parent directory:", err)
		return
	}

	currentFiles, err := listFilesInCurrent(currentDir)
	if err != nil {
		fmt.Println("Error listing files in current directory:", err)
		return
	}

	fmt.Printf("%d file in lc detected, %d files in repo detected \n", len(parentFiles), len(currentFiles))
	// Copy files that are in parent directory but not in current directory
	for fileName := range parentFiles {
		if !currentFiles[fileName] {
			srcPath := filepath.Join(parentDir, fileName)
			dstPath := filepath.Join(currentDir, fileName)
			if err := copyFile(srcPath, dstPath); err != nil {
				fmt.Println("Error copying file:", err)
			} else {
				fmt.Printf("Copied %s to %s \n", fileName, currentDir)
			}
		}
	}
}
