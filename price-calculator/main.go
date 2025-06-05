package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("prices_with_tax_%.0f.json", taxRate*100))
		// cmd := cmdmanager.New()
		fmt.Printf("Calculating prices with tax rate: %.2f\n", taxRate)
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go job.Process(doneChans[index])
		// if err != nil {
		// 	fmt.Printf("Error processing job for tax rate %.2f: %v\n", taxRate, err)
		// 	continue
		// }
	}
	for _, doneChan := range doneChans {
		<-doneChan
	}
}
