package datasource

import (
	"encoding/json"
	"errors"
	"log/slog"
	"srenity/entities"

	"github.com/influxdata/influxdb/client/v2"
)

// InfluxV1Repository is the repository for InfluxDB v1
type InfluxV1Repository struct {
	client client.Client
	config entities.InfluxV1Configuration
}

// NewInfluxV1Repository creates a new InfluxV1Repository
func NewInfluxV1Repository(config entities.InfluxV1Configuration) *InfluxV1Repository {
	return &InfluxV1Repository{
		config: config,
	}
}

// Connect to the database
func (r *InfluxV1Repository) Connect() error {
	// Validate the configuration
	err := r.validateConfig()
	if err != nil {
		return err
	}

	slog.Info("Connecting to InfluxDB v1 on " + r.config.Host)

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     r.config.Host,
		Username: r.config.Username,
		Password: r.config.Password,
	})

	if err != nil {
		return err
	}

	_, _, err = c.Ping(30)
	if err != nil {
		return err
	}

	slog.Info("Connected to InfluxDB v1")
	r.client = c

	// check if the database exists
	err = r.checkDatabaseExists()
	if err != nil {
		return err
	}

	return nil
}

func (r *InfluxV1Repository) validateConfig() error {
	if r.config.Host == "" {
		return errors.New("host is required")
	}

	if r.config.Database == "" {
		return errors.New("database is required")
	}

	return nil
}

// Disconnect from the database
func (r *InfluxV1Repository) Disconnect() error {
	err := r.client.Close()
	if err != nil {
		return err
	}

	r.client = nil
	return nil
}

// Query the database
func (r *InfluxV1Repository) Query(query string) (entities.QueryMetric, error) {
	var queryMetric entities.QueryMetric
	q := client.Query{
		Command:  query,
		Database: r.config.Database,
	}
	response, err := r.client.Query(q)
	if err != nil {
		return queryMetric, err
	}

	// Convert the response to a QueryMetric
	for _, result := range response.Results {
		for _, series := range result.Series {
			for _, value := range series.Values {
				if value[0] == nil || value[1] == nil {
					continue
				}
				metric := value[1].(json.Number)

				// Convert the metric to a float64
				metricFloat, err := metric.Float64()
				if err != nil {
					return queryMetric, err
				}

				// Add the metric to the QueryMetric
				queryMetric.Values = append(queryMetric.Values, metricFloat)
			}
		}
	}

	return queryMetric, nil
}

// Write to the database
func (r *InfluxV1Repository) Write(data entities.WriteMetric) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  r.config.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// Create a new point batch
	point, err := client.NewPoint(
		data.Name,
		data.Tags,
		data.Values,
		data.Timestamp,
	)
	if err != nil {
		return err
	}

	bp.AddPoint(point)

	return r.client.Write(bp)
}

func (r *InfluxV1Repository) checkDatabaseExists() error {
	// check if the database exists
	q := client.Query{
		Command:  "SHOW DATABASES;",
		Database: r.config.Database,
	}

	resp, err := r.client.Query(q)
	if err != nil {
		return err
	}

	for _, dbName := range resp.Results[0].Series[0].Values {
		if dbName[0] == r.config.Database {
			return nil
		}
	}

	return errors.New("database does not exist: " + r.config.Database)
}
