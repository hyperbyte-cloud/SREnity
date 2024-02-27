package datasource

import (
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
	// Assume you have a running InfluxDB instance for testing

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
	// Assume you have data in your InfluxDB instance for testing

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
	// Assume you have a running InfluxDB instance for testing
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
