package print

import (
	"os"

	"github.com/albert-widi/sqlimporter/print/printer"
	"github.com/fatih/color"
)

var (
	debugPrinter = printer.New("[debug]", color.FgHiYellow)
	infoPrinter  = printer.New("[info]", color.FgHiGreen)
	warnPrinter  = printer.New("[warn]", color.FgHiCyan)
	errorPrinter = printer.New("[error]", color.FgHiRed)
	fatalPrinter = printer.New("[fatal]", color.FgHiRed)

	// for debug flag
	isVerbose     bool
	isVeryVerbose bool
)

// SetVerbose to flag allowed to print debug or not
func SetVerbose(v bool) {
	isVerbose = v
}

// SetVeryVerbose is useful for very verbose info
func SetVeryVerbose(vv bool) {
	isVeryVerbose = vv
}

// Debug print
func Debug(v ...interface{}) {
	// idiomatic debug
	if !isVerbose && !isVeryVerbose {
		return
	}
	debugPrinter.Print(v...)
}

// Debugv print for very verbose
func Debugv(v ...interface{}) {
	// idiomatic debug
	if !isVeryVerbose {
		return
	}
	debugPrinter.Print(v...)
}

// Debugp print with prefix
func Debugp(prefix string, v ...interface{}) {
	// idiomatic debug
	if !isVerbose && !isVeryVerbose {
		return
	}
	debugPrinter.WithPrefix(prefix).Print(v...)
}

// Debugpv print with prefix for very verbose
func Debugpv(prefix string, v ...interface{}) {
	if !isVeryVerbose {
		return
	}
	debugPrinter.WithPrefix(prefix).Print(v...)
}

// Info print
func Info(v ...interface{}) {
	infoPrinter.Print(v...)
}

// Infop print with prefix
func Infop(prefix string, v ...interface{}) {
	infoPrinter.WithPrefix(prefix).Print(v...)
}

// Warn print
func Warn(v ...interface{}) {
	warnPrinter.Print(v...)
}

// Warnp print with prefix
func Warnp(prefix string, v ...interface{}) {
	warnPrinter.WithPrefix(prefix).Print(v...)
}

// Error print
func Error(v ...interface{}) {
	errorPrinter.Print(v...)
}

// Errorp print with prefix
func Errorp(prefix string, v ...interface{}) {
	errorPrinter.WithPrefix(prefix).Print(v...)
}

// Fatal print
func Fatal(v ...interface{}) {
	Error(v...)
	os.Exit(1)
}

// Fatalp print with prefix
func Fatalp(prefix string, v ...interface{}) {
	Errorp(prefix, v...)
	os.Exit(1)
}

// FatalError print
func FatalError(err error) {
	if err == nil {
		return
	}
	Fatal(err)
}

// FatalErrorp print with prefix
func FatalErrorp(prefix string, err error) {
	if err == nil {
		return
	}
	Fatal(prefix, err)
}
