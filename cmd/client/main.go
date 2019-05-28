package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// gitCommit will be the hash that the binary was built from
// and will be populated by the Makefile
var gitCommit = ""

// version will be populated by the Makefile, read from
// VERSION file of the source code.
var version = ""

func main() {
	app := cli.NewApp()
	app.Name = "cs-tools"
	if gitCommit != "" {
		app.Version = fmt.Sprintf("%s, commit: %s", version, gitCommit)
	} else {
		app.Version = version
	}
	app.Usage = "Clean Source tools"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-level",
			Value: "error",
			Usage: "Log level (panic, fatal, error, warn, info, or debug)",
		},
	}

	app.Commands = []cli.Command{
		trainCommand,
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func before(context *cli.Context) error {
	logLevelString := context.GlobalString("log-level")
	logLevel, err := logrus.ParseLevel(logLevelString)
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	logrus.SetLevel(logLevel)

	return nil
}
