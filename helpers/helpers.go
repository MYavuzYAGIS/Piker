package helpers

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	Hit_color = color.New(color.FgGreen, color.Bold)
	Miss      = color.New(color.FgRed, color.Bold)
	BColor    = color.New(color.FgWhite, color.BlinkSlow)
)

func Erred(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

var Banner = `
██████╗ ██╗██╗  ██╗███████╗██████╗ 
██╔══██╗██║██║ ██╔╝██╔════╝██╔══██╗
██████╔╝██║█████╔╝ █████╗  ██████╔╝
██╔═══╝ ██║██╔═██╗ ██╔══╝  ██╔══██╗
██║     ██║██║  ██╗███████╗██║  ██║
╚═╝     ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
 `

func PrintBanner() {
	BColor.Println(Banner)
}
