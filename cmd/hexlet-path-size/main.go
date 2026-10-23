package main

import (
	"code/code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:      "hexlet-path-size",                  //описательнй характер (название)
		Usage:     "print size of a file or directory", //описательный характер (что делает)
		ArgsUsage: "<path>",                            //справка о том что программа ожидает путь
		Action: func(ctx context.Context, cmd *cli.Command) error {

			path := cmd.Args().First()
			size, err := code.GetPathSize(path)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}

//СДЕЛАТЬ КОММИТ!
