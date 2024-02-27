package entities

import "time"

// Metric is a struct that represents a metric expected to be returned by a datasource
type WriteMetric struct {
	Timestamp time.Time
	Name      string
	Tags      map[string]string
	Values    map[string]interface{}
}

type QueryMetric struct {
	Values []float64
}
