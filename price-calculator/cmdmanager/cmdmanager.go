package cmdmanager

import "fmt"

type CMDManager struct{}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter the prices (one per line, end with an empty line):")
	var lines []string
	for {
		var line string
		fmt.Print("Price: ")
		fmt.Scanln(&line)
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	if len(lines) == 0 {
		return nil, fmt.Errorf("no prices entered")
	}
	return lines, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() *CMDManager {
	return &CMDManager{}
}
