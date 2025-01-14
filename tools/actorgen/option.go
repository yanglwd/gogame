package main

import "flag"

type GenerateOptions struct {
	InputFile  string
	OutputFile string
	ChannelNum int
	Timeout    int
}

var options GenerateOptions

func init() {
	flag.StringVar(&options.InputFile, "input", "", "input file")
	flag.StringVar(&options.OutputFile, "output", "", "output file")
	flag.IntVar(&options.ChannelNum, "channel", 1024, "channel num")
	flag.IntVar(&options.Timeout, "timeout", 1000, "timeout ms")
}
