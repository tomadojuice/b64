package main

import (
	base64 "encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	clipboard "golang.design/x/clipboard"

	"github.com/urfave/cli/v2"
)

func encode(s string, iterations int, doCopy bool) string {
	for i := 0; i < iterations; i++ {
		s = base64.StdEncoding.EncodeToString([]byte(s))
	}
	if doCopy {
		copy(s)
	}
	return s
}

func decode(s string, iterations int, doCopy bool) string {
	for i := 0; i < iterations; i++ {
		b, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			log.Fatal(err)
		}
		s = string(b)
	}
	if doCopy {
		copy(s)
	}
	return s
}

func copy(s string) {
	clipboard.Write(clipboard.FmtText, []byte(s))

	// this delay is needed for Linux clipboard managers to properly pick up the new clipboard contents HACKY I KNOW
	time.Sleep(time.Millisecond * 20)
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
					&cli.BoolFlag{
						Name:    "copy",
						Aliases: []string{"c"},
						Usage:   "Copy to clipboard",
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println(encode(c.Args().First(), c.Int("iterations"), c.Bool("copy")))
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
					&cli.BoolFlag{
						Name:    "copy",
						Aliases: []string{"c"},
						Usage:   "Copy to clipboard",
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println(decode(c.Args().First(), c.Int("iterations"), c.Bool("copy")))
					return nil
				},
			},
		},
	}
	errClip := clipboard.Init()
	if errClip != nil {
		panic(errClip)
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
