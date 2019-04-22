package streammux

import core "github.com/libp2p/go-libp2p-core/mux"

// Deprecated: Use github.com/libp2p/go-libp2p-core/mux.ErrReset instead.
var ErrReset = core.ErrReset

// Deprecated: Use github.com/libp2p/go-libp2p-core/mux.MuxedStream instead.
type Stream = core.MuxedStream

// Deprecated: Use github.com/libp2p/go-libp2p-core/mux.NoOpHandler instead.
var NoOpHandler = core.NoOpHandler

// Deprecated: Use github.com/libp2p/go-libp2p-core/mux.MuxedConn instead.
type Conn = core.MuxedConn

// Deprecated: Use github.com/libp2p/go-libp2p-core/mux.Multiplexer instead.
type Transport = core.Multiplexer
