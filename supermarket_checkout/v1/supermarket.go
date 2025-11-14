package v1

import (
	"math"
)

// TODO: add CashMachine.Scan(article string) method to scan one or multiple articles decoupled from Checkout
type CashMachine struct {
	Articles map[SKU]PriceByQty
	Ticket   map[SKU]int
}

type SKU rune

type PriceByQty map[int]int

func NewCashMachine() *CashMachine {
	return &CashMachine{
		Articles: map[SKU]PriceByQty{},
		Ticket:   map[SKU]int{},
	}
}

// RegisterArticle enregistre le prix unitaire d'un article
func (cm *CashMachine) RegisterArticle(sku SKU, price int) {
	if cm.Articles[sku] == nil {
		cm.Articles[sku] = PriceByQty{}
	}
	cm.Articles[sku][1] = price
}

// RegisterOffer enregistre une offre spéciale (ex: 3 pour 130)
func (cm *CashMachine) RegisterOffer(sku SKU, qty int, price int) {
	if cm.Articles[sku] == nil {
		cm.Articles[sku] = PriceByQty{}
	}
	cm.Articles[sku][qty] = price
}

func (cm *CashMachine) GetArticleUnitPrice(sku SKU) int {
	return cm.Articles[sku][1]
}

// buildLookupTable construit la table de lookup pour un produit
// Elle contient le prix optimal pour chaque quantité de 0 à maxQty
func (cm *CashMachine) buildLookupTable(sku SKU, maxQty int) []int {
	lookup := make([]int, maxQty+1)
	lookup[0] = 0

	priceByQty := cm.Articles[sku]

	for q := 1; q <= maxQty; q++ {
		minPrice := math.MaxInt

		// Essayer toutes les offres disponibles
		for offerQty, offerPrice := range priceByQty {
			if offerQty <= q {
				// Prix = prix optimal du reste + cette offre
				price := lookup[q-offerQty] + offerPrice
				if price < minPrice {
					minPrice = price
				}
			}
		}

		lookup[q] = minPrice
	}

	return lookup
}

func (cm *CashMachine) Checkout(articles string) int {
	artSlice := []rune(articles)

	// Compter les articles scannés
	for _, art := range artSlice {
		cm.Ticket[SKU(art)]++
	}

	total := 0

	// Calculer le prix pour chaque type d'article
	for sku, qty := range cm.Ticket {
		// Construire la lookup table pour ce produit
		lookup := cm.buildLookupTable(sku, qty)
		// Récupérer le prix optimal pour cette quantité
		total += lookup[qty]
	}

	// Réinitialiser le ticket
	cm.Ticket = map[SKU]int{}
	return total
}
