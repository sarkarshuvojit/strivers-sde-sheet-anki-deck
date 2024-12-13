package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/sarkarshuvojit/striver-sde-anki-deck/pkg/tmpl"
	"github.com/sarkarshuvojit/striver-sde-anki-deck/pkg/types"
	"github.com/sarkarshuvojit/striver-sde-anki-deck/pkg/utils"
)

const NO_LINK_PROVIDED = "No link provided"
const UNTAGGED = "untagged"
var (
	souceFilePath string
	outputDirPath string
	shouldCreateOutputDir bool
)

func setupFlags() error {
	flag.StringVar(
		&souceFilePath,
		"path",
		"./assets/snapshot-zero.json",
		"Path to the source snapshot JSON file to be processed", 
	)
	flag.StringVar(
		&outputDirPath,
		"output-dir-path",
		"assets/decks-"+utils.NowInFilesafeFormat(),
		"Destination directory path for output files (defaults to a timestamped folder)", 
	)
	flag.BoolVar(
		&shouldCreateOutputDir,
		"create-dir",
		false,
		"Automatically create the output directory if it does not exist",
	)
	flag.Parse()

	if !strings.HasSuffix(souceFilePath, ".json") {
		return errors.New(fmt.Sprintf("%s not a valid snapshot", souceFilePath))
	}

	if _, err := os.ReadDir(souceFilePath); err == nil {
		return errors.Join(errors.New(fmt.Sprintf("%s not a valid snapshot", souceFilePath)), err)
	}

	// TODO: Add additional check to make sure dir is not empty
	// which can be overridden using --force
	if !shouldCreateOutputDir {
		if _, err := os.ReadDir(outputDirPath); err != nil {
			return errors.Join(errors.New(fmt.Sprintf("%s not a directory", outputDirPath)), err)
		}
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
	topicNameEscaped := strings.ReplaceAll(
		topic.HeadStepNo, " ", "-",
	)
	tags := []string{"topic-"+strings.ToLower(topicNameEscaped)}
	for _, questionTopic := range questionTopics {
		tags = append(tags, questionTopic.Value)
	}
	return tags
}

func getBackBytes(topic types.Topic) []byte {
	postLink := getOrDefault(topic.PostLink, NO_LINK_PROVIDED)
	youtubeLink := getOrDefault(topic.YtLink, NO_LINK_PROVIDED)
	return []byte(fmt.Sprintf(
		tmpl.BackHTMLTmpl, youtubeLink, postLink,
	))
}

func getOrDefault(target *string, _default string) string {
	if target != nil {
		return *target
	} else {
		return _default
	}
}

func getFrontBytes(topic types.Topic) []byte {
	lcLink := getOrDefault(topic.LcLink, NO_LINK_PROVIDED)
	gfgLink := getOrDefault(topic.GfgLink, NO_LINK_PROVIDED)

	return []byte(fmt.Sprintf(
		tmpl.FrontHTMLTmpl, topic.Title, lcLink, gfgLink,
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


func createDecksGroupedByHeadStep(snapshot *types.StriverQuestions) (*map[string]*types.Deck, error) {
	stepsToSkip := []int{}
	decks := map[string]*types.Deck{}
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
			return nil, err
		}

		decks[headingKey] = headingDeck
	}

	log.Println("Decks Created:")
	for _, deckValue := range decks {
		log.Printf("%s\n", (*deckValue).Display())
	}
	return &decks, nil
}

func writeDeckToDisk(deck *types.Deck) error {
	filename := fmt.Sprintf(
		"%s-%s.csv",
		strings.ReplaceAll(deck.Meta.Title, " ","_"), utils.NowInFilesafeFormat(),
	)
	absFilePath := fmt.Sprintf("%s/%s", outputDirPath, filename)
	
	// headers := []string{"Front", "Back", "Tags"}
	
	outputMat := [][]string{
		// headers,
	}

	for _, item := range deck.Items {
		outputMat = append(outputMat, []string{
			string(item.Front), string(item.Back), strings.Join(item.Tags, " "),
		})
	}

	outputFile, err := os.Create(absFilePath)
	if err != nil {
		return err
	}

	cw := csv.NewWriter(outputFile)
	cw.WriteAll(outputMat)

	if err := cw.Error(); err != nil {
		return err
	}

	return nil
}

func main() {
	log.SetOutput(io.Discard)
	if err := setupFlags(); err != nil {
		utils.ErrAndExit(err)
	}

	snapshot, err := loadFile(souceFilePath)
	if err != nil {
		utils.ErrAndExit(err)
	}

	decks, err := createDecksGroupedByHeadStep(snapshot)
	if err != nil {
		utils.ErrAndExit(err)
	}

	if err := os.Mkdir(outputDirPath, 0777); err != nil {
		utils.ErrAndExit(err)
	}

	for deckKey, deck := range *decks {
		if err := writeDeckToDisk(deck); err != nil {
			// TODO: if write error break out from loop intead of trying t
			// to write to an imaginary dir
			// also validtiy of the dir should already been checked before 
			// reaching this point
			log.Println(err)
			utils.PPrinter.Warning(fmt.Sprintf(
				"Failed to write to disk: %s", deckKey, 
			))
			
		}
	}

	utils.PPrinter.Success(
		fmt.Sprintf("Decks created successfull at %s", outputDirPath),
	)
}



