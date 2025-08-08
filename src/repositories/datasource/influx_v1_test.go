package datasource

import (
	"net"
	"testing"
	"time"

	"srenity/entities"

	"github.com/stretchr/testify/assert"
)

const (
	testInfluxV1Host        = "http://localhost:8086"
	testInfluxV1Database    = "testdb"
	testInfluxV1Measurement = "test_measurement"
)

func TestInfluxV1Repository_Connect(t *testing.T) {
	// Skip if no local InfluxDB is running
	if !isPortOpen("localhost:8086", 300*time.Millisecond) {
		t.Skip("InfluxDB is not running on localhost:8086")
	}

	config := entities.InfluxV1Configuration{
		Host:     testInfluxV1Host,
		Database: testInfluxV1Database,
	}

	repo := NewInfluxV1Repository(config)

	err := repo.Connect()
	assert.NoError(t, err)
	assert.NotNil(t, repo.client)

	err = repo.Disconnect()
	assert.NoError(t, err)
}

func TestInfluxV1Repository_Query(t *testing.T) {
	// Skip if no local InfluxDB is running
	if !isPortOpen("localhost:8086", 300*time.Millisecond) {
		t.Skip("InfluxDB is not running on localhost:8086")
	}

	config := entities.InfluxV1Configuration{
		Host:     testInfluxV1Host,
		Database: testInfluxV1Database,
	}

	repo := NewInfluxV1Repository(config)

	// Connect to the InfluxDB instance
	err := repo.Connect()
	assert.NoError(t, err)

	query := "SELECT field1 FROM test_measurement"
	result, err := repo.Query(query)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInfluxV1Repository_Write(t *testing.T) {
	// Skip if no local InfluxDB is running
	if !isPortOpen("localhost:8086", 300*time.Millisecond) {
		t.Skip("InfluxDB is not running on localhost:8086")
	}
	config := entities.InfluxV1Configuration{
		Host:     testInfluxV1Host,
		Database: testInfluxV1Database,
	}

	repo := NewInfluxV1Repository(config)

	// Connect to the InfluxDB instance
	err := repo.Connect()
	assert.NoError(t, err)

	// With this repo lets add 5 minutes of data to the InfluxDB instance
	// Add data to the InfluxDB instance
	currentTime := time.Now()

	// Set current time to

	for i := 0; i < 5; i++ {
		data := entities.WriteMetric{
			Name:      testInfluxV1Measurement,
			Tags:      map[string]string{"tag1": "value1"},
			Values:    map[string]interface{}{"field1": 42.0, "field2": 100.0},
			Timestamp: currentTime.Add(time.Duration(i) * time.Minute),
		}
		err := repo.Write(data)
		assert.NoError(t, err)
	}
}

// isPortOpen checks if a TCP address is reachable in a short timeout
func isPortOpen(address string, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}
