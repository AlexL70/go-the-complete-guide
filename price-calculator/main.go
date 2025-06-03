package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}
	taxMap := make(map[string]prices.TaxIncludedPriceJob)
	globalFm := filemanager.New("prices.txt", "tax_included_prices.json")

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("prices_with_tax_%.0f.json", taxRate*100))
		// cmd := cmdmanager.New()
		fmt.Printf("Calculating prices with tax rate: %.2f\n", taxRate)
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := job.Process()
		if err != nil {
			fmt.Printf("Error processing job for tax rate %.2f: %v\n", taxRate, err)
			continue
		}
		job.IOManager.WriteResult(job)
		taxMap[fmt.Sprintf("%.2f", taxRate)] = *job
	}
	fmt.Println(taxMap)
	globalFm.WriteResult(taxMap)
}
