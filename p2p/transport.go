package p2p

// Peer represents the remote node
type Peer interface {
	Close() error
}

// Transport handles communication between nodes in the network.
// This can be in the form (TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}