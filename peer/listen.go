// SPDX-License-Identifier: MIT-0
// LICENSE: https://spdx.org/licenses/MIT-0.html

package peer

import (
	"log"
	"net"
	"strconv"
	"strings"
	"encoding/base64"
)

func listen(peerInfo chan Info) {
	conn, err := net.ListenMulticastUDP("udp", nil, 
		&net.UDPAddr{
			IP: []byte{239, 255, 255, 19},
			Port: 25519,
			// Zone: "udp4",
		},
	)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	err = conn.SetReadBuffer(maxUDPSize)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		b := make([]byte, maxUDPSize)
		l, address, err := conn.ReadFromUDP(b)
		if err != nil {
			log.Println(err)
			continue
		} else {
			s := strings.Split(string(b[:l]), "@")
			if len(s) == 2 {
				port, err := strconv.Atoi(s[0])
				if err != nil {
					log.Println(err)
					continue
				}
				k, err := base64.StdEncoding.DecodeString(
					strings.Split(s[1], ".")[0],
				)
				if err != nil {
					log.Println(err)
					continue
				}
				id := [32]byte{}
				copy(id[:], k)
				go func() {
					peerInfo <- Info{
						IP: address.IP,
						Port: port,
						ID: id,
					}
				}()
			}
		}
	}
}
