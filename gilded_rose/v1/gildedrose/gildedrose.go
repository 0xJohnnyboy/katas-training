package gildedrose

import (
	"errors"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

type CategorizedItem struct {
	Item     *Item
	Category Category
}

type Category int

const (
	StandardCategory Category = iota
	AgedBrieCategory
	TicketCategory
	LegendaryCategory
	ConjuredCategory
)

var CategoriesMap = map[string]Category{
	"Aged Brie": AgedBrieCategory,
	"Backstage passes to a TAFKAL80ETC concert": TicketCategory,
	"Sulfuras, Hand of Ragnaros":                LegendaryCategory,
	"Conjured Cape":                             ConjuredCategory,
}

var CategoryDoesNotExistError = errors.New("Category does not exist.")

func NewCategorizedItem(item *Item) *CategorizedItem {
	categorized := &CategorizedItem{
		item,
		StandardCategory,
	}

	categorized.Categorize()

	return categorized
}

func (ci *CategorizedItem) Categorize() Category {
	cat, ok := CategoriesMap[ci.Item.Name]

	if ok {
		ci.Category = cat
		return cat
	}

	return StandardCategory
}

func (ci *CategorizedItem) increaseQuality(factor int) {
	ci.Item.Quality += (1 * factor)
	if ci.Item.Quality > 50 {
		ci.Item.Quality = 50
	}
}

func (ci *CategorizedItem) decreaseQuality(factor int) {
	ci.Item.Quality -= (1 * factor)
	if ci.Item.Quality < 0 {
		ci.Item.Quality = 0
	}
}
func (ci *CategorizedItem) decreaseSellIn() {
	ci.Item.SellIn--
}

func (ci *CategorizedItem) UpdateQuality() {
	switch ci.Category {
	case LegendaryCategory:
		return
	case AgedBrieCategory:
		ci.decreaseSellIn()
		ci.increaseQuality(1)
		if ci.Item.SellIn < 0 {
			ci.increaseQuality(1)
		}
	case TicketCategory:
		ci.decreaseSellIn()
		ci.increaseQuality(1)
		if ci.Item.SellIn <= 10 {
			ci.increaseQuality(1)
		}
		if ci.Item.SellIn <= 5 {
			ci.increaseQuality(1)
		}
		if ci.Item.SellIn < 0 {
			ci.Item.Quality = 0
		}
	case ConjuredCategory:
		ci.decreaseSellIn()
		if ci.Item.SellIn < 0 {
			ci.decreaseQuality(4)
		} else {
			ci.decreaseQuality(2)
		}
	default:
		ci.decreaseSellIn()
		ci.decreaseQuality(1)
		if ci.Item.SellIn < 0 {
			ci.decreaseQuality(1)
		}
	}
}

func UpdateQuality(categorizedItems []*CategorizedItem) {
	for i := 0; i < len(categorizedItems); i++ {
		categorizedItems[i].UpdateQuality()

	}
}
