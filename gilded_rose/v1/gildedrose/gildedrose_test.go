package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestUpdateQuality(t *testing.T) {
	type TestCase = struct {
		name          string
		days          int
		item          *gildedrose.CategorizedItem
		expected      *gildedrose.CategorizedItem
		expectedError error
	}
	tests := []TestCase{
		{
			name:          "simple item, 1 day, should lower quality and sellin",
			days:          1,
			item:          gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20}),
			expected:      gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "+5 Dexterity Vest", SellIn: 9, Quality: 19}),
			expectedError: nil,
		},
		{
			name:          "simple item, 2 days, with quality 1 should cap quality at 0",
			days:          2,
			item:          gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Unknown Potion", SellIn: 3, Quality: 1}),
			expected:      gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Unknown Potion", SellIn: 1, Quality: 0}),
			expectedError: nil,
		},
		{
			name:          "simple item, 3 days, past sellin should decrease quality 2x faster",
			days:          3,
			item:          gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "+1 Strength Stone Axe", SellIn: 1, Quality: 20}),
			expected:      gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "+1 Strength Stone Axe", SellIn: -2, Quality: 15}),
			expectedError: nil,
		},
		{
			name:          "aged brie, 3 days, should increase quality",
			days:          3,
			item:          gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Aged Brie", SellIn: 3, Quality: 5}),
			expected:      gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Aged Brie", SellIn: 0, Quality: 8}),
			expectedError: nil,
		},
		{
			name:          "aged brie with 49, 3 days, should cap quality at 50",
			days:          3,
			item:          gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Aged Brie", SellIn: 0, Quality: 49}),
			expected:      gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Aged Brie", SellIn: -3, Quality: 50}),
			expectedError: nil,
		},
		{
			name:          "sulfuras, 3 days, should not decrease sellin or quality",
			days:          3,
			item:          gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: 2, Quality: 80}),
			expected:      gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: 2, Quality: 80}),
			expectedError: nil,
		},
		{
			name:          "backstage pass, 3 days, should increase quality by 2",
			days:          3,
			item:          gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 20}),
			expected:      gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 7, Quality: 26}),
			expectedError: nil,
		},
		{
			name:          "backstage pass, 3 days, should increase quality by 3",
			days:          3,
			item:          gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 20}),
			expected:      gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 2, Quality: 29}),
			expectedError: nil,
		},
		{
			name:          "backstage pass, 3 days, should drop quality to 0 past sellin",
			days:          3,
			item:          gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 2, Quality: 50}),
			expected:      gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: -1, Quality: 0}),
			expectedError: nil,
		},
		{
			name:          "conjured item, 1 day, should lower quality 2x and sellin",
			days:          1,
			item:          gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Conjured Cape", SellIn: 10, Quality: 20}),
			expected:      gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Conjured Cape", SellIn: 9, Quality: 18}),
			expectedError: nil,
		},
		{
			name:          "conjured item, 3 days, should lower quality 2x and sellin",
			days:          3,
			item:          gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Conjured Cape", SellIn: 2, Quality: 20}),
			expected:      gildedrose.NewCategorizedItem(&gildedrose.Item{Name: "Conjured Cape", SellIn: -1, Quality: 12}),
			expectedError: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < tc.days; i++ {
				tc.item.UpdateQuality()
			}

			if tc.expectedError != nil {
				t.Fatalf("Not implemented")
			}
			if !isItemEqual(tc.expected, tc.item) {
				t.Fatalf("Expected %v, got %v", tc.expected, tc.item)
			}
		})
	}
}

func isItemEqual(a *gildedrose.CategorizedItem, b *gildedrose.CategorizedItem) bool {
	return a.Item.Name == b.Item.Name && a.Item.Quality == b.Item.Quality && a.Item.SellIn == b.Item.SellIn && a.Category == b.Category
}

func TestCategorizedItem_Categorize(t *testing.T) {
	type TestCase = struct {
		name             string
		item             *gildedrose.CategorizedItem
		expectedCategory gildedrose.Category
		expectedError    error
	}
	tests := []TestCase{
		{
			name: "should categorize aged brie",
			item: &gildedrose.CategorizedItem{
				Item:     &gildedrose.Item{Name: "Aged Brie", SellIn: 0, Quality: 49},
				Category: 0,
			},
			expectedCategory: gildedrose.AgedBrieCategory,
		},
		{
			name: "should categorize ticket",
			item: &gildedrose.CategorizedItem{
				Item:     &gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 20},
				Category: 0,
			},
			expectedCategory: gildedrose.TicketCategory,
		},
		{
			name: "should categorize legendary",
			item: &gildedrose.CategorizedItem{
				Item:     &gildedrose.Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: 10, Quality: 80},
				Category: 0,
			},
			expectedCategory: gildedrose.LegendaryCategory,
		},
		{
			name: "should categorize standard",
			item: &gildedrose.CategorizedItem{
				Item:     &gildedrose.Item{Name: "Mystery Item", SellIn: 10, Quality: 30},
				Category: 0,
			},
			expectedCategory: gildedrose.StandardCategory,
		},
		{
			name: "should categorize conjured",
			item: &gildedrose.CategorizedItem{
				Item:     &gildedrose.Item{Name: "Conjured Cape", SellIn: 10, Quality: 30},
				Category: 0,
			},
			expectedCategory: gildedrose.ConjuredCategory,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.item.Categorize()

			if tc.expectedCategory != actual {
				t.Fatalf("Expected category %v, got %v", tc.expectedCategory, actual)
			}
		})
	}
}
