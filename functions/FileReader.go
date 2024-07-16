package ascii

import (
	"bufio"
	"log"
	"os"
)

func FileReader(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error: failed to open banner file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: failed to read banner file: %v", err)
	}
	return lines
}
