package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func generate(tokenInfo *FileInfo) error {
	os.Remove(tokenInfo.outputFile)

	f, err := os.Create(tokenInfo.outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	var buf bytes.Buffer

	// header
	fmt.Fprintln(&buf, "// Code generated by actorgen .; DO NOT EDIT.")
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, tokenInfo.packageName)
	fmt.Fprintln(&buf, "import (")
	fmt.Fprintln(&buf, "\"time\"")
	fmt.Fprintln(&buf, "\"context\"")
	fmt.Fprintln(&buf, ")")

	// body
	fmt.Fprintln(&buf, "func NewActor(p *", tokenInfo.structName, ")*Actor {")
	fmt.Fprintln(&buf, "ctx, cancel := context.WithCancel(context.Background())")
	fmt.Fprintln(&buf, "return &Actor{")
	fmt.Fprintln(&buf, "p : p,")
	fmt.Fprintf(&buf, "mailbox: make(chan chan struct{}, %d),", options.ChannelNum)
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, "ctx: ctx,")
	fmt.Fprintln(&buf, "cancel: cancel,")
	fmt.Fprintln(&buf, "join: make(chan struct{}),")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf)

	fmt.Fprintln(&buf, "type Actor struct {")
	fmt.Fprintln(&buf, "p *", tokenInfo.structName)
	fmt.Fprintln(&buf, "mailbox chan chan struct{}")
	fmt.Fprintln(&buf, "ctx context.Context")
	fmt.Fprintln(&buf, "cancel context.CancelFunc")
	fmt.Fprintln(&buf, "join chan struct{}")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf)

	fmt.Fprintln(&buf, "func (a *Actor) Start() {")
	fmt.Fprintln(&buf, "go a.run()")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, "func (a *Actor) run() {")
	fmt.Fprintln(&buf, "runLoop:")
	fmt.Fprintln(&buf, "for {")
	fmt.Fprintln(&buf, "select {")
	fmt.Fprintln(&buf, "case <-a.ctx.Done():")
	fmt.Fprintln(&buf, "break runLoop")
	fmt.Fprintln(&buf, "case m := <-a.mailbox:")
	fmt.Fprintln(&buf, "close(m)")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf, "for m := range a.mailbox {")
	fmt.Fprintln(&buf, "close(m)")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf, "close(a.join) ")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, "func (a *Actor) Stop() {")
	fmt.Fprintln(&buf, "a.cancel()")
	fmt.Fprintln(&buf, "<-a.join")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf)

	for _, info := range tokenInfo.tokens {
		rets := generateDefaultReturnValue(info.ret)
		fmt.Fprintf(&buf, "func (a *Actor) %s() %s {", info.token, info.ret)
		fmt.Fprintln(&buf, "m := make(chan struct{})")
		fmt.Fprintln(&buf, "select {")
		fmt.Fprintln(&buf, "case a.mailbox <- m:")
		fmt.Fprintf(&buf, "case <-time.After(time.Duration(%d) * time.Millisecond):", options.Timeout)
		fmt.Fprintln(&buf, "return", rets)
		fmt.Fprintln(&buf, "default:")
		fmt.Fprintln(&buf, "close(m)")
		fmt.Fprintf(&buf, "return %s", rets)
		fmt.Fprintln(&buf, "}")
		fmt.Fprintln(&buf, "<-m")
		if len(info.ret) > 0 {
			fmt.Fprintf(&buf, "return a.p.%s()", info.token)
		} else {
			fmt.Fprintf(&buf, "a.p.%s()", info.token)
		}
		fmt.Fprintln(&buf, "}")
		fmt.Fprintln(&buf)
	}

	f.WriteString(buf.String())
	return nil
}

func generateDefaultReturnValue(ret string) string {
	words := strings.Split(ret, " ")
	if len(words) == 0 {
		return ""
	}
	val := ""
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		switch word {
		case "int8", "int16", "int32", "int64", "int":
			val += "0" + " "
		case "uint8", "uint16", "uint32", "uint64", "uint":
			val += "0" + " "
		case "float32", "float64":
			val += "0.0" + " "
		case "string":
			val += "\"\"" + " "
		case "bool":
			val += "false" + " "
		default:
			val += "nil" + " "
		}
	}
	return val
}
