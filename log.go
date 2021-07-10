package discover

type Logger interface {
	// Log a message at the given level with context key/value pairs
	Trace(msg string, ctx ...interface{})
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Warn(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
	Crit(msg string, ctx ...interface{})
}

type LazyLogValue func() interface{}

type nilLogger struct{}

func (n nilLogger) Trace(msg string, ctx ...interface{}) {}

func (n nilLogger) Debug(msg string, ctx ...interface{}) {}

func (n nilLogger) Info(msg string, ctx ...interface{}) {}

func (n nilLogger) Warn(msg string, ctx ...interface{}) {}

func (n nilLogger) Error(msg string, ctx ...interface{}) {}

func (n nilLogger) Crit(msg string, ctx ...interface{}) {}
