package main

import (
	base64 "encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func encode(s string, iterations int) string {
	for i := 0; i < iterations; i++ {
		s = base64.StdEncoding.EncodeToString([]byte(s))
	}
	return s
}

func decode(s string, iterations int) string {
	for i := 0; i < iterations; i++ {
		b, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			log.Fatal(err)
		}
		s = string(b)
	}
	return s
}
func main() {
	app := &cli.App{
		Name:      "Base 64 Go",
		Usage:     "Iteratively encode/decode strings using base64",
		ArgsUsage: "string",
		Commands: []*cli.Command{
			{
				Name:      "encode",
				Aliases:   []string{"e"},
				Usage:     "Encode a string using base64",
				ArgsUsage: "string",
				UsageText: "b64 encode <string>",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "iterations",
						Aliases: []string{"i"},
						Usage:   "Number of iterations",
						Value:   1,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println(encode(c.Args().First(), c.Int("iterations")))
					return nil
				},
			},
			{
				Name:      "decode",
				Aliases:   []string{"d"},
				Usage:     "Decode a string using base64",
				ArgsUsage: "string",
				UsageText: "b64 encode <string>",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "iterations",
						Aliases: []string{"i"},
						Usage:   "Number of iterations",
						Value:   1,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println(decode(c.Args().First(), c.Int("iterations")))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
