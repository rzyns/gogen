package main

import (
	"os"

	"github.com/rzyns/gogen/cli"
)

func main() {
	code := cli.Main(os.Args)
	os.Exit(code)
}
