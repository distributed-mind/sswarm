// SPDX-License-Identifier: MIT-0
// LICENSE: https://spdx.org/licenses/MIT-0.html

package cli

import (
	// "log"
	"encoding/base64"
	"sswarm/peer"
	"sswarm/svc"
)

func run() {
	port := make(chan int)
	go svc.Serve(port)
	peerInfo := make(chan peer.Info)
	myID := base64.StdEncoding.EncodeToString(peer.ID.PublicKey[:])
	go peer.Discover(port, peerInfo, "@"+myID+".ed25519")
	go peer.Log(peerInfo)
	select {} // cheap hang
}
