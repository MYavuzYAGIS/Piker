package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

//base64
//md5
//sha1

var (
	base64Success = color.New(color.FgBlue)
	md5Success    = color.New(color.FgCyan)
	shaSuccess    = color.New(color.FgGreen, color.Bold)
	Failue        = color.New(color.FgRed)
	Fatal         = color.New(color.FgRed, color.Bold)
)

const (
	exitOK = iota
	exitOpenFile
	exitReadInput
	exitFetchURL
	exitParseStatements
)

func spawn() {
	flag.Usage = func() {

		h := "Search for a Sha1/md5 (from a file, stdin, or online) or encode/decode Base64 (from a file or stdin)\n\n"
		h += "Usage:\n"
		h += "Piker [OPTIONS] [FILE/URL|-]    \n\n  "

		h += "Options:\n"
		h += "-o, 	--online	Use the online endpoint to use readied API \n"
		h += "-f, 	--offline 	Use a custom file to search for the results \n"
		h += "-be,  --base64 encode the given string\n"
		h += "-bd,  --base64 decode the given string\n"
		h += "-m,   --md5	reverse lookup for  for md5 hash\n"
		h += "-s,   --sha1  reverse lookup for for sha1 hash values\n"
		h += "-l, 	--list  custom dict that contains API endpoints to hit.\n"

		h += "Exit Codes:\n"
		h += fmt.Sprintf("  %d\t%s\n", exitOK, "OK")
		h += fmt.Sprintf("  %d\t%s\n", exitOpenFile, "Failed to open file")
		h += fmt.Sprintf("  %d\t%s\n", exitReadInput, "Failed to read input")
		h += fmt.Sprintf("  %d\t%s\n", exitFetchURL, "Failed to fetch URL")
		h += fmt.Sprintf("  %d\t%s\n", exitParseStatements, "Failed to parse statements")
		h += "\n"

		h += "Examples:\n"
		h += "  Piker -f -s /tmp/apiresponse.pk\n"
		h += "  Piker -o -l listfile filetocrack.pk\n"
		h += "  Piker -f -m listfile ec55d3e698d289f2afd663725127bace "

		fmt.Fprintf(os.Stderr, h)

	}
}

func main() {
	var (
		ungronFlag     bool
		colorizeFlag   bool
		monochromeFlag bool
		streamFlag     bool
		noSortFlag     bool
		versionFlag    bool
		insecureFlag   bool
		jsonFlag       bool
		valuesFlag     bool
	)

	flag.BoolVar(&ungronFlag, "ungron", false, "")
	flag.BoolVar(&ungronFlag, "u", false, "")
	flag.BoolVar(&colorizeFlag, "colorize", false, "")
	flag.BoolVar(&colorizeFlag, "c", false, "")
	flag.BoolVar(&monochromeFlag, "monochrome", false, "")
	flag.BoolVar(&monochromeFlag, "m", false, "")
	flag.BoolVar(&streamFlag, "s", false, "")
	flag.BoolVar(&streamFlag, "stream", false, "")
	flag.BoolVar(&noSortFlag, "no-sort", false, "")
	flag.BoolVar(&versionFlag, "version", false, "")
	flag.BoolVar(&insecureFlag, "k", false, "")
	flag.BoolVar(&insecureFlag, "insecure", false, "")
	flag.BoolVar(&jsonFlag, "j", false, "")
	flag.BoolVar(&jsonFlag, "json", false, "")
	flag.BoolVar(&valuesFlag, "values", false, "")
	flag.BoolVar(&valuesFlag, "value", false, "")
	flag.BoolVar(&valuesFlag, "v", false, "")

	flag.Parse()

}
