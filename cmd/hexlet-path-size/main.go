package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:      "hexlet-path-size",                  //описательнй характер (название)
		Usage:     "print size of a file or directory", //описательный характер (что делает)
		ArgsUsage: "<path>",                            //справка о том что программа ожидает путь
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
