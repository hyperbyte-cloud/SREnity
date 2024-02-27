package entities

// Database Configuration
type InfluxV1Configuration struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type DatasourceRepository interface {
	// Connect to the database
	Connect() error

	// Disconnect from the database
	Disconnect() error

	// Query the database
	Query(string) (QueryMetric, error)

	// Write to the database
	Write(WriteMetric) error
}
