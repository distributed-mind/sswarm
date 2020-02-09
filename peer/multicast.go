// SPDX-License-Identifier: MIT-0
// LICENSE: https://spdx.org/licenses/MIT-0.html

package peer

import (
	//
	"net"
	"log"
	"time"
	"strconv"
)


func multicast(p int, id string) error {
	conn, err := net.DialUDP("udp", nil,
		&net.UDPAddr{
			IP:   []byte{239, 255, 255, 19},
			Port: 25519,
			// Zone: "udp4",
		},
	)
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()
	err = conn.SetDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		log.Println(err)
		return err
	}
	i, err := conn.Write([]byte(strconv.Itoa(p) + id)) // "45799@Sg8QsBUgjflSjMeVBUhMxVDq7ynTlvwxLzFRUKF4k74=.ed25519"
	if err != nil {
		log.Println(i, err)
		return err
	}
	return nil
}
