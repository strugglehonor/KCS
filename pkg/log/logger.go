package log

import (
	"io"
	"log"
	"os"
)

// Level type
type level int

const (
	// DEBUG level
	Ldebug level = iota
	// INFO level
	Linfo
	// WARNING level
	Lwarning
	// ERROR level
	Lerror
	// FATAL level
	Lfatal

	flag = log.Ldate | log.Ltime
)

// Log level prefix map
var prefix = map[level]string{
	Ldebug:   "DEBUG: ",
	Linfo:    "INFO: ",
	Lwarning: "WARNING: ",
	Lerror:   "ERROR: ",
	Lfatal:   "FATAL: ",
}

// Logger ...
type Logger map[level]LoggerInterface

// New returns instance of Logger
func NewLogger(out, errOut io.Writer, f Formatter) Logger {
	// Fall back to stdout if out not set
	if out == nil {
		out = os.Stdout
	}

	// Fall back to stderr if errOut not set
	if errOut == nil {
		errOut = os.Stderr
	}

	// Fall back to DefaultFormatter if f not set
	if f == nil {
		f = new(DefaultFormatter)
	}

	l := make(map[level]LoggerInterface, 5)
	l[Ldebug] = &Wrapper{lvl: Ldebug, formatter: f, logger: log.New(out, f.GetPrefix(Ldebug)+prefix[Ldebug], flag)}
	l[Linfo] = &Wrapper{lvl: Linfo, formatter: f, logger: log.New(out, f.GetPrefix(Linfo)+prefix[Linfo], flag)}
	l[Lwarning] = &Wrapper{lvl: Linfo, formatter: f, logger: log.New(out, f.GetPrefix(Lwarning)+prefix[Lwarning], flag)}
	l[Lerror] = &Wrapper{lvl: Linfo, formatter: f, logger: log.New(errOut, f.GetPrefix(Lerror)+prefix[Lerror], flag)}
	l[Lfatal] = &Wrapper{lvl: Linfo, formatter: f, logger: log.New(errOut, f.GetPrefix(Lfatal)+prefix[Lfatal], flag)}

	return Logger(l)
}

// Wrapper ...
type Wrapper struct {
	lvl       level
	formatter Formatter
	logger    LoggerInterface
}

// Print ...
func (w *Wrapper) Print(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Print(v...)
}

// Printf ...
func (w *Wrapper) Printf(format string, v ...interface{}) {
	suffix := w.formatter.GetSuffix(w.lvl)
	v = w.formatter.Format(w.lvl, v...)
	w.logger.Printf("%s"+format+suffix, v...)
}

// Println ...
func (w *Wrapper) Println(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Println(v...)
}

// Fatal ...
func (w *Wrapper) Fatal(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Fatal(v...)
}

// Fatalf ...
func (w *Wrapper) Fatalf(format string, v ...interface{}) {
	suffix := w.formatter.GetSuffix(w.lvl)
	v = w.formatter.Format(w.lvl, v...)
	w.logger.Fatalf("%s"+format+suffix, v...)
}

// Fatalln ...
func (w *Wrapper) Fatalln(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Fatalln(v...)
}

// Panic ...
func (w *Wrapper) Panic(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Fatal(v...)
}

// Panicf ...
func (w *Wrapper) Panicf(format string, v ...interface{}) {
	suffix := w.formatter.GetSuffix(w.lvl)
	v = w.formatter.Format(w.lvl, v...)
	w.logger.Panicf("%s"+format+suffix, v...)
}

// Panicln ...
func (w *Wrapper) Panicln(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Panicln(v...)
}
