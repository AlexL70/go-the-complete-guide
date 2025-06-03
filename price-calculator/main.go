package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}
	taxMap := make(map[string]prices.TaxIncludedPriceJob)

	for _, taxRate := range taxRates {
		job := prices.NewTaxIncludedPriceJob(taxRate)
		job.Process()
		fmt.Printf("Tax Rate: %.2f\n", taxRate)
		fmt.Println("Price with taxes:", job.TaxIncludedPrices)
		//filemanager.WriteJSON(job.TaxIncludedPrices, fmt.Sprintf("prices_with_tax_%.0f.json", taxRate*100))
		taxMap[fmt.Sprintf("%.2f", taxRate)] = *job
	}
	fmt.Println(taxMap)
	filemanager.WriteJSON(taxMap, "tax_included_prices.json")
}
