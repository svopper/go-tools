package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/svopper/go-tools/utils/bufio"
	"github.com/yeqown/go-qrcode"
)

func main() {

	inputFilePath := os.Args[1]

	lines, err := bufio.ReadLines(inputFilePath)
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
