// SPDX-License-Identifier: MIT-0
// LICENSE: https://spdx.org/licenses/MIT-0.html

package svc

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"crypto/rand"
)

// Serve .
func Serve(p chan int) {
	for {
		err := httpService(p)
		log.Printf("Serve error: %v", err)
	}
}

func httpService(p chan int) error {
	l, err := net.Listen("tcp", ":0")
	defer l.Close()
	if err != nil {
		return err
	}
	uuid := UUID()
	http.HandleFunc("/", 
		func(w http.ResponseWriter, r *http.Request){
			w.Write([]byte("Random data: " + uuid))
		},
	)
	log.Printf("Service on port: %v", l.Addr().(*net.TCPAddr).Port)
	p <- l.Addr().(*net.TCPAddr).Port
	for {
		log.Println("starting http serve")
		err := http.Serve(l, nil)
		log.Printf("%v\n", err)
	}
	// return nil
}

// UUID just outputs a random uuid
func UUID() (string) {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        panic(err)
    }
    return fmt.Sprintf(
        "%x-%x-%x-%x-%x",
        b[0:4], b[4:6], b[6:8], b[8:10], b[10:],
    )
}
