package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		panic("Usage: actor <file_path>")
	}

	inputBase := filepath.Base(os.Args[1])
	inputDir := filepath.Dir(os.Args[1])
	outputName := strings.TrimSuffix(inputBase, ".go") + "_gen.go"

	tokenInfo := &FileInfo{}
	tokenInfo.inputFile = os.Args[1]
	tokenInfo.outputFile = inputDir + "/" + outputName

	if err := readFile(tokenInfo); err != nil {
		panic(err)
	}

	if err := generate(tokenInfo); err != nil {
		panic(err)
	}

	exec.Command("gofmt", "-w", tokenInfo.outputFile).Run()
	exec.Command("goimports", tokenInfo.outputFile).Run()
}
