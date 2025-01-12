package main

import (
	"bytes"
	"fmt"
	"os"
)

func generate(tokenInfo *TokenInfo) error {
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
	fmt.Fprintln(&buf, "import \"context\"")

	// body

	fmt.Fprintln(&buf, "func NewActor(p *", tokenInfo.structName, ")*Actor {")
	fmt.Fprintln(&buf, "ctx, cancel := context.WithCancel(context.Background())")
	fmt.Fprintln(&buf, "return &Actor{")
	fmt.Fprintln(&buf, "p : p,")
	fmt.Fprintln(&buf, "mailbox: make(chan chan interface{}),")
	fmt.Fprintln(&buf, "ctx: ctx,")
	fmt.Fprintln(&buf, "cancel: cancel,")
	fmt.Fprintln(&buf, "join: make(chan interface{}),")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf)

	fmt.Fprintln(&buf, "type Actor struct {")
	fmt.Fprintln(&buf, "p *", tokenInfo.structName)
	fmt.Fprintln(&buf, "mailbox chan chan interface{}")
	fmt.Fprintln(&buf, "ctx context.Context")
	fmt.Fprintln(&buf, "cancel context.CancelFunc")
	fmt.Fprintln(&buf, "join chan interface{}")
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
	fmt.Fprintln(&buf, "close(a.join) ")
	fmt.Fprintln(&buf, "break runLoop")
	fmt.Fprintln(&buf, "case m := <-a.mailbox:")
	fmt.Fprintln(&buf, "m <- struct{}{}")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, "func (a *Actor) Stop() {")
	fmt.Fprintln(&buf, "a.cancel()")
	fmt.Fprintln(&buf, "<-a.join")
	fmt.Fprintln(&buf, "}")
	fmt.Fprintln(&buf)

	for _, token := range tokenInfo.tokens {
		fmt.Fprintf(&buf, "func (a *Actor) %s() {", token)
		fmt.Fprintln(&buf, "m := make(chan interface{})")
		fmt.Fprintln(&buf, "a.mailbox <- m")
		fmt.Fprintln(&buf, "<-m")
		fmt.Fprintln(&buf, "close(m)")
		fmt.Fprintf(&buf, "a.p.%s()", token)
		fmt.Fprintln(&buf, "}")
		fmt.Fprintln(&buf)
	}

	f.WriteString(buf.String())
	return nil
}
