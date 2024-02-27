package domain_test

import (
	"srenity/domain"
	"srenity/entities/pkl"
	"testing"

	apkl "github.com/apple/pkl-go/pkl"
	"github.com/stretchr/testify/assert"
)

func TestDomainValidate(t *testing.T) {
	t.Parallel()

	// this test is checking to see if the output of the validate config is correct

	// define the config file which is correct
	correctConfig := NewTestConfig(t)

	// create a new domain
	d := domain.NewDomain()

	// run the test
	err := d.Validate(correctConfig)
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}

	// using the above config, we will now test the validate function with incorrect config
	// define the config file which is incorrect
	missingOutput := NewTestConfig(t)
	missingOutput.Outputs = []*pkl.Datasource{}

	// run the test
	err = d.Validate(missingOutput)
	assert.Equal(t, "no outputs found", err.Error())

	missingInputs := NewTestConfig(t)
	missingInputs.Inputs = []*pkl.Datasource{}

	// run the test
	err = d.Validate(missingInputs)
	assert.Equal(t, "no inputs found", err.Error())

	missingSLOs := NewTestConfig(t)
	missingSLOs.SLOs = []*pkl.SLO{}

	// run the test
	err = d.Validate(missingSLOs)
	assert.Equal(t, "no SLO's found", err.Error())

	missingSLIs := NewTestConfig(t)
	missingSLIs.SLOs[0].SLIs = []*pkl.SLI{}

	// run the test
	err = d.Validate(missingSLIs)
	assert.Equal(t, "No SLI's found for test SLO", err.Error())

	missingSLIInputMapping := NewTestConfig(t)
	missingSLIInputMapping.SLOs[0].SLIs[0].Input = "not_a_input"

	// run the test
	err = d.Validate(missingSLIInputMapping)
	assert.Equal(t, "Input not_a_input not found for test SLI", err.Error())

	missingSLOOutputMapping := NewTestConfig(t)
	missingSLOOutputMapping.SLOs[0].Output = ""

	// run the test
	err = d.Validate(missingSLOOutputMapping)
	assert.Equal(t, "No output found for test SLO", err.Error())

	incorrectSLOOutputMapping := NewTestConfig(t)
	incorrectSLOOutputMapping.SLOs[0].Output = "not_a_output"

	// run the test
	err = d.Validate(incorrectSLOOutputMapping)
	assert.Equal(t, "Output not_a_output not found for test SLO", err.Error())

	incorrectSLIInputMapping := NewTestConfig(t)
	incorrectSLIInputMapping.SLOs[0].SLIs[0].Input = "test1"

	// run the test
	err = d.Validate(incorrectSLIInputMapping)
	assert.Equal(t, "Input test1 not found for test SLI", err.Error())

	// test the SLI interval
	missingSLIInterval := NewTestConfig(t)
	missingSLIInterval.SLOs[0].SLIs[0].Interval = nil

	// run the test
	err = d.Validate(missingSLIInterval)
	assert.Equal(t, "No interval found for test SLI", err.Error())

	// test the function with a nil config
	err = d.Validate(nil)
	assert.Equal(t, "config is nil", err.Error())

}

func NewTestConfig(t *testing.T) *pkl.Configuration {
	return &pkl.Configuration{
		Inputs: []*pkl.Datasource{
			{
				Name: "test",
				Type: "influxdb_v1",
				Config: &apkl.Object{
					Properties: map[string]any{
						"host":     "http://localhost:8086",
						"database": "testdb",
					},
				},
			},
		},
		Outputs: []*pkl.Datasource{
			{
				Name: "test",
				Type: "influxdb_v1",
				Config: &apkl.Object{
					Properties: map[string]any{
						"host":     "http://localhost:8086",
						"database": "testdb",
					},
				},
			},
		},
		SLOs: []*pkl.SLO{
			{
				Name:   "test SLO",
				Output: "test",
				SLIs: []*pkl.SLI{
					{
						Name:  "test SLI",
						Query: "SELECT mean(field1) FROM test_measurement",
						Goal:  99.9,
						Input: "test",
						Interval: &apkl.Duration{
							Value: 60,
							Unit:  apkl.Second,
						},
					},
				},
			},
		},
	}
}
