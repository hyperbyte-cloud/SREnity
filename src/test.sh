#!/bin/bash

# Run tests
go test $(go list ./... | grep -v /vendor/)