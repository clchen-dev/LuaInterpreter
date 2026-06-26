package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/clchen-dev/LuaInterpreter/interpreter"
)

var version = "dev"

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("luago", flag.ContinueOnError)
	flags.SetOutput(stderr)
	showVersion := flags.Bool("version", false, "print version information")
	flags.Usage = func() {
		_, _ = fmt.Fprintln(stderr, "Usage: luago [flags] <script.lua>")
		flags.PrintDefaults()
	}

	if err := flags.Parse(args); err != nil {
		return 2
	}
	if *showVersion {
		_, _ = fmt.Fprintf(stdout, "luago %s\n", version)
		return 0
	}
	if flags.NArg() != 1 {
		flags.Usage()
		return 2
	}

	if err := interpreter.New(stdout).RunFile(flags.Arg(0)); err != nil {
		_, _ = fmt.Fprintf(stderr, "luago: %v\n", err)
		return 1
	}
	return 0
}
