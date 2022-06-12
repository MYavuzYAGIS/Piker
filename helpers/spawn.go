package helpers

import (
	"flag"
	"fmt"
)

var (
	OnlineFlag     bool
	OfflineFlag    bool
	BaseDecodeFlag bool
	BaseEncodeFlag bool
	Md5Flag        bool
	Sha1Flag       bool
	ListFlag       bool
)

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
		h += "\n"

		h += "Examples:\n"
		h += "  Piker -f -s /tmp/apiresponse.pk\n"
		h += "  Piker -o -l listfile filetocrack.pk\n"
		h += "  Piker -f -m listfile ec55d3e698d289f2afd663725127bace "

		fmt.Println(h)
	}
}
