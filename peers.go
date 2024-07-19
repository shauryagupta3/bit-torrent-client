package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

func Unmarshal(peersbin []byte) ([]Peer, error) {
	const peerSize = 6 // 4 + 2
	numPeers := len(peersbin) / peerSize
	if len(peersbin)%peerSize != 0 {
		err := fmt.Errorf("Recieved malformed peers")
		return nil, err
	}
	peers := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i * peerSize
		peers[i].IP = net.IP(peersbin[offset : offset+4])
		peers[i].Port = binary.BigEndian.Uint16(peersbin[offset+4 : offset+6])
	}
	return peers, nil
}
