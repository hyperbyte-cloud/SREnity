package entities

type Configuration struct {
	Inputs  []Datasource `pkl:"inputs"`
	Outputs []Datasource `pkl:"outputs"`
	SLOs    []SLO        `pkl:"slos"`
}

type Datasource struct {
	Name string `pkl:"name"`
	Type string `pkl:"type"`

	// Config is a datasource interface that can be used to store any configuration
	// this is defined by the type of the datasource
	Config any `pkl:"config"`
}

type SLO struct {
	Name        string `pkl:"name"`
	Description string `pkl:"description"`
	Output      string `pkl:"output"`
	SLIs        []SLI  `pkl:"slis"`
}

type SLI struct {
	Name        string  `pkl:"name"`
	Description string  `pkl:"description"`
	Input       string  `pkl:"input"`
	Interval    string  `pkl:"interval"`
	Query       string  `pkl:"query"`
	Goal        float64 `pkl:"goal"`
}
