package main

import (
	"encoding/binary"
	"fmt"

	"github.com/jackpal/bencode-go"

	// "github.com/veggiedefender/torrent-client/peers"
	"io"
	"net"

	// "net/http"
	"net/url"
	"strconv"
	// "time"
)

type bencodeInfo struct {
	Pieces      string `bencode: "pieces"`
	PieceLength int    `bencode: "pieces length"`
	Length      int    `bencode: "length"`
	Name        string `bencode: "name"`
}

type bencodeTorrent struct {
	Announce string      `bencode: "announce"`
	Info     bencodeInfo `bencode:"info"`
}

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func Open(r io.Reader) (*bencodeTorrent, error) {
	bto := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)
	if err != nil {
		return nil, err
	}
	// resp, err := http.Get('')
	return &bto, nil
}

// func (bto bencodeTorrent) toTorrentFile() (TorrentFile, error) {
// 	return nil, nil
// }

func (t *TorrentFile) buildTrackerURL(peerID [20]byte, port uint16) (string, error) {
	base, err := url.Parse(t.Announce)
	if err != nil {
		return "", err
	}
	params := url.Values{
		"info_hash":  []string{string(t.InfoHash[:])},
		"peer_id":    []string{string(peerID[:])},
		"port":       []string{strconv.Itoa(int(9999))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(t.Length)},
	}
	base.RawQuery = params.Encode()
	return base.String(), nil
}

type Peer struct {
	IP   net.IP
	Port uint16
}

func Unmarsha1(peersBin []byte) ([]Peer, error) {
	const peerSize = 6
	numPeers := len(peersBin) / peerSize
	if len(peersBin)%peerSize != 0 {
		err := fmt.Errorf("Received malformed peers")
		return nil, err
	}
	peers := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i * peerSize
		peers[i].IP = net.IP(peersBin[offset : offset+4])
		peers[i].Port = binary.BigEndian.Uint16(peersBin[offset+4 : offset+6])
	}
	return peers, nil
}

// conn, err := net.DialTCP("tcp", peer.String(), 3*time.Second)
