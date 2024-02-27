package domain

import (
	"errors"
	"log/slog"
	"srenity/entities"
	"srenity/entities/pkl"
	"srenity/repositories/datasource"

	"github.com/mitchellh/mapstructure"
)

// LoadInputs is to load the input connections
func (d *Domain) LoadInputs(inputs []*pkl.Datasource) error {
	// Open input connections
	for _, input := range inputs {
		// Open the input connection
		datasouce, err := d.OpenDatasource(input)
		if err != nil {
			return err
		}

		// Add the input datasource to the map
		d.inputDatasources[input.Name] = datasouce
	}

	return nil
}

// LoadOutputs is to load the output connections
func (d *Domain) LoadOutputs(outputs []*pkl.Datasource) error {
	// Open output connections
	for _, output := range outputs {
		// Open the output connection
		datasouce, err := d.OpenDatasource(output)
		if err != nil {
			return err
		}

		// Add the output datasource to the map
		d.outputDatasources[output.Name] = datasouce
	}

	return nil
}

// OpenInput is to open the input connection
func (d *Domain) OpenDatasource(input *pkl.Datasource) (entities.DatasourceRepository, error) {
	slog.Info("Opening " + input.Name)

	var datasourceRepo entities.DatasourceRepository

	// Open a new connection from repository
	switch input.Type {
	case "influxdb_v1":

		var influxV1Configuration entities.InfluxV1Configuration
		err := mapstructure.Decode(input.Config.Properties, &influxV1Configuration)
		if err != nil {
			return nil, err
		}

		// Open the connection
		datasourceRepo = datasource.NewInfluxV1Repository(
			influxV1Configuration,
		)
	default:
		return nil, errors.New("unknown datasource type: " + input.Type)
	}

	// Connect to the database
	err := datasourceRepo.Connect()
	if err != nil {
		return nil, err
	}

	return datasourceRepo, nil
}

func (d *Domain) StartOutputWriter(slo *pkl.SLO) {
	// Start the writer
	for metric := range d.metricChan {
		// Add our tags to the metric
		metric.Tags["slo_name"] = slo.Name

		// Write the metric
		err := d.outputDatasources[slo.Output].Write(metric)
		if err != nil {
			slog.Error(err.Error())
		}
	}
}
