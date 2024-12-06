package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
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
	if err != nil {
		return nil, err
	}
	return &striverQuestions, nil
}

func createDecksGroupedByHeadStep(snapshot *types.StriverQuestions) error {
	stepsToSkip := []int{}
	for _, deckHeading := range (*snapshot).SheetData {
		if slices.Contains(stepsToSkip, int(deckHeading.StepNo)) {
			log.Printf("Skipping step no: %d\n", deckHeading.StepNo)
			continue
		}

		log.Printf(
			"Processing step no: %d, title: %s\n", 
			deckHeading.StepNo, deckHeading.HeadStepNo,
		)

		for topicIdx, topic := range deckHeading.Topics {
			log.Printf(
				"\t%d. Processing topic no: %d, title: %s\n", 
				topicIdx, topic.SlNoInStep, topic.Title,
			)
		}
	}
	return nil
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

	if err := createDecksGroupedByHeadStep(snapshot); err != nil {
		utils.ErrAndExit(err)
	}

	utils.PPrinter.Success(
		fmt.Sprintf("Decks created successfull at %s", "some/random/folder"),
	)
}


