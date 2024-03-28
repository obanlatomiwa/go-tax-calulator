package main

import (
	// "example.com/tax-calculator/cmdmanager"
	"fmt"

	"example.com/tax-calculator/filemanager"
	"example.com/tax-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.08, 0.7, 0.19}
	for _, rate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", rate * 100))
		// cmdm := cmdmanager.New()
		newPrice := prices.NewTaxIncludedPrice(fm, rate)
		err := newPrice.Process()

		if err != nil {
			fmt.Println(" Could not process request, try again with the correct inputs")
			fmt.Print(err)
		}
	}
}