package errorchecker

import "fmt"

func Erred(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
