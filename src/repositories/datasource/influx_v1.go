package datasource

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"srenity/entities"
	"strconv"

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
	if r.client == nil {
		// Nothing to do
		return nil
	}
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
	if r.client == nil {
		return queryMetric, errors.New("not connected to InfluxDB")
	}
	q := client.Query{
		Command:  query,
		Database: r.config.Database,
	}
	response, err := r.client.Query(q)
	if err != nil {
		return queryMetric, err
	}
	if response == nil {
		return queryMetric, errors.New("nil response from InfluxDB")
	}
	if response.Error() != nil {
		return queryMetric, response.Error()
	}

	// Convert the response to a QueryMetric
	for _, result := range response.Results {
		for _, series := range result.Series {
			for _, value := range series.Values {
				if len(value) < 2 || value[1] == nil {
					continue
				}

				var metricFloat float64
				switch v := value[1].(type) {
				case json.Number:
					mf, err := v.Float64()
					if err != nil {
						return queryMetric, err
					}
					metricFloat = mf
				case float64:
					metricFloat = v
				case float32:
					metricFloat = float64(v)
				case int64:
					metricFloat = float64(v)
				case int:
					metricFloat = float64(v)
				case string:
					mf, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return queryMetric, fmt.Errorf("unable to parse string metric '%s' to float: %w", v, err)
					}
					metricFloat = mf
				default:
					// unsupported type; skip
					continue
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
	if r.client == nil {
		return errors.New("not connected to InfluxDB")
	}
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

	if r.client == nil {
		return errors.New("not connected to InfluxDB")
	}

	resp, err := r.client.Query(q)
	if err != nil {
		return err
	}
	if resp == nil {
		return errors.New("nil response from InfluxDB")
	}
	if resp.Error() != nil {
		return resp.Error()
	}
	if len(resp.Results) == 0 {
		return errors.New("no results returned from SHOW DATABASES")
	}
	if len(resp.Results[0].Series) == 0 {
		return errors.New("no series returned from SHOW DATABASES")
	}
	for _, dbName := range resp.Results[0].Series[0].Values {
		if len(dbName) == 0 || dbName[0] == nil {
			continue
		}
		// Database name is expected to be in the first column
		switch name := dbName[0].(type) {
		case string:
			if name == r.config.Database {
				return nil
			}
		default:
			if fmt.Sprint(name) == r.config.Database {
				return nil
			}
		}
	}

	return errors.New("database does not exist: " + r.config.Database)
}
