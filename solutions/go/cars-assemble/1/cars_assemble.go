package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const carsGroup = 10
	const groupPrice = 95000
	const individualPrice = 10000
	notGrouped := carsCount % carsGroup
	groupedCars := carsCount - notGrouped
	groupedPrice := (groupedCars / carsGroup) * groupPrice
	notGroupedPrice := notGrouped * individualPrice
	return uint(groupedPrice + notGroupedPrice)
}
