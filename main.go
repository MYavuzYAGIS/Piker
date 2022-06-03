package main

import (
	"flag"
	"github.com/fatih/color"
)

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
		h += "-o, --online	Use the online endpoint to use readied API \n"
		h += "-f, --offline 	Use a custom file to search for the results \n"
		h += "-be  --base64 encode the given string\n"
		h += "-bd  --base64 decode the given string\n"
		h += "

	}
}
