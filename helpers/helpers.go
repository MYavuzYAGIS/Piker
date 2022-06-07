package helpers

import (
	"fmt"

	"github.com/fatih/color"
)

var Banner = `
██████╗ ██╗██╗  ██╗███████╗██████╗ 
██╔══██╗██║██║ ██╔╝██╔════╝██╔══██╗
██████╔╝██║█████╔╝ █████╗  ██████╔╝
██╔═══╝ ██║██╔═██╗ ██╔══╝  ██╔══██╗
██║     ██║██║  ██╗███████╗██║  ██║
╚═╝     ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
 `

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

func PrintBanner() {
	BColor.Println(Banner)
	BColor.Println("Base64 Encode Decode and  Online/Offline hashing/rever lookup tool\n")
	BColor.Println("Please use --help or -h to see the options\n")
}
