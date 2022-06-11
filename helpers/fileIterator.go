package helpers

import (
	"bufio"
	"fmt"
	"os"
)

// Cutting the redundancy in file reading functions.
// Takes the file as input from the enclosed function,
// returns a slice of strings to iterate.
func FileIterator(filename string) []string {
	dat, err := os.Open(filename)
	Erred(err)
	fmt.Println("opened the file")
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)
	fmt.Println("Scanned the file...")
	var lines []string
	for fileScanner.Scan() {
		text := fileScanner.Text()
		lines = append(lines, text)
	}
	dat.Close()
	return lines
}
