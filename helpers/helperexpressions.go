package helpers

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func Erred(e error) {
	if e != nil {
		fmt.Println(e)
		log.Fatal(e)
	}
}

func PrintBanner() {
	BColor.Println(Banner)
	BColor.Println("Base64 Encode Decode and  Online/Offline hash reverse lookup tool")
	BColor.Println("Please use --help or -h to see the options")
}

func Spawn() {
	flag.Usage = func() {
		h := "Search for a Sha1/md5 (from a file, stdin, or online) or encode/decode Base64 (from a file or stdin)\n\n"
		h += "Usage:\n"
		h += "Piker [OPTIONS] [FILE/URL|-]    \n\n  "

		h += "Options:\n"
		h += "-o, 	--online	Use the online endpoint to use readied API \n"
		h += "-f, 	--offline 	Use a custom file to search for the results \n"
		h += "-be,  --encode 	encode the given string\n"
		h += "-bd,  --decode 	decode the given string\n"
		h += "-m,   --md5		reverse lookup for  for md5 hash\n"
		h += "-s,   --sha1  	reverse lookup for for sha1 hash values\n"
		h += "-l, 	--list  	custom dict that contains API endpoints to hit.\n"

		h += "Exit Codes:\n"
		h += fmt.Sprintf("  %d\t%s\n", ExitOK, "OK")
		h += fmt.Sprintf("  %d\t%s\n", ExitOpenFile, "Failed to open file")
		h += fmt.Sprintf("  %d\t%s\n", ExitReadInput, "Failed to read input")
		h += fmt.Sprintf("  %d\t%s\n", ExitFetchURL, "Failed to fetch URL")
		h += fmt.Sprintf("  %d\t%s\n", ExitParseStatements, "Failed to parse statements")
		h += "\n"

		h += "Examples:\n"
		h += "  Piker -f -s /tmp/apiresponse.pk\n"
		h += "  Piker -o -l listfile filetocrack.pk\n"
		h += "  Piker -f -m listfile ec55d3e698d289f2afd663725127bace "

		fmt.Fprintf(os.Stderr, h)
	}
}

// func FileIterator(file *os.File) {
// 	dat, err := os.Open("file")
// 	Erred(err)
// 	fmt.Println("opened the file")
// 	fileScanner := bufio.NewScanner(dat)
// 	fileScanner.Split(bufio.ScanLines)
// 	fmt.Println("Scanned the file...")
// 	fmt.Println("Now Iterating...")
// 	for fileScanner.Scan() {
// 		text := fileScanner.Text()
// 		println(text)
// 	}
// }
func FileIterator(filename string) {
	mockPath := "mockData/" + filename
	dat, err := os.Open(mockPath)
	Erred(err)
	fmt.Println("opened the file")
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)
	fmt.Println("Scanned the file...")
	fmt.Println("Now Iterating...")
	for fileScanner.Scan() {
		text := fileScanner.Text()
		println(text)
	}
}
