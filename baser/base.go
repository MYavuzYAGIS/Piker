package baser

// TODO -> Read file path from stdin
// TODO -> Read string from stdin
// TODO -> Check file extension
// TODO -> maybe a better error handling
// TODO -> enable Flags and coloring schemes

import (
	"encoding/base64"
	"fmt"
	"log"
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
func BaseFileDecode(filename string) {
	mockPath := "mockData/" + filename
	iterable := helpers.FileIterator(mockPath)
	helpers.Md5Success.Println("Iterating the file for decoding the content")
	for _, iter := range iterable {
		next := BaseTextDecode(iter)
		fmt.Println(next)
	}
}

func BaseFileEncode(filename string) {
	mockPath := "mockData/" + filename
	iterable := helpers.FileIterator(mockPath)
	helpers.Md5Success.Println("Iterating the file for encoding the content")
	for _, iter := range iterable {
		next := BaseTextEncode(iter)
		fmt.Println(next)
	}
}
