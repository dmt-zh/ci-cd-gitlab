package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

func main() {
	commitReportName := flag.String("commit-report-name", "commit-report", "Name of the output commit report file")
	flag.Parse()

	tpl := `---COMMIT REPORT---

Author: {{ .author }}
Branch: {{ .branch }}
SHA:    {{ .sha }}
`
	t, err := template.New("commit-report").Parse(tpl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	f, err := os.OpenFile(*commitReportName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0640)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	data := map[string]string{
		"author": os.Getenv("CI_COMMIT_AUTHOR"),
		"branch": os.Getenv("CI_COMMIT_BRANCH"),
		"sha":    os.Getenv("CI_COMMIT_SHA"),
	}

	err = t.Execute(f, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}