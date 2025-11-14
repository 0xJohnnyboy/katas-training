package v1

import (
	"testing"
)

func TestSupermarket(t *testing.T) {
	t.Run("Should register an article", func(t *testing.T) {
		cm := NewCashMachine()
		expected := 50
		var sku SKU = SKU('A')
		cm.RegisterArticle(sku, expected)
		actual := cm.GetArticleUnitPrice(sku)

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("Should retrieve an article", func(t *testing.T) {
		cm := NewCashMachine()
		expected := 50
		var sku_A SKU = SKU('A')
		var sku_B SKU = SKU('B')
		cm.RegisterArticle(sku_A, expected)
		cm.RegisterArticle(sku_B, expected)

		actual := cm.GetArticleUnitPrice(sku_A)

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	// t.Run("Should retrieve an article", func(t *testing.T) {
	// })

	t.Run("Simple prices", func(t *testing.T) {
		testCases := []struct {
			Desc     string
			Articles string
			Expected int
		}{
			{"No article", "", 0},
			{"50", "A", 50},
			{"20", "C", 20},
			{"15", "D", 15},
			{"50 + 20", "AC", 70},
			{"20 + 15", "CD", 35},
			{"50+20+15+20", "ACDC", 105},
		}

		cm := NewCashMachine()
		cm.RegisterArticle(SKU('A'), 50)
		cm.RegisterArticle(SKU('C'), 20)
		cm.RegisterArticle(SKU('D'), 15)

		for _, tc := range testCases {
			t.Run(tc.Desc, func(t *testing.T) {
				actual := cm.Checkout(tc.Articles)
				if actual != tc.Expected {
					t.Errorf("Expected %d, got %d", tc.Expected, actual)
				}
			})
		}
	})
	t.Run("Simple offers", func(t *testing.T) {
		testCases := []struct {
			Desc     string
			Articles string
			Expected int
		}{
			{"50 + 50", "AA", 100},
			{"130", "AAA", 130},
			{"130 + 50", "AAAA", 180},
			{"130 + 50 + 50", "AAAAA", 230},
			{"130 + 130", "AAAAAA", 260},
			{"30", "B", 30},
			{"45", "BB", 45},
			{"45 + 30", "BBB", 75},
			{"45 + 45", "BBBB", 90},
		}

		cm := NewCashMachine()
		cm.RegisterArticle(SKU('A'), 50)
		cm.RegisterOffer(SKU('A'), 3, 130)
		cm.RegisterArticle(SKU('B'), 30)
		cm.RegisterOffer(SKU('B'), 2, 45)

		for _, tc := range testCases {
			t.Run(tc.Desc, func(t *testing.T) {
				actual := cm.Checkout(tc.Articles)
				if actual != tc.Expected {
					t.Errorf("Expected %d, got %d", tc.Expected, actual)
				}
			})
		}
	})

	t.Run("Offers with random scanning order", func(t *testing.T) {
		testCases := []struct {
			Desc     string
			Articles string
			Expected int
		}{
			{"1: 3A + 2B = 130 + 45", "ABABA", 175},
			{"2: 3A + 2B = 130 + 45", "BABAA", 175},
			{"3: 3A + 2B = 130 + 45", "AABBA", 175},
			{"4: 3A + 2B = 130 + 45", "BAAAB", 175},
		}

		cm := NewCashMachine()
		cm.RegisterArticle(SKU('A'), 50)
		cm.RegisterOffer(SKU('A'), 3, 130)
		cm.RegisterArticle(SKU('B'), 30)
		cm.RegisterOffer(SKU('B'), 2, 45)

		for _, tc := range testCases {
			t.Run(tc.Desc, func(t *testing.T) {
				actual := cm.Checkout(tc.Articles)
				if actual != tc.Expected {
					t.Errorf("Expected %d, got %d", tc.Expected, actual)
				}
			})
		}
	})

	t.Run("Mixed with and without offers", func(t *testing.T) {
		testCases := []struct {
			Desc     string
			Articles string
			Expected int
		}{
			{"130 + 30", "AAAB", 160},
			{"130 + 45", "AAABB", 175},
			{"130 + 45 + 15", "AAABBD", 190},
			{"3A + 2B + D = 130+45+15", "DABABA", 190},
			{"130 + 45 + 20 + 15", "AAABBCD", 210},
		}

		cm := NewCashMachine()
		cm.RegisterArticle(SKU('A'), 50)
		cm.RegisterArticle(SKU('B'), 30)
		cm.RegisterArticle(SKU('D'), 15)
		cm.RegisterArticle(SKU('C'), 20)

		cm.RegisterOffer(SKU('A'), 3, 130)
		cm.RegisterOffer(SKU('B'), 2, 45)

		for _, tc := range testCases {
			t.Run(tc.Desc, func(t *testing.T) {
				actual := cm.Checkout(tc.Articles)
				if actual != tc.Expected {
					t.Errorf("Expected %d, got %d", tc.Expected, actual)
				}
			})
		}
	})

	t.Run("Multiple offers", func (t *testing.T) {
		testCases := []struct {
			Desc     string
			Articles string
			Expected int
		}{
			{"45", "BB", 45},
			{"45 + 30", "BBB", 75},
			{"45 + 45", "BBBB", 90},
			{"100", "BBBBB", 100},
			{"100 + 30", "BBBBBB", 130},
			{"100 + 45", "BBBBBBB", 145},
		}

		cm := NewCashMachine()
		cm.RegisterArticle(SKU('B'), 30)

		cm.RegisterOffer(SKU('B'), 2, 45)
		cm.RegisterOffer(SKU('B'), 5, 100)

		for _, tc := range testCases {
			t.Run(tc.Desc, func(t *testing.T) {
				actual := cm.Checkout(tc.Articles)
				if actual != tc.Expected {
					t.Errorf("Expected %d, got %d", tc.Expected, actual)
				}
			})
		}
	})
}
