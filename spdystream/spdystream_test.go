package peerstream_spdystream

import (
	"testing"

	test "github.com/jbenet/go-stream-mux/test"
)

func TestSpdyStreamTransport(t *testing.T) {
	test.SubtestAll(t, Transport)
}
