package main

import (
	linters "github.com/pentops/golint-modfile-replace"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(linters.GoModReplaceAnalyzer)
}
