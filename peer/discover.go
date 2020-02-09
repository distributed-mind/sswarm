// SPDX-License-Identifier: MIT-0
// LICENSE: https://spdx.org/licenses/MIT-0.html

package peer

import (
	"log"
	// "net"
	"time"
)

const (
	maxUDPSize int = 65535
)

// Discover .
func Discover(port chan int, info chan Info, id string) { 
	// open udp listener
	// send multicast public key, and service port
	// if public key received is mine (log), ignore
	// if its a new peer (log), connect to service port
	// if connection successful (log), put peer info into bucket
	// log contents of bucket to stdout

	go func() {
		listen(info)
	}()
	p := <- port
	for {
		err := multicast(p, id)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(5 * time.Second)
	}
}
