package reporter

import (
	"fmt"
	"io"

	"github.com/google/osv-scanner/internal/output"
	"github.com/google/osv-scanner/pkg/models"
)

type GHAnnotationsReporter struct {
	hasPrintedError bool
	stdout          io.Writer
	stderr          io.Writer
}

func NewGHAnnotationsReporter(stdout io.Writer, stderr io.Writer) *GHAnnotationsReporter {
	return &GHAnnotationsReporter{
		stdout:          stdout,
		stderr:          stderr,
		hasPrintedError: false,
	}
}

func (r *GHAnnotationsReporter) PrintError(msg string) {
	fmt.Fprint(r.stderr, msg)
	r.hasPrintedError = true
}

func (r *GHAnnotationsReporter) HasPrintedError() bool {
	return r.hasPrintedError
}

func (r *GHAnnotationsReporter) PrintText(msg string) {
	fmt.Fprint(r.stderr, msg)
}

func (r *GHAnnotationsReporter) PrintResult(vulnResult *models.VulnerabilityResults) error {
	return output.PrintGHAnnotationReport(vulnResult, r.stderr)
}
