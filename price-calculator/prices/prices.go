package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	strPrices, err := job.IOManager.ReadLines()
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

func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	job := &TaxIncludedPriceJob{
		IOManager: io,
		TaxRate:   taxRate,
	}
	return job
}
