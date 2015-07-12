package multistream

import (
	"testing"

	test "github.com/jbenet/go-stream-muxer/test"
)

func TestMultiStreamTransport(t *testing.T) {
	test.SubtestAll(t, NewTransport())
}
