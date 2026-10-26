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
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "format size to human",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			isHuman := cmd.Bool("human")

			path := cmd.Args().First()
			size, err := code.GetPathSize(path)
			if err != nil {
				return err
			}

			if isHuman {
				fmt.Printf("%s\t%s\n", code.SizeToHuman(size), path)
				return nil
			}

			fmt.Printf("%d\t%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
