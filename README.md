# go-stream-muxer - generalized stream multiplexing


go-stream-muxer is a common interface for stream muxers, with common tests. It wraps other stream muxers (like [muxado](https://github.com/inconshreveable/muxado), [spdystream](https://github.com/docker/spdystream) and [yamux](https://github.com/hashicorp/yamux)).

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://ipn.io) [![](https://img.shields.io/badge/freenode-%23ipfs-blue.svg?style=flat-square)](http://webchat.freenode.net/?channels=%23ipfs)

> A test suite and interface you can use to implement a stream muxer.

### Godoc: https://godoc.org/github.com/jbenet/go-stream-mux

## Implementations

* [yamux](yamux)
* [muxado](muxado)
* [multiplex](multiplex)
* [spdystream](spdystream)

## Badge

Include this badge in your readme if you make a new module that uses abstract-stream-muxer API.

![](/img/badge.png)
