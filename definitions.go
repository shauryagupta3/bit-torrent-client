package main

import "net"

type Peer struct {
	IP   net.IP
	Port uint16
}

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

type Handshake struct {
	Pstr     string
	InfoHash [20]byte
	PeerID   [20]byte
}

type Message struct {
	ID      messageID
	Payload []byte
}