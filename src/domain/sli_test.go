package domain_test

import (
	"srenity/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This is a test function of the server to output the SLI Calculation
func TestDomainCalculateSLI(t *testing.T) {
	t.Parallel()

	// define the config file
	config := NewTestConfig(t)

	// create a new domain
	d := domain.NewDomain()

	// Load the inputs
	err := d.LoadInputs(config.Inputs)
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}

	// run the test
	SLIMet, SLIMax, SLIMean, SLIMin, _, _, err := d.CalculateSLI(config.SLOs[0].SLIs[0])
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}

	// check the output
	if SLIMet != true {
		t.Errorf("Test failed: expected SLOMet to be true, got %t", SLIMet)
	}

	if SLIMax != 42.0 {
		t.Errorf("Test failed: expected SLIMax to be 42, got %f", SLIMax)
	}

	if SLIMean != 42.0 {
		t.Errorf("Test failed: expected SLIMean to be 42, got %f", SLIMean)
	}

	if SLIMin != 0 {
		t.Errorf("Test failed: expected SLIMin to be 0, got %f", SLIMin)
	}

}

// This is a test function to test the MonitorSLI function
func TestDomainMonitorSLI(t *testing.T) {
	t.Parallel()

	// define the config file
	config := NewTestConfig(t)

	// create a new domain
	d := domain.NewDomain()

	// Load the inputs
	err := d.LoadInputs(config.Inputs)
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}

	// run the test
	err = d.MonitorSLI(config.SLOs[0].SLIs[0])
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}

	// test the error output
	badSLI := NewTestConfig(t)
	badSLI.SLOs[0].SLIs[0].Input = "not_a_input"
	err = d.MonitorSLI(badSLI.SLOs[0].SLIs[0])
	if err == nil {
		t.Errorf("Test failed: expected error, got nil")
	}
	assert.Equal(t, "input not_a_input not found", err.Error())

}
