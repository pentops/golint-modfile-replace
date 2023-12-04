package linters

import (
	"fmt"
	"go/token"
	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/analysis"
	"os"
	"path/filepath"
)

type position struct {
	Start token.Pos
	End   token.Pos
}

var GoModReplaceAnalyzer = &analysis.Analyzer{
	Name: "gomodreplace",
	Doc:  "finds replace directives in go.mod files",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	var posByModFilePath = make(map[string]position)
	var modFilePaths []string

	for _, file := range pass.Files {
		goFilePath := pass.Fset.File(file.Package).Name()
		split := filepath.SplitList(goFilePath)
		curr := split[len(split)-1]

		for curr != "" && curr != string(filepath.Separator) {
			prospect := filepath.Join(curr, "go.mod")
			if _, err := os.Stat(prospect); err == nil {
				hasPath := false

				for _, path := range modFilePaths {
					if path == prospect {
						hasPath = true
						break
					}
				}

				if !hasPath {
					modFilePaths = append(modFilePaths, prospect)
					posByModFilePath[prospect] = position{Start: file.FileStart, End: file.FileEnd}
				}

				break
			}

			curr = filepath.Dir(curr)
		}
	}

	for _, modFilePath := range modFilePaths {
		file, err := os.ReadFile(modFilePath)
		if err != nil {
			return nil, err
		}

		modFile, err := modfile.Parse(modFilePath, file, nil)
		if err != nil {
			return nil, err
		}

		for _, replace := range modFile.Replace {
			if pos, ok := posByModFilePath[modFilePath]; ok {
				pass.Report(analysis.Diagnostic{
					Pos:            pos.Start,
					End:            pos.End,
					Category:       "gomodreplace",
					Message:        fmt.Sprintf("replace directive found in %s: replace %s => %s\n", modFilePath, replace.Old.String(), replace.New.String()),
					SuggestedFixes: nil,
				})
			}
		}
	}

	return nil, nil
}
