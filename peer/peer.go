// SPDX-License-Identifier: MIT-0
// LICENSE: https://spdx.org/licenses/MIT-0.html

package peer

import (
	"log"
	"net"
	"crypto/rand"
	"crypto/ed25519"
)

// Info .
type Info struct {
	IP net.IP
	Port int
	ID [32]byte
}

var (
	// ID .
	ID struct {
		PublicKey [32]byte
		PrivateKey [64]byte
	}
	seenPeers = make(map[string]Info)
)

func init() {
	ed25519PublicKey, ed25519PrivateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	copy(ID.PublicKey[:], ed25519PublicKey)
	copy(ID.PrivateKey[:], ed25519PrivateKey)
}