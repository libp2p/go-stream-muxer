package streammux

import moved "github.com/libp2p/go-libp2p-core/mux"

// Deprecated: Use github.com/libp2p/go-libp2p-core/mux.ErrReset instead.
var ErrReset = moved.ErrReset

// Deprecated: Use github.com/libp2p/go-libp2p-core/mux.MuxedStream instead.
type Stream = moved.MuxedStream

// Deprecated: Use github.com/libp2p/go-libp2p-core/mux.NoOpHandler instead.
var NoOpHandler = moved.NoOpHandler

// Deprecated: Use github.com/libp2p/go-libp2p-core/mux.MuxedConn instead.
type Conn = moved.MuxedConn

// Deprecated: Use github.com/libp2p/go-libp2p-core/mux.Multiplexer instead.
type Transport = moved.Multiplexer
