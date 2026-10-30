package main

import (
	"code"
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
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			isHuman := cmd.Bool("human")
			isAll := cmd.Bool("all")
			isRecursive := cmd.Bool("recursive")

			path := cmd.Args().First()
			size, err := code.GetPathSize(path, isRecursive, isHuman, isAll)
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
