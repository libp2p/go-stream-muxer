package sm_test

import (
	"net"
	"testing"

	"github.com/libp2p/go-libp2p-core/mux"
	core "github.com/libp2p/go-libp2p-core/mux/test"
)

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.Options instead.
type Options = core.Options

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.LogWriter instead.
type LogWriter = core.LogWriter

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.GoServe instead.
func GoServe(t *testing.T, tr mux.Multiplexer, l net.Listener) (done func()) {
	return core.GoServe(t, tr, l)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestSimpleWrite instead.
func SubtestSimpleWrite(t *testing.T, tr mux.Multiplexer) {
	core.SubtestSimpleWrite(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress instead.
func SubtestStress(t *testing.T, opt Options) {
	core.SubtestStress(t, opt)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStreamOpenStress instead.
func SubtestStreamOpenStress(t *testing.T, tr mux.Multiplexer) {
	core.SubtestStreamOpenStress(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStreamReset instead.
func SubtestStreamReset(t *testing.T, tr mux.Multiplexer) {
	core.SubtestStreamReset(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestWriteAfterClose instead.
func SubtestWriteAfterClose(t *testing.T, tr mux.Multiplexer) {
	core.SubtestWriteAfterClose(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress1Conn1Stream1Msg instead.
func SubtestStress1Conn1Stream1Msg(t *testing.T, tr mux.Multiplexer) {
	core.SubtestStress1Conn1Stream1Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress1Conn1Stream100Msg instead.
func SubtestStress1Conn1Stream100Msg(t *testing.T, tr mux.Multiplexer) {
	core.SubtestStress1Conn1Stream100Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress1Conn100Stream100Msg instead.
func SubtestStress1Conn100Stream100Msg(t *testing.T, tr mux.Multiplexer) {
	core.SubtestStress1Conn100Stream100Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress50Conn10Stream50Msg instead.
func SubtestStress50Conn10Stream50Msg(t *testing.T, tr mux.Multiplexer) {
	core.SubtestStress50Conn10Stream50Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress1Conn1000Stream10Msg instead.
func SubtestStress1Conn1000Stream10Msg(t *testing.T, tr mux.Multiplexer) {
	core.SubtestStress1Conn1000Stream10Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestStress1Conn100Stream100Msg10MB instead.
func SubtestStress1Conn100Stream100Msg10MB(t *testing.T, tr mux.Multiplexer) {
	core.SubtestStress1Conn100Stream100Msg10MB(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.Subtests instead.
var Subtests = core.Subtests

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.SubtestAll instead.
func SubtestAll(t *testing.T, tr mux.Multiplexer) {
	core.SubtestAll(t, tr)

}

// Deprecated: use github.com/libp2p/go-libp2p-core/mux/test.TransportTest instead.
type TransportTest = core.TransportTest
