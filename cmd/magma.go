package main

import (
	"log"
	"os"

	"magma-cipher/internal/files"
	"magma-cipher/internal/key"
	"magma-cipher/internal/simple"

	"github.com/urfave/cli/v2"
)

var (
	flags = []cli.Flag{
		&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Usage: "Name of input `FILE`", Required: true},
		&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "Name of output `FILE`", Required: true},
		&cli.StringFlag{Name: "key", Aliases: []string{"k"}, Usage: "Name of key `FILE`", Required: true},
	}
)

func main() {
	app := &cli.App{
		Name:  "magma",
		Usage: "CLI tool for encrypting and decrypting files with GOST 28147-89 cipher (a.k.a. Magma)",
		Commands: []*cli.Command{
			{
				Name:   "encrypt",
				Usage:  "Encrypt a file",
				Flags: flags,
				Action: func(ctx *cli.Context) error {
					inputBytes, err := files.ReadInput(ctx.String("input"))
					if err != nil {
						return err
					}
					keyBytes, err := files.ReadInput(ctx.String("key"))
					if err != nil {
						return err
					}
					k, err := key.New(keyBytes)
					if err != nil {
						return err
					}

					encrypted := simple.Encrypt(inputBytes, k)

					if err := files.WriteOutput(ctx.String("output"), encrypted); err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:   "decrypt",
				Usage:  "Decrypt a file",
				Flags: flags,
				Action: func(ctx *cli.Context) error {
					inputBytes, err := files.ReadInput(ctx.String("input"))
					if err != nil {
						return err
					}
					keyBytes, err := files.ReadInput(ctx.String("key"))
					if err != nil {
						return err
					}
					k, err := key.New(keyBytes)
					if err != nil {
						return err
					}

					decrypted := simple.Decrypt(inputBytes, k)

					if err := files.WriteOutput(ctx.String("output"), decrypted); err != nil {
						return err
					}

					return nil
				},
			},
		},

	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
