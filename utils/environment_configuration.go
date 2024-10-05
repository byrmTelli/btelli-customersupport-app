package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// To read .env file
func LoadEnvFile(FilePath string) {

	file, err := os.Open(FilePath)

	if err != nil {
		log.Fatalf("Error occured while reading .env file. Error: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Pass empty lines

		if len(line) == 0 {
			continue
		}

		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			log.Printf("Invalid line in file .env. Line: %s", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading .env file. Error: %v", err)
	}
}
