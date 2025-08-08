package domain_test

import (
	"srenity/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This is a test function to make sure that datasources are loaded correctly
func TestDomainLoadInputs(t *testing.T) {
	t.Parallel()

	// define the config file
	config := NewTestConfig(t)

	// Skip if no local InfluxDB is running
	if !isInfluxAvailable() {
		t.Skip("InfluxDB is not running on localhost:8086")
	}

	// create a new domain
	d := domain.NewDomain()

	// run the test
	err := d.LoadInputs(config.Inputs)
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}
}

// This is a test function to make sure that datasources are loaded correctly
func TestDomainLoadOutputs(t *testing.T) {
	t.Parallel()

	// create a new domain
	d := domain.NewDomain()

	// Skip if no local InfluxDB is running
	if !isInfluxAvailable() {
		t.Skip("InfluxDB is not running on localhost:8086")
	}

	// run the test
	config := NewTestConfig(t)
	config.Outputs[0].Config.Properties["database"] = "testdb"
	err := d.LoadOutputs(config.Outputs)
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}
}

// This is a test function to make sure that datasources are loaded correctly
// we are not testing the output of Connect in this fucntion as that is up to the repository
func TestDomainOpenDatasource(t *testing.T) {
	t.Parallel()

	// create a new domain
	d := domain.NewDomain()

	// Skip if no local InfluxDB is running
	if !isInfluxAvailable() {
		t.Skip("InfluxDB is not running on localhost:8086")
	}

	// run the test
	config := NewTestConfig(t)
	config.Inputs[0].Config.Properties["database"] = "testdb"
	_, err := d.OpenDatasource(config.Inputs[0])
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}

	// check error output with bad database
	badOutput := NewTestConfig(t)
	badOutput.Outputs[0].Config.Properties["database"] = "i_do_not_exist"
	_, err = d.OpenDatasource(badOutput.Outputs[0])
	if err == nil {
		t.Errorf("Test failed: expected error, got nil")
	}
	assert.Equal(t, "database does not exist: i_do_not_exist", err.Error())

	// check error output with bad type
	badOutput = NewTestConfig(t)
	badOutput.Outputs[0].Type = "i_do_not_exist"
	_, err = d.OpenDatasource(badOutput.Outputs[0])
	if err == nil {
		t.Errorf("Test failed: expected error, got nil")
	}
	assert.Equal(t, "unknown datasource type: i_do_not_exist", err.Error())
}
