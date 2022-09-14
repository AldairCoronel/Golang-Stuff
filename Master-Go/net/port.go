package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "Url to scan")

func main() {
	flag.Parse()
	var wg sync.WaitGroup

	for port := 0; port < 1000; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			// tcp conn
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open\n", port)
		}(port)
	}
	wg.Wait()
}
