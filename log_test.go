package discover

import (
	"fmt"
	"strings"
	"testing"
)

// Minimal logger implementation for testing,
func newTestLogger(t *testing.T) Logger {
	return &testLogger{"", t}
}

type testLogger struct {
	prefix string
	t      *testing.T
}

func (n *testLogger) write(lvl string, msg string, ctx ...interface{}) {
	if len(ctx)%2 != 0 {
		n.t.Errorf("malformatted log msg: %s %s %v", lvl, msg, ctx)
		return
	}
	var fields strings.Builder
	for i := 0; i < len(ctx); i += 2 {
		k, ok := ctx[i].(string)
		if !ok {
			n.t.Errorf("malformatted log field %d: %v %v", i, ctx[i], ctx[i+1])
			continue
		}
		v := ctx[i+1]
		if lazy, ok := v.(LazyLogValue); ok {
			v = lazy()
		}
		fields.WriteString(" ")
		fields.WriteString(k)
		fields.WriteString("=")
		fields.WriteString(fmt.Sprintf("%v", v))
	}
	n.t.Logf("%s [%s]  %s%s", n.prefix, lvl, msg, fields.String())
}

func (n *testLogger) Trace(msg string, ctx ...interface{}) {
	n.write("trace", msg, ctx...)
}

func (n *testLogger) Debug(msg string, ctx ...interface{}) {
	n.write("debug", msg, ctx...)
}

func (n *testLogger) Info(msg string, ctx ...interface{}) {
	n.write("info", msg, ctx...)
}

func (n *testLogger) Warn(msg string, ctx ...interface{}) {
	n.write("warn", msg, ctx...)
}

func (n *testLogger) Error(msg string, ctx ...interface{}) {
	n.write("error", msg, ctx...)
}

func (n *testLogger) Crit(msg string, ctx ...interface{}) {
	n.write("crit", msg, ctx...)
}
