package cars

// Defining key values in constants to avoid magic number appearing in the code.
const completeSuccess = 100.0
const hourInMinutes = 60
const batchProdCount = 10
const batchProdCost = 95000
const individualProdCost = 10000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / completeSuccess)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    workingCarsPerHour := int(float64(productionRate) * (successRate / completeSuccess))
	return workingCarsPerHour / hourInMinutes
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	batchProdTotal := uint(carsCount / batchProdCount) * batchProdCost
    individualProdCost := uint(carsCount % batchProdCount) * individualProdCost
    return batchProdTotal + individualProdCost
}
