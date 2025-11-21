package movie_rental

type PricingStrategy interface {
	CalculateCost(daysRented float64) float64
	CalculateFrequentRenterPoints(daysRented float64) uint
}

type RegularPricing struct{}

func (p RegularPricing) CalculateCost(days float64) float64 {
	amount := 2.0
	if days > 2 {
		amount += (days - 2) * 1.5
	}
	return amount
}

func (p RegularPricing) CalculateFrequentRenterPoints(days float64) uint {
	return 1
}

type NewReleasePricing struct{}

func (p NewReleasePricing) CalculateCost(days float64) float64 {
	return days * 3
}

func (p NewReleasePricing) CalculateFrequentRenterPoints(days float64) uint {
	if days > 1 {
		return 2
	}
	return 1
}

type ChildrenPricing struct{}

func (p ChildrenPricing) CalculateCost(days float64) float64 {
	amount := 1.5
	if days > 3 {
		amount += (days - 3) * 1.5
	}
	return amount
}

func (p ChildrenPricing) CalculateFrequentRenterPoints(days float64) uint {
	return 1
}

var strategyMap = map[PriceCode]PricingStrategy{
	REGULAR:     RegularPricing{},
	NEW_RELEASE: NewReleasePricing{},
	CHILDREN:    ChildrenPricing{},
}

func GetPricingStrategy(code PriceCode) PricingStrategy {
	if strategy, ok := strategyMap[code]; ok {
		return strategy
	}
	return RegularPricing{} 
}
