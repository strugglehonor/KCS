package log

var (
	logger = NewLogger(nil, nil, new(ColouredFormatter))

	// DEBUG ...
	DEBUG = logger[Ldebug]
	// INFO ...
	INFO = logger[Linfo]
	// WARNING ...
	WARNING = logger[Lwarning]
	// ERROR ...
	ERROR = logger[Lerror]
	// FATAL ...
	FATAL = logger[Lfatal]
)

// Set sets a custom logger for all log levels
func Set(l LoggerInterface) {
	DEBUG = l
	INFO = l
	WARNING = l
	ERROR = l
	FATAL = l
}

// SetDebug sets a custom logger for DEBUG level logs
func SetDebug(l LoggerInterface) {
	DEBUG = l
}

// SetInfo sets a custom logger for INFO level logs
func SetInfo(l LoggerInterface) {
	INFO = l
}

// SetWarning sets a custom logger for WARNING level logs
func SetWarning(l LoggerInterface) {
	WARNING = l
}

// SetError sets a custom logger for ERROR level logs
func SetError(l LoggerInterface) {
	ERROR = l
}

// SetFatal sets a custom logger for FATAL level logs
func SetFatal(l LoggerInterface) {
	FATAL = l
}
