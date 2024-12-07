package types

type DeckItem struct {
	Front []byte
	Back []byte
	Tags []string
}

type DeckMeta struct {
	Tags []string
	Title string
}

type Deck struct {
	Meta DeckMeta
	Items []DeckItem
}
