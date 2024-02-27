package entities

// SLIResults is a struct that contains the results of the SLI calculation
type SLIResult struct {
	SLOName     string
	Name        string
	Status      string
	Goal        float64
	ErrorBudget string
	Max         float64
	Mean        float64
	Min         float64
	Data        []float64
	Error       string
}
