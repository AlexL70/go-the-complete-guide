package filemanager

import (
	"bufio"
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
