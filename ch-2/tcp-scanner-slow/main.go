package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered.
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
	elapsed := time.Since(start)
	fmt.Printf("The script took %s to run.\n", elapsed)
}
