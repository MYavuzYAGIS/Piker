package helpers

import (
	"fmt"
	"log"
)

func Erred(e error) {
	if e != nil {
		fmt.Println(e)
		log.Fatal(e)
	}
}
