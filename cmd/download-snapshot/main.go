package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sarkarshuvojit/striver-sde-anki-deck/utils"
)

var (
	outFolderPath string
)

func setupFlags() error {
	flag.StringVar(
		&outFolderPath,
		"path",
		"./assets/",
		"[cmd] --path folder/to/save/snapshot/",
	)
	flag.Parse()

	if _, err := os.ReadDir(outFolderPath); err != nil {
		return errors.Join(errors.New(fmt.Sprintf("%s not a directory", outFolderPath)), err)
	}

	return nil
}

func getFilename() string {
	currentTime := time.Now()

	timestamp := currentTime.Format("2006-01-02_15-04-05")

	// Return the formatted filename
	return fmt.Sprintf("snapshot-%s.json", timestamp)
}

func errAndExit(err error) {
	utils.PPrinter.Error(err.Error())
	os.Exit(1)
}

func downloadLatestIntoFile(filename string) error {
	// Make HTTP GET request
	url := "https://backend.takeuforward.org/api/sheets/single/strivers_sde_sheet"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch URL: %v", err)
		return err
	}
	defer resp.Body.Close()

	// Create the file
	file, err := os.Create(fmt.Sprintf("%s/%s", outFolderPath, filename))
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
		return err
	}
	defer file.Close()

	// Copy content from response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
		return err
	}
	return nil
}

func main() {
	if err := setupFlags(); err != nil {
		errAndExit(err)
	}
	utils.PPrinter.Info(fmt.Sprintf("Creating snapshot in %s", outFolderPath))
	
	filename := getFilename()
	if err := downloadLatestIntoFile(filename); err != nil {
		errAndExit(err)
	}
}

