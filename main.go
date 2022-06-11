package main

import (
	"flag"

	"github.com/MYavuzYAGIS/Piker/baser"
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

	// digest5.Md5lineLookupOnline("md5list.pk")
	// shaone.Sha1lineLookupOnline("shaonelist.pk")
	baser.BaseFileEncode("yavuz.pk")
	baser.BaseFileDecode("yavuzdecoded.pk")

	// TODO: add argparse
	// TODO: add file extension checker (.pk)
	// TODO : create the easiest function which is base64 encode/decode
}
