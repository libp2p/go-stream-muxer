package peerstream_multiplex

import (
	"testing"

	test "github.com/jbenet/go-stream-mux/test"
)

func TestMultiplexTransport(t *testing.T) {
	test.SubtestAll(t, DefaultTransport)
}
