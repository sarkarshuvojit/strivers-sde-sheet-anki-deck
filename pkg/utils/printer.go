package utils

import (
	"os"

	"github.com/sarkarshuvojit/pprinter"
)

var PPrinter = pprinter.WithTheme(&pprinter.AyuLightTheme)

func ErrAndExit(err error) {
	PPrinter.Error(err.Error())
	os.Exit(1)
}
