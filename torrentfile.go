package main

import (
	"net/url"
	"strconv"
)

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

// func (bto *bencodeTorrent) toTorrentFile() (TorrentFile, error) {
// 	infoHash, err := bto.Info.hash()
// 	if err != nil {
// 		return TorrentFile{}, err
// 	}
// 	pieceHashes, err := bto.Info.splitPieceHashes()
// 	if err != nil {
// 		return TorrentFile{}, err
// 	}
// 	t := TorrentFile{
// 		Announce:    bto.Announce,
// 		InfoHash:    infoHash,
// 		PieceHashes: pieceHashes,
// 		PieceLength: bto.Info.PieceLength,
// 		Length:      bto.Info.Length,
// 		Name:        bto.Info.Name,
// 	}
// 	return t, nil
// }

func (t *TorrentFile) buildTrackerURL(peerID [20]byte, port uint16) (string, error) {
	base, err := url.Parse(t.Announce)

	if err != nil {
		return "", err
	}

	params := url.Values{
		"info_hash":  []string{string(t.InfoHash[:])},
		"peer_id":    []string{string(peerID[:])},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(t.Length)},
	}

	base.RawQuery = params.Encode()
	return "", nil
}
