package main

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
	"io/ioutil"
	"os"
)

// Alpine doesn't do point releases, but if you are reading this, 3.8 downloads
// 3.8.1 or newer, which contains the security fix for this RCE:

const baseDockerImage = "alpine:3.8"

func main() {

	app := &cli.App{
		Name:    "godockerize",
		Usage:   "build Docker images from GO packages",
		Version: "0.0.2",
		Commands: []*cli.Command{
			{
				Name:        "build",
				Usage:       "build a Docker image from Go packages",
				ArgsUsage:   "[packages]",
				Description: "Build compiles and installs the packages by the import paths to /usr/local/bin\n   in the docker image. The first package is used as the entrypoint.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "tag",
						Aliases: []string{"t"},
						Usage:   "output Docker image and optionally a tag in the 'name:tag' format",
					},
					&cli.StringFlag{
						Name:  "base",
						Usage: "base Docker image name",
						Value: baseDockerImage,
					},
					&cli.StringSliceFlag{
						Name:  "env",
						Usage: "additional environment variables for Dockerfile",
					},
					&cli.StringSliceFlag{
						Name:  "go-build-flags",
						Usage: "additional flags to pass to go build",
					},
					&cli.BoolFlag{
						Name:  "dry-run",
						Usage: "only print generated Dockerfile",
					},
				},
				Action: doBuild,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}

func doBuild(c *cli.Context) error {
	_, err := os.Getwd()
	if err != nil {
		return err
	}

	//go111module, useModules := os.LookupEnv("GO111MODULE")

	args := c.Args()
	if args.Len() < 1 {
		return errors.New(`"godockerize build" requires 1 or more arguments`)
	}

	tmpDir, err := ioutil.TempDir("", "godockerrize")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)
	panic("")
}
