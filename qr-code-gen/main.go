package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

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

	outputPath := "out"
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		os.Mkdir(outputPath, os.ModeDir)
	}

	linesLength := len(lines)

	var wg sync.WaitGroup
	wg.Add(linesLength)
	start := time.Now()
	for i := 0; i < linesLength; i++ {
		go func(i int) {
			defer wg.Done()
			qrc, err := qrcode.New(lines[i])
			if err != nil {
				fmt.Printf("could not generate QRCode: %v\n", err)
			}

			// save file
			if err := qrc.Save(fmt.Sprintf("out/%s.png", lines[i])); err != nil {
				fmt.Printf("could not save image: %v\n", err)
			}
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Gen took %s", elapsed)
}
