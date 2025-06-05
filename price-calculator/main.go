package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("prices_with_tax_%.0f.json", taxRate*100))
		// cmd := cmdmanager.New()
		fmt.Printf("Calculating prices with tax rate: %.2f\n", taxRate)
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go job.Process(doneChans[index], errorChans[index])
	}
	for index, taxRate := range taxRates {
		select {
		case err := <-errorChans[index]:
			fmt.Printf("Error processing job for the Tax Rate %.2f: %v\n", taxRate, err)
		case <-doneChans[index]:
		}
	}
}
