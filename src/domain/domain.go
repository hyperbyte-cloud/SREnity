package domain

import "srenity/entities"

type Domain struct {
	inputDatasources  map[string]entities.DatasourceRepository
	outputDatasources map[string]entities.DatasourceRepository
	metricChan        chan entities.WriteMetric
}

func NewDomain() *Domain {
	return &Domain{
		inputDatasources:  make(map[string]entities.DatasourceRepository),
		outputDatasources: make(map[string]entities.DatasourceRepository),
		metricChan:        make(chan entities.WriteMetric),
	}
}
