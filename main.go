package main

import (
	"flag"
	"os"

	"github.com/MYavuzYAGIS/Piker/helpers"
)

func main() {
	flag.BoolVar(&helpers.OnlineFlag, "online", false, "")
	flag.BoolVar(&helpers.OnlineFlag, "o", false, "")
	flag.BoolVar(&helpers.OfflineFlag, "offline", false, "")
	flag.BoolVar(&helpers.OfflineFlag, "f", false, "")
	flag.BoolVar(&helpers.BaseDecodeFlag, "bd", false, "")
	flag.BoolVar(&helpers.BaseDecodeFlag, "decode", false, "")
	flag.BoolVar(&helpers.BaseEncodeFlag, "be", false, "")
	flag.BoolVar(&helpers.BaseEncodeFlag, "encode", false, "")
	flag.BoolVar(&helpers.Md5Flag, "m", false, "")
	flag.BoolVar(&helpers.Md5Flag, "md5", false, "")
	flag.BoolVar(&helpers.Sha1Flag, "s", false, "")
	flag.BoolVar(&helpers.Sha1Flag, "sha1", false, "")
	flag.BoolVar(&helpers.ListFlag, "l", false, "")
	flag.BoolVar(&helpers.ListFlag, "list", false, "")

	flag.Parse()

	if helpers.VersionFlag {
		// fmt.Printf("Piker version %s\n", PikerVersion)
		os.Exit(helpers.ExitOK)
	}

	// helpers.PrintBanner()
	// shaone.Sha1OnlineLookup("aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d")
	// digest5.Md5OnlineLookup("5d41402abc4b2a76b9719d911017c592")
	// digest5.Md5lineLookupOnline()
	helpers.FileIterator("yavuz.pk")

	// TODO: add argparse
	// TODO: add file extension checker (.pk)
	// TODO : create the easiest function which is base64 encode/decode
	// ====================
	// a := baser.BaseTextDecode("aGVsbG8sd29ybGQhCg==")
	// b := baser.BaseTextEncode(a)
	// fmt.Println(a, b)
	// baser.BaseFileEncode()
	// baser.BaseFileDecode()
	// digest5.Md5StringHasher("yavuz")
	// digest5.Md5OnlineLookup("5d41402abc4b2a76b9719d911017c592")
	// digest5.Md5OnlineLookup("aGVsbG8sd29ybCg==")

	// ========================
}
