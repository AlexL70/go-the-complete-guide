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

func (job *TaxIncludedPriceJob) LoadData() error {
	strPrices, err := job.IOManager.ReadLines()
	if err != nil {
		return fmt.Errorf("Error reading prices file: %w", err)
	}
	if len(strPrices) == 0 {
		return fmt.Errorf("No prices found in the file.")
	}
	job.InputPrices, err = conversion.StringToFloat64(strPrices)
	if err != nil {
		return fmt.Errorf("Error converting prices: %w", err)
	}
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errChan chan error) {
	err := job.LoadData()
	if err != nil {
		errChan <- fmt.Errorf("Error loading data: %w", err)
		return
	}
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		total := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", total)
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	job := &TaxIncludedPriceJob{
		IOManager: io,
		TaxRate:   taxRate,
	}
	return job
}
