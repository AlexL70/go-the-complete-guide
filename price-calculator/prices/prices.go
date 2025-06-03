package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var price float64
		s := scanner.Text()
		if _, err := fmt.Sscanf(s, "%f", &price); err == nil {
			job.InputPrices = append(job.InputPrices, price)
		} else {
			fmt.Println("Error parsing price:", s)
			file.Close()
			return
		}
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	file.Close()
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		total := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = total
	}
	job.TaxIncludedPrices = result
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	job := &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
	job.LoadData()
	return job
}
