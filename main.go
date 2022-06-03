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

var PikerVersion = "dev"

func spawn() {
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
		onlineFlag     bool
		offlineFlag    bool
		baseDecodeFlag bool
		baseEncodeFlag bool
		md5Flag        bool
		sha1Flag       bool
		listFlag       bool
	)

	flag.BoolVar(&onlineFlag, "online", false, "")
	flag.BoolVar(&onlineFlag, "o", false, "")
	flag.BoolVar(&offlineFlag, "offline", false, "")
	flag.BoolVar(&offlineFlag, "f", false, "")
	flag.BoolVar(&baseDecodeFlag, "bd", false, "")
	flag.BoolVar(&baseDecodeFlag, "decode", false, "")
	flag.BoolVar(&baseEncodeFlag, "be", false, "")
	flag.BoolVar(&baseEncodeFlag, "encode", false, "")
	flag.BoolVar(&md5Flag, "m", false, "")
	flag.BoolVar(&md5Flag, "md5", false, "")
	flag.BoolVar(&sha1Flag, "s", false, "")
	flag.BoolVar(&sha1Flag, "sha1", false, "")
	flag.BoolVar(&listFlag, "l", false, "")
	flag.BoolVar(&listFlag, "list", false, "")

	flag.Parse()

	if versionFlag {
		fmt.Printf("Piker version %s\n", PikerVersion)
		os.Exit(exitOK)
	}

}
