package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/svopper/go-tools/utils/bufio"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(fmt.Sprintf("out/%s", filepath))
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

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
			fromIndex := strings.LastIndex(lines[i], "/")
			filename := string(lines[i][fromIndex+1:])
			err := DownloadFile(filename, lines[i])
			if err != nil {
				panic(err)
			}
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Gen took %s", elapsed)

}
