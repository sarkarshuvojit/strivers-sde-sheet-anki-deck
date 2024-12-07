package types

import (
	"fmt"
	"strings"
)

type DeckItem struct {
	Front []byte
	Back []byte
	Tags []string
}

func(di DeckItem) Display() string {
	return fmt.Sprintf(
		"\tFront: %s\n\tBack: %s\n\tTags: %s\n",
		string(di.Front), string(di.Back), strings.Join(di.Tags, ","),
	)
}

type DeckMeta struct {
	Tags []string
	Title string
}

type Deck struct {
	Meta DeckMeta
	Items []DeckItem
}

func(d Deck) Display() string {
	itemDisplays := []string{}
	for _, item := range d.Items {
		itemDisplays = append(itemDisplays, item.Display())
	}
	return fmt.Sprintf(
		"Title: %s\n\nContent:\n%s\n",
		d.Meta.Title, strings.Join(itemDisplays, "\n"),
	)
}

