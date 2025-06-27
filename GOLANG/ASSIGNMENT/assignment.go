package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func ProcessLogs(inputFiles []string, outputFile string) error {
	errors := make(chan string, 100)
	var wg sync.WaitGroup
	writeDone := make(chan struct{})
	go func() {
		defer close(writeDone)
		out, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("failed to create output file: %v\n", err)
			return
		}
		defer out.Close()
		for line := range errors {
			out.WriteString(line + "\n")
		}
	}()

	for _, file := range inputFiles {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			f, err := os.Open(filename)
			if err != nil {
				fmt.Printf("failed to open %s: %v\n", filename, err)
				return
			}
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "ERROR") {
					errors <- line
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("error scanning %s: %v\n", filename, err)
			}
		}(file)
	}

	go func() {
		wg.Wait()
		close(errors)
	}()

	<-writeDone
	return nil
}

func main() {
	inputFiles := []string{"server1.log", "server2.log", "server3.log"}
	if err := ProcessLogs(inputFiles, "errors.log"); err != nil {
		fmt.Printf("Error processing logs: %v", err)
	}
}
