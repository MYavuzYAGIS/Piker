package baser

// TODO -> Read file path from stdin
// TODO -> Read string from stdin
// TODO -> Check file extension
// TODO -> maybe a better error handling
// TODO -> enable Flags and coloring schemes

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/MYavuzYAGIS/Piker/helpers"
)

func BaseTextDecode(text string) string {
	text = strings.ReplaceAll(text, `\/`, `/`)
	dcd, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Fatal(err)
	}
	return string(dcd)
}

func BaseTextEncode(text string) string {
	enc := base64.StdEncoding.EncodeToString([]byte(text))
	a := fmt.Sprint(enc)
	return a
}

// read from file, encode each line into base 64.
func BaseFileEncode() int {
	dat, err := os.Open("mockData/yavuz.pk")
	helpers.Erred(err)
	fmt.Println("opened the file")
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)
	fmt.Println("Scanned the file...")
	fmt.Println("Now Iterating...")
	for fileScanner.Scan() {
		text := fileScanner.Text()
		encoded := BaseTextEncode(text)
		fmt.Println(encoded)
	}
	dat.Close()
	return 0
}

func BaseFileDecode() int {
	dat, err := os.Open("mockData/yavuzdecoded.pk")
	helpers.Erred(err)
	fmt.Println("opened the file")
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)
	fmt.Println("Scanned the file...")
	fmt.Println("Now Iterating...")
	for fileScanner.Scan() {
		text := fileScanner.Text()
		encoded := BaseTextDecode(text)
		fmt.Println(encoded)
	}
	dat.Close()
	return 0
}
