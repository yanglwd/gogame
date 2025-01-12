package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type TokenInfo struct {
	inputFile   string
	outputFile  string
	packageName string
	structName  string
	tokens      []string
}

func (t *TokenInfo) String() string {
	return fmt.Sprintf("inputFile: %s, outputFile: %s, packageName: %s, structName: %s, tokens: %v", t.inputFile, t.outputFile, t.packageName, t.structName, t.tokens)
}

func readFile(tokenInfo *TokenInfo) error {
	f, err := os.Open(tokenInfo.inputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(tokenInfo.packageName) < 1 && strings.Contains(scanner.Text(), "package") {
			tokenInfo.packageName = scanner.Text()
			continue
		}
		if len(tokenInfo.structName) < 1 && strings.Contains(scanner.Text(), "type") {
			words := strings.Split(scanner.Text(), " ")
			tokenInfo.structName = words[1]
			continue
		}
		if !strings.Contains(scanner.Text(), "func") {
			continue
		}

		words := strings.Split(scanner.Text(), " ")
		if !unicode.IsUpper(rune(words[3][0])) {
			continue
		}
		token := strings.TrimSuffix(words[3], "()")
		tokenInfo.tokens = append(tokenInfo.tokens, token)
	}

	fmt.Println(tokenInfo)
	return nil
}
