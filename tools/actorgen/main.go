package main

import (
	"flag"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()

	inputBase := filepath.Base(options.InputFile)
	inputDir := filepath.Dir(options.InputFile)
	outputName := strings.TrimSuffix(inputBase, ".go") + "_gen.go"

	tokenInfo := &FileInfo{}
	tokenInfo.inputFile = options.InputFile
	tokenInfo.outputFile = inputDir + "/" + outputName

	if err := readFile(tokenInfo); err != nil {
		panic(err)
	}

	if err := generate(tokenInfo); err != nil {
		panic(err)
	}

	exec.Command("gofmt", "-w", tokenInfo.outputFile).Run()
	exec.Command("goimports", "-w", tokenInfo.outputFile).Run()
}
