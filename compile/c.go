package compile

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

type CCompile struct {
}

func (c *CCompile) Compile(codeFile, exeFile string) (string, error) {
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c",
		fmt.Sprintf("gcc  -DONLINE_JUDGE  -O2  -w -fmax-errors=3  -std=c11 %s -lm -o %s", codeFile, exeFile))
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", errors.WithMessage(err, stderr.String())
	}
	return exeFile, nil
}
