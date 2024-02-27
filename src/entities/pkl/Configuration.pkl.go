// Code generated from Pkl module `srenity.entities.Configuration`. DO NOT EDIT.
package pkl

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Configuration struct {
	Inputs []*Datasource `pkl:"inputs"`

	Outputs []*Datasource `pkl:"outputs"`

	SLOs []*SLO `pkl:"slos"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Configuration
func LoadFromPath(ctx context.Context, path string) (ret *Configuration, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Configuration
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Configuration, error) {
	var ret Configuration
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
