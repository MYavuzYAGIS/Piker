package baser

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

func baseTextDecode(text string) []byte {

	text = strings.ReplaceAll(text, `\/`, `/`)
	dcd, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dcd))
	return dcd

}

//
//func base_text_encode(text) string {
//	v := validator.New()
//	text = b64{
//		Target: "text",
//	}
//	err := v.Struct(text)
//	for _, e := range err.(validator.ValidationErrors) {
//		fmt.Println(e)
//	}
//
//}
//
//func base_file_encode() string {
//
//}
//
//func base_file_decode() string {
//
//}
