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

const NO_LINK_PROVIDED = "No link provided"
const UNTAGGED = "untagged"

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
func getTags(topic types.Topic) []string {
	if topic.QuesTopic == nil {
		return []string{UNTAGGED}
	}
	questionTopics, err := types.UnmarshalQuestionTopics(
		[]byte(*topic.QuesTopic),
	)
	if err != nil {
		return []string{UNTAGGED}
	}
	tags := []string{}
	for _, questionTopic := range questionTopics {
		tags = append(tags, questionTopic.Value)
	}
	return tags
}

func getBackBytes(topic types.Topic) []byte {
	return []byte(fmt.Sprintf(
		`%s`, topic.Title,
	))
}

func getFrontBytes(topic types.Topic) []byte {
	var youtubeLink string
	if topic.YtLink != nil {
		youtubeLink = *topic.YtLink
	} else {
		youtubeLink = NO_LINK_PROVIDED
	}
	return []byte(fmt.Sprintf(
		`%s`, youtubeLink,
	))
}

func createDeckForHeading(deckHeading types.SheetDatum) (*types.Deck, error) {
	items := []types.DeckItem{}
	for topicIdx, topic := range deckHeading.Topics {
		log.Printf(
			"\t%d. Processing topic no: %d, title: %s\n", 
			topicIdx, topic.SlNoInStep, topic.Title,
		)
		_item := types.DeckItem{
			Front: getFrontBytes(topic),
			Back:  getBackBytes(topic),
			Tags:  getTags(topic),
		}
		items = append(items, _item)
	}
	return &types.Deck{
		Meta:  types.DeckMeta{
			Tags:  []string{},
			Title: deckHeading.HeadStepNo,
		},
		Items: items,
	}, nil
}


func createDecksGroupedByHeadStep(snapshot *types.StriverQuestions) error {
	stepsToSkip := []int{}
	decks := map[string][]*types.Deck{}
	for _, deckHeading := range (*snapshot).SheetData {
		if slices.Contains(stepsToSkip, int(deckHeading.StepNo)) {
			log.Printf("Skipping step no: %d\n", deckHeading.StepNo)
			continue
		}

		log.Printf(
			"Processing step no: %d, title: %s\n", 
			deckHeading.StepNo, deckHeading.HeadStepNo,
		)

		headingKey := fmt.Sprintf(
			"%d_%s", deckHeading.StepNo, deckHeading.HeadStepNo,
		)
		headingDeck, err := createDeckForHeading(deckHeading)
		if err != nil {
			return err
		}

		if val, ok := decks[headingKey]; ok {
			val = append(val, headingDeck)
			decks[headingKey] = val
		} else {
			newDeckList := []*types.Deck{headingDeck}
			decks[headingKey] = newDeckList
		}

	}

	log.Println("Decks Created:")
	for deckKey, deckList := range decks {
		log.Printf("%s => %v", deckKey, len(deckList))
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


