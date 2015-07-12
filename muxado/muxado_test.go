package peerstream_muxado

import (
	"testing"

	test "github.com/jbenet/go-stream-mux/test"
)

func TestMuxadoTransport(t *testing.T) {
	test.SubtestAll(t, Transport)
}
