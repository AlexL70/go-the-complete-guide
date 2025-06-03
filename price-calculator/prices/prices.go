package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {
	strPrices, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println("Error reading prices file:", err)
		return
	}
	if len(strPrices) == 0 {
		fmt.Println("No prices found in the file.")
		return
	}
	job.InputPrices, err = conversion.StringToFloat64(strPrices)
	if err != nil {
		fmt.Println("Error converting prices:", err)
		return
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		total := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", total)
	}
	job.TaxIncludedPrices = result
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	job := &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
	return job
}
