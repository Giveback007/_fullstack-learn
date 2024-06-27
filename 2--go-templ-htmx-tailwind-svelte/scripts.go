package main

import (
	"fmt"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	// Specify the path of the file you want to delete
	filePath := ".task/checksum/go-watch"

	if !fileExists(filePath) {
		fmt.Println("No Checksum file to delete")
		return
	}

	// Call os.Remove to delete the file
	err := os.Remove(filePath)
	if err != nil {
		// Handle the error if the file could not be deleted
		fmt.Println("Error deleting file:", err)
	} else {
		// Success message
		fmt.Println("Checksum deleted...")
	}
}
