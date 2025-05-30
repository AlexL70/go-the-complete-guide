package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, fmt.Errorf("Error reading file: %v", err)
	}
	var value float64
	value, err = strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, fmt.Errorf("Error parsing float from file: %v", err)
	}
	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	balanceStr := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(balanceStr), 0644)
}
