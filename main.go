package main

import (
	"github.com/javijuol/git-remind/app/cli"
	"os"
)

func main() {
	cli.App.Run(os.Args)
}
