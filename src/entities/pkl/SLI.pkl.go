// Code generated from Pkl module `srenity.entities.Configuration`. DO NOT EDIT.
package pkl

import "github.com/apple/pkl-go/pkl"

type SLI struct {
	Name string `pkl:"name"`

	Description string `pkl:"description"`

	Input string `pkl:"input"`

	Interval *pkl.Duration `pkl:"interval"`

	Query string `pkl:"query"`

	Goal float64 `pkl:"goal"`
}
