package main

import (
	linters "github.com/pentops/golint-modfile-replace"
	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	// The configuration type will be map[string]any or []interface, it depends on your configuration.
	// You can use https://github.com/mitchellh/mapstructure to convert map to struct.

	return []*analysis.Analyzer{linters.GoModReplaceAnalyzer}, nil
}
