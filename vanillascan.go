package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	// Define command line flags
	host := flag.String("host", "localhost", "Host to connect to")

	// Parse command line arguments
	flag.Parse()

	// Access the flag values
	fmt.Printf("Scanning host: %s\n", *host)

	var wg sync.WaitGroup

	// Iterate over port range and attempt connection
	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", *host, p)
			conn, err := net.DialTimeout("tcp", address, 10000*time.Millisecond)
			if err != nil {
				// Port is closed
				return
			}
			defer conn.Close()
			fmt.Printf("Port %d is open\n", p)
		}(port)
	}

	wg.Wait()
}
