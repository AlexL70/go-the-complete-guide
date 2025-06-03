package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open("prices.txt")
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	strPrices := []string{}
	for scanner.Scan() {
		s := scanner.Text()
		strPrices = append(strPrices, s)
	}
	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %w", err)
	}
	return strPrices, nil
}

func WriteJSON(data any, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error creating file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("Error encoding JSON: %w", err)
	}
	return nil
}
