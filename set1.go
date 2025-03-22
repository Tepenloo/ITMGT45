package main

import (
	"fmt"
	"math"
)

// main
func main() {
	// Test
	fmt.Println(savings(10000, 0.12, 2000))
	fmt.Println(materialWaste(1000, "kg", 4, 100))
	fmt.Println(interest(5500, 0.05, 10))
}

// savings
func savings(grossPay int, taxRate float64, expenses int) int {
	after_tax := int(math.Floor(float64(grossPay) * (1 - taxRate)))
	remaining := after_tax - expenses
	if remaining < 0 {
		return 0
	}
	return remaining
}

// Material waste
func materialWaste(totalMaterial int, materialUnits string, numJobs int, jobConsumption int) string {
	waste := totalMaterial - (numJobs * jobConsumption)
	if waste < 0 {
		waste = 0
	}
	return fmt.Sprintf("%d%s", waste, materialUnits)
}

// Interest
func interest(principal int, rate float64, periods int) int {
	interest_earned := int(math.Floor(float64(principal) * rate * float64(periods)))
	investment := principal + interest_earned
	return investment
}
