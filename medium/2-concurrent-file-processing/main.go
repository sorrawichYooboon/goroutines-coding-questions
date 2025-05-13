package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Write a Go program that reads multiple files concurrently using Go routines.
// Each Go routine should count the number of words in a file and send the result back to the main function.
// The main function should then sum up the word counts from all files and print the total.

// Hint: Use channels to send the word count results back to the main function.

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}

	var wg sync.WaitGroup
	ch := make(chan int, len(files))

	for _, file := range files {
		wg.Add(1)
		go countWords(file, ch, &wg)
	}

	wg.Wait()
	close(ch)

	total := 0
	for c := range ch {
		total += c
	}

	fmt.Println("Total word count:", total)
}

func countWords(filename string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error:", err)
		ch <- 0
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}
	ch <- count
}
