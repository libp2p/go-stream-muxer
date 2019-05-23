package sm_test

import (
	"net"
	"testing"

	"github.com/libp2p/go-libp2p-core/mux"
	tmux "github.com/libp2p/go-libp2p-testing/suites/mux"
)

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.Options instead.
type Options = tmux.Options

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.LogWriter instead.
type LogWriter = tmux.LogWriter

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.GoServe instead.
func GoServe(t *testing.T, tr mux.Multiplexer, l net.Listener) (done func()) {
	return tmux.GoServe(t, tr, l)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestSimpleWrite instead.
func SubtestSimpleWrite(t *testing.T, tr mux.Multiplexer) {
	tmux.SubtestSimpleWrite(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestStress instead.
func SubtestStress(t *testing.T, opt Options) {
	tmux.SubtestStress(t, opt)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestStreamOpenStress instead.
func SubtestStreamOpenStress(t *testing.T, tr mux.Multiplexer) {
	tmux.SubtestStreamOpenStress(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestStreamReset instead.
func SubtestStreamReset(t *testing.T, tr mux.Multiplexer) {
	tmux.SubtestStreamReset(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestWriteAfterClose instead.
func SubtestWriteAfterClose(t *testing.T, tr mux.Multiplexer) {
	tmux.SubtestWriteAfterClose(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestStress1Conn1Stream1Msg instead.
func SubtestStress1Conn1Stream1Msg(t *testing.T, tr mux.Multiplexer) {
	tmux.SubtestStress1Conn1Stream1Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestStress1Conn1Stream100Msg instead.
func SubtestStress1Conn1Stream100Msg(t *testing.T, tr mux.Multiplexer) {
	tmux.SubtestStress1Conn1Stream100Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestStress1Conn100Stream100Msg instead.
func SubtestStress1Conn100Stream100Msg(t *testing.T, tr mux.Multiplexer) {
	tmux.SubtestStress1Conn100Stream100Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestStress50Conn10Stream50Msg instead.
func SubtestStress50Conn10Stream50Msg(t *testing.T, tr mux.Multiplexer) {
	tmux.SubtestStress50Conn10Stream50Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestStress1Conn1000Stream10Msg instead.
func SubtestStress1Conn1000Stream10Msg(t *testing.T, tr mux.Multiplexer) {
	tmux.SubtestStress1Conn1000Stream10Msg(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestStress1Conn100Stream100Msg10MB instead.
func SubtestStress1Conn100Stream100Msg10MB(t *testing.T, tr mux.Multiplexer) {
	tmux.SubtestStress1Conn100Stream100Msg10MB(t, tr)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.Subtests instead.
var Subtests = tmux.Subtests

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.SubtestAll instead.
func SubtestAll(t *testing.T, tr mux.Multiplexer) {
	tmux.SubtestAll(t, tr)

}

// Deprecated: use github.com/libp2p/go-libp2p-testing/mux.TransportTest instead.
type TransportTest = tmux.TransportTest
