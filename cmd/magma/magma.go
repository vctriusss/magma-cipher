package main

import (
	"log"
	"os"

	"github.com/vctriusss/magma-cipher/internal/files"
	"github.com/vctriusss/magma-cipher/internal/key"
	"github.com/vctriusss/magma-cipher/internal/simple"
	"github.com/vctriusss/magma-cipher/internal/utils"

	"github.com/urfave/cli/v2"
)

var (
	encFlags = []cli.Flag{
		&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Usage: "Name of input `FILE`", Required: true},
		&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "Name of output `FILE`", Required: true},
		&cli.StringFlag{Name: "key", Aliases: []string{"k"}, Usage: "Name of key `FILE`", Required: false},
	}

	decFlags = []cli.Flag{
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
				Name:  "encrypt",
				Usage: "Encrypt a file",
				Flags: encFlags,
				Action: func(ctx *cli.Context) error {
					inputBytes, err := files.ReadInput(ctx.String("input"))
					if err != nil {
						return err
					}
					var k key.Key

					if ctx.String("key") == "" {
						k, err = key.Generate()
						if err != nil {
							return err
						}
						if err := files.WriteOutput("key.txt", utils.BytesToHex(k.Bytes())); err != nil {
							return err
						}

					} else {
						keyBytes, err := files.ReadInput(ctx.String("key"))
						if err != nil {
							return err
						}
						k, err = key.New(keyBytes)
						if err != nil {
							return err
						}
					}
					
					encrypted := simple.Encrypt(inputBytes, k)

					if err := files.WriteOutput(ctx.String("output"), encrypted); err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:  "decrypt",
				Usage: "Decrypt a file",
				Flags: decFlags,
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
