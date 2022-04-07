package main

import (
	"fmt"

	"github.com/Noisyboy-9/go-redis/cli"
	"github.com/Noisyboy-9/go-redis/container"
)

func main() {
	cmdParser := cli.New(container.New())

	if err := cmdParser.StartProgramLoop(); err != nil {
		fmt.Println(err)
	}
}
