package main

import (
	"fmt"
	"math"
)

// savings calculates the remaining money after taxes and expenses.
func savings(grossPay int, taxRate float64, expenses int) int {
	afterTaxPay := int(math.Floor(float64(grossPay) * (1 - taxRate)))
	remaining := afterTaxPay - expenses
	if remaining < 0 {
		return 0
	}
	return remaining
}

// materialWaste calculates how much material is left after running jobs.
func materialWaste(totalMaterial int, materialUnits string, numJobs int, jobConsumption int) string {
	waste := totalMaterial - (numJobs * jobConsumption)
	if waste < 0 {
		waste = 0
	}
	return fmt.Sprintf("%d%s", waste, materialUnits)
}

// interest calculates the final investment value after simple interest.
func interest(principal int, rate float64, periods int) int {
	interestEarned := int(math.Floor(float64(principal) * rate * float64(periods)))
	finalValue := principal + interestEarned
	return finalValue
}

func main() {
	// Example cases to test functions
	fmt.Println(savings(100000, 0.12, 20000))     // Example usage
	fmt.Println(materialWaste(500, "kg", 4, 100)) // Example usage
	fmt.Println(interest(100000, 0.05, 10))       // Example usage
}
