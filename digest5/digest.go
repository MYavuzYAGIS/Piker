package digest5

// TODO -> Read file path from stdin
// TODO -> Verify MD5 format before running lookup
// TODO -> Read string from stdin
// TODO -> Check file extension
// TODO -> maybe a better error handling
// TODO -> enable Flags and coloring schemes
// TODO {DONE} -> enable online lookup

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MYavuzYAGIS/Piker/helpers"
	"github.com/PuerkitoBio/goquery"
)

// Md5StringHasher creates a md5 hash of a given string as argument
func Md5StringHasher(input string) string {
	hash := md5.Sum([]byte(input))
	hashed := hex.EncodeToString(hash[:])
	fmt.Println(hashed)
	return hashed
}

// Takes a file and hashes the entire file
func Md5FileHasher() {
}

// Takes a file, checks all the items in the file online
func Md5lineLookupOnline() {
	dat, err := os.Open("mockData/md5list.pk")
	helpers.Erred(err)
	fmt.Println("opened the file")
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)
	fmt.Println("Scanned the file...")
	fmt.Println("Now Iterating...")
	for fileScanner.Scan() {
		text := fileScanner.Text()
		Md5OnlineLookup(text)

	}
	dat.Close()
}

// Takes a file, checks all the items in the file online
func Md5lineLookupOffline() {
}

// Takes an MD5 hash and looks online db to perform a reverse lookup
func Md5OnlineLookup(text string) {
	url := fmt.Sprint("https://md5.gromweb.com/?md5=", text)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("#content").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		hit := s.Find(".string").Text()
		if len(hit) == 0 {
			helpers.Miss.Println("Hash {", text, "} is not  in our database")
		} else {
			helpers.Hit_color.Println("Hit!", text, "======>", hit)
		}
	})
}
