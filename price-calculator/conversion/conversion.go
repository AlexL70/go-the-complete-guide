package conversion

import (
	"fmt"
	"strconv"
)

func StringToFloat64(strings []string) ([]float64, error) {
	var result []float64
	for _, str := range strings {
		floatNum, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, fmt.Errorf("error converting string '%s' to float64: %w", str, err)
		}
		result = append(result, floatNum)
	}
	return result, nil
}
