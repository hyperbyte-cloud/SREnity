package domain_test

import (
	"srenity/domain"
	"testing"
)

func TestDomainTest(t *testing.T) {
	t.Parallel()

	// this test is checking to see if the output of the test function is correct

	// define the config file
	config := NewTestConfig(t)

	// create a new domain
	d := domain.NewDomain()

	// Open input and output connections
	err := d.LoadInputs(config.Inputs)
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}

	// run the test
	results, err := d.Test(config)
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}

	// check the output
	if len(results) != 1 {
		t.Errorf("Test failed: expected 1 result, got %d", len(results))
	}

	// check the result
	result := results[0]
	if result.SLOName != "test SLO" {
		t.Errorf("Test failed: expected SLOName to be 'test SLO', got %s", result.SLOName)
	}

	if result.Name != "test SLI" {
		t.Errorf("Test failed: expected Name to be 'test SLI', got %s", result.Name)
	}

	if result.Status != "Passing" {
		t.Errorf("Test failed: expected Status to be 'Failing', got %s", result.Status)
	}

	if result.Goal != 99.9 {
		t.Errorf("Test failed: expected Goal to be 99.9, got %f", result.Goal)
	}

	if result.ErrorBudget != "0.1" {
		t.Errorf("Test failed: expected ErrorBudget to be '0.1', got %s", result.ErrorBudget)
	}

}
