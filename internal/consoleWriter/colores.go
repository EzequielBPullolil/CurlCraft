package consolewriter

import (
	"github.com/fatih/color"
)

var (
	Succes        = color.New(color.FgGreen).Add(color.Underline)
	SuccesFc      = Succes.PrintFunc()
	SuccesPl      = Succes.PrintlnFunc()
	SuccesSp      = Succes.SprintFunc()
	Fail          = color.New(color.FgHiWhite, color.BgRed, color.Bold)
	FailSp        = Fail.SprintFunc()
	Redirect      = color.New(color.FgBlue)
	RedirectSp    = Redirect.SprintFunc()
	ServerError   = color.New(color.FgRed)
	ServerErrorSp = ServerError.SprintFunc()
	Info          = color.New(color.FgWhite, color.Bold, color.Underline)
	InfoPL        = Info.PrintlnFunc()
	InfoSp        = Info.SprintFunc()

	ExtensiveData   = color.New(color.FgBlue, color.Bold)
	ExtensiveDataPL = ExtensiveData.PrintlnFunc()
)
