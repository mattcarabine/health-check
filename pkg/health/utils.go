package health

import (
	"bufio"
	"log"
	"os"
)

func readFileSingleLine(fileLoc string) (string, error) {
	file, err := os.Open(fileLoc)
	if err != nil {
		log.Println(err)
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		if scanner.Scan() {
			return string(scanner.Bytes()), err
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}
	return "", err
}