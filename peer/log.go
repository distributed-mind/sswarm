package peer

import (
	//
	"log"
	"encoding/base64"
)

// Log .
func Log(info chan Info) {
	//
	for {
		p := <- info
		if isNewPeer(p) {
			log.Printf(
				"Seeing New Peer:\n----- IP:   %v\n----- Port: %v\n----- ID:   %v\n\n",
				p.IP.String(),
				p.Port,
				"@" + base64.StdEncoding.EncodeToString(p.ID[:]) + ".ed25519",
			)
		}
	}
}

func isNewPeer(p Info) bool {
	new := true
	for id := range seenPeers {
		if string(p.ID[:]) == id {
			new = false
		}
	}
	if new {
		seenPeers[string(p.ID[:])] = p
	}
	return new
}