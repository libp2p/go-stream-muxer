package multistream

import (
	"testing"

	spdy "github.com/jbenet/go-stream-muxer/spdystream"
	test "github.com/jbenet/go-stream-muxer/test"
)

func TestMultiStreamTransport(t *testing.T) {
	tpt := NewBlankTransport()
	tpt.AddTransport("/spdy", spdy.Transport)
	test.SubtestAll(t, tpt)
}
