package main

import "flag"

type GenerateOptions struct {
	InputFile  string
	OutputFile string
	ChannelNum int
	Timeout    int
	Debug      bool
	Async      bool
}

var options GenerateOptions

func init() {
	flag.StringVar(&options.InputFile, "input", "", "input file")
	flag.StringVar(&options.OutputFile, "output", "", "output file")
	flag.IntVar(&options.ChannelNum, "channel", 1024, "channel num")
	flag.IntVar(&options.Timeout, "timeout", 1000, "timeout ms")
	flag.BoolVar(&options.Debug, "debug", false, "debug")
	flag.BoolVar(&options.Async, "async", false, "async run actor")
}
