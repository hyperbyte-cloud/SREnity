package domain_test

import (
	"net"
	"time"
)

// isInfluxAvailable checks if InfluxDB is reachable locally to decide whether to skip integration tests.
func isInfluxAvailable() bool {
	conn, err := net.DialTimeout("tcp", "localhost:8086", 300*time.Millisecond)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}
