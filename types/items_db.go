package types

// ItemsDatabase is a map type used by the application as a simple
// in memory database.
type ItemsDatabase map[string]Item

// NewItemDatabase returns an ItemsDatabase instance populate with
// two items.
func NewItemDatabase() ItemsDatabase {
	return ItemsDatabase{
		"1": Item{
			ID:     1,
			Name:   "item one",
			Colour: "red",
		},
		"2": Item{
			ID:     2,
			Name:   "item two",
			Colour: "yellow",
		},
	}
}
