package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	elapsed := time.Since(start)
	fmt.Printf("The script took %s to run.\n", elapsed)
}
