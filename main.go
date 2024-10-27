package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func organizeFolder(folder string) error {
	// File types and extensions
	fileTypes := map[string][]string{
		"Images":    {".jpeg", ".jpg", ".png", ".gif"},
		"Videos":    {".mp4", ".avi", ".mov"},
		"Documents": {".pdf", ".docx", ".txt"},
		"Archives":  {".zip", ".rar", ".gz"},
		"Windows":   {".exe", ".com"},
		"Torrents":  {".torrent"},
		"Isos":      {".iso"},
	}

	// Create folders if they do not exist
	for folderName := range fileTypes {
		targetFolder := filepath.Join(folder, folderName)
		if _, err := os.Stat(targetFolder); os.IsNotExist(err) {
			if err := os.Mkdir(targetFolder, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create the folder %s: %w", folderName, err)
			}
		}
	}

	// Traverse and move files
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only regular files
		if info.Mode().IsRegular() {
			ext := filepath.Ext(info.Name())
			for folderName, extensions := range fileTypes {
				// Check if the extension matches
				if contains(extensions, ext) {
					targetPath := filepath.Join(folder, folderName, info.Name())
					if err := os.Rename(path, targetPath); err != nil {
						return err
					}
					fmt.Printf("Moved %s to %s\n", info.Name(), folderName)
					break
				}
			}
		}
		return nil
	})
	return err
}

// Contains function to check if an extension is in the slice
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	pathToDownloads := "/home/tloszkov/Downloads"

	if err := organizeFolder(pathToDownloads); err != nil {
		fmt.Println("Error:", err)
	}
}
