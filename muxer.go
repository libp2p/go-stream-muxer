package streammux

import moved "github.com/libp2p/go-libp2p/skel/mux"

// Deprecated: Use github.com/libp2p/go-libp2p/skel/mux.ErrReset instead.
var ErrReset = moved.ErrReset

// Deprecated: Use github.com/libp2p/go-libp2p/skel/mux.MuxStream instead.
type Stream = moved.MuxStream

// Deprecated: Use github.com/libp2p/go-libp2p/skel/mux.NoOpHandler instead.
var NoOpHandler = moved.NoOpHandler

// Deprecated: Use github.com/libp2p/go-libp2p/skel/mux.MuxedConn instead.
type Conn = moved.MuxedConn

// Deprecated: Use github.com/libp2p/go-libp2p/skel/mux.Multiplexer instead.
type Transport = moved.Multiplexer
