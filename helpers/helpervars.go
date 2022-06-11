package helpers

import (
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
	Base64Success = color.New(color.FgBlue)
	Md5Success    = color.New(color.FgCyan)
	ShaSuccess    = color.New(color.FgGreen, color.Bold)
	Failure       = color.New(color.FgRed)
	Fatal         = color.New(color.FgRed, color.Bold)
	Hit_color     = color.New(color.FgGreen, color.Bold)
	Miss          = color.New(color.FgRed, color.Bold)
	BColor        = color.New(color.FgWhite, color.BlinkSlow)
)

var (
	OnlineFlag     bool
	OfflineFlag    bool
	BaseDecodeFlag bool
	BaseEncodeFlag bool
	Md5Flag        bool
	Sha1Flag       bool
	ListFlag       bool
	VersionFlag    bool
)

const (
	ExitOK = iota
	ExitOpenFile
	ExitReadInput
	ExitFetchURL
	ExitParseStatements
)
