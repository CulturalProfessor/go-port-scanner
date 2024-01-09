package port

import (
	"net"
	"strconv"
	"sync"
	"time"
)

type ScanResult struct {
	Port  int
	State string
}

func ScanPort(protocol, hostname string, port int, wg *sync.WaitGroup, mu *sync.Mutex, results *[]ScanResult) {
	defer wg.Done()

	result := ScanResult{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 5*time.Second)

	mu.Lock()
	defer mu.Unlock()

	if err != nil {
		result.State = "Closed"
		*results = append(*results, result)
		return
	}
	defer conn.Close()

	result.State = "Open"
	*results = append(*results, result)
}

func InitialScan(hostname string, startPort, endPort int) []ScanResult {
	var results []ScanResult
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := startPort; i <= endPort; i++ {
		wg.Add(1)
		go ScanPort("tcp", hostname, i, &wg, &mu, &results)
	}

	wg.Wait()
	return results
}
