package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yeqown/go-qrcode"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	inputFilePath := os.Args[1]

	lines, err := readLines(inputFilePath)
	if err != nil {
		panic(err)
	}

	for _, invitationCode := range lines {
		qrc, err := qrcode.New(invitationCode)
		if err != nil {
			fmt.Printf("could not generate QRCode: %v", err)
		}

		// save file
		if err := qrc.Save(fmt.Sprintf("out/%s.png", invitationCode)); err != nil {
			fmt.Printf("could not save image: %v", err)
		}
	}
}
