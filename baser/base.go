package baser

import (
	"encoding/base64"
	"errors"
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

	// fmt.Println(string(dcd))
	return string(dcd)

}

//
func BaseTextEncode(text string) string {
	enc := string(base64.StdEncoding.EncodeToString([]byte(text)))
	if strings.Contains(enc, "illegal base64 data at") {
		errorVal := errors.New("error! This is not a valid base64 line")
		return errorVal.Error()

	}
	enc = base64.StdEncoding.EncodeToString([]byte(text))
	encoded := string(enc)
	a := fmt.Sprint(encoded)
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
