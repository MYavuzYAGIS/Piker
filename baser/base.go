package baser

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
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

//
//func base_file_encode() string {
//
//}
//
//func base_file_decode() string {
//
//}
