package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type TokenInfo struct {
	token string
	ret   string
}

type FileInfo struct {
	inputFile   string
	outputFile  string
	packageName string
	structName  string
	tokens      []TokenInfo
}

func (t *FileInfo) String() string {
	return fmt.Sprintf("inputFile: %s\noutputFile: %s\npackageName: %s\nstructName: %s\ntokens: %v",
		t.inputFile, t.outputFile, t.packageName, t.structName, t.tokens)
}

func readFile(tokenInfo *FileInfo) error {
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
		token := TokenInfo{}
		token.token = strings.TrimSuffix(words[3], "()")
		for i := 4; i < len(words)-1; i++ {
			token.ret += words[i] + " "
		}
		tokenInfo.tokens = append(tokenInfo.tokens, token)
	}

	fmt.Println(tokenInfo)
	return nil
}
