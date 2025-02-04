package main_test

import (
	"path"
	"testing"

	"github.com/infracost/infracost/internal/testutil"
)

func TestGenerateConfig(t *testing.T) {
	dir := testutil.CalcGoldenFileTestdataDirName()
	GoldenFileCommandTest(
		t,
		dir,
		[]string{
			"generate",
			"config",
			"--template-path",
			path.Join("./testdata", dir, "infracost.yml.tmpl"),
			"--project-path",
			path.Join("./testdata", dir),
		},
		nil)
}
