package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("The script took %s to run.\n", elapsed)
}
