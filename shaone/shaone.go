package shaone

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MYavuzYAGIS/Piker/helpers"
	"github.com/PuerkitoBio/goquery"
)

// Takes an Sha1 hash and looks online db to perform a reverse lookup
func Sha1OnlineLookup(text string) {
	url := fmt.Sprint("https://sha1.gromweb.com/?hash=", text)
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
