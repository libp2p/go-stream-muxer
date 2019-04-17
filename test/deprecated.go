package sm_test

import (
	"net"
	"testing"

	"github.com/libp2p/go-libp2p-core/mux"
	moved "github.com/libp2p/go-libp2p-core/mux/test"
)

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.Options instead.
type Options = moved.Options

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.LogWriter instead.
type LogWriter = moved.LogWriter

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.GoServe instead.
func GoServe(t *testing.T, tr mux.Multiplexer, l net.Listener) (done func()) {
	return moved.GoServe(t, tr, l)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestSimpleWrite instead.
func SubtestSimpleWrite(t *testing.T, tr mux.Multiplexer) {
	moved.SubtestSimpleWrite(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress instead.
func SubtestStress(t *testing.T, opt Options) {
	moved.SubtestStress(t, opt)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStreamOpenStress instead.
func SubtestStreamOpenStress(t *testing.T, tr mux.Multiplexer) {
	moved.SubtestStreamOpenStress(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStreamReset instead.
func SubtestStreamReset(t *testing.T, tr mux.Multiplexer) {
	moved.SubtestStreamReset(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestWriteAfterClose instead.
func SubtestWriteAfterClose(t *testing.T, tr mux.Multiplexer) {
	moved.SubtestWriteAfterClose(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress1Conn1Stream1Msg instead.
func SubtestStress1Conn1Stream1Msg(t *testing.T, tr mux.Multiplexer) {
	moved.SubtestStress1Conn1Stream1Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress1Conn1Stream100Msg instead.
func SubtestStress1Conn1Stream100Msg(t *testing.T, tr mux.Multiplexer) {
	moved.SubtestStress1Conn1Stream100Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress1Conn100Stream100Msg instead.
func SubtestStress1Conn100Stream100Msg(t *testing.T, tr mux.Multiplexer) {
	moved.SubtestStress1Conn100Stream100Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress50Conn10Stream50Msg instead.
func SubtestStress50Conn10Stream50Msg(t *testing.T, tr mux.Multiplexer) {
	moved.SubtestStress50Conn10Stream50Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress1Conn1000Stream10Msg instead.
func SubtestStress1Conn1000Stream10Msg(t *testing.T, tr mux.Multiplexer) {
	moved.SubtestStress1Conn1000Stream10Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress1Conn100Stream100Msg10MB instead.
func SubtestStress1Conn100Stream100Msg10MB(t *testing.T, tr mux.Multiplexer) {
	moved.SubtestStress1Conn100Stream100Msg10MB(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.Subtests instead.
var Subtests = moved.Subtests

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestAll instead.
func SubtestAll(t *testing.T, tr mux.Multiplexer) {
	moved.SubtestAll(t, tr)

}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.TransportTest instead.
type TransportTest = moved.TransportTest
