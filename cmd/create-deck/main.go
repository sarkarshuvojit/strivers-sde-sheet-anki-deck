package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sarkarshuvojit/striver-sde-anki-deck/pkg/types"
	"github.com/sarkarshuvojit/striver-sde-anki-deck/pkg/utils"
)

var (
	souceFilePath string
)

func setupFlags() error {
	flag.StringVar(
		&souceFilePath,
		"path",
		"./assets/snapshot-zero.json",
		"[cmd] --path path-to-snapshot-file.json",
	)
	flag.Parse()

	if !strings.HasSuffix(souceFilePath, ".json") {
		return errors.New(fmt.Sprintf("%s not a valid snapshot", souceFilePath))
	}

	if _, err := os.ReadDir(souceFilePath); err == nil {
		return errors.Join(errors.New(fmt.Sprintf("%s not a valid snapshot", souceFilePath)), err)
	}

	return nil
}

func loadFile(filePath string) (*types.StriverQuestions, error){
	fileConent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
    striverQuestions, err := types.UnmarshalStriverQuestions(fileConent)
	return &striverQuestions, nil
}

func main() {
	if err := setupFlags(); err != nil {
		utils.ErrAndExit(err)
	}
	snapshot, err := loadFile(souceFilePath)
	if err != nil {
		utils.ErrAndExit(err)
	}
	log.Printf("Snapshot: %v\n", snapshot)
	/*
	topics := getTopics()
	for topic := range topics {
		createDeckForTopic(topic)
	}
	*/
}

