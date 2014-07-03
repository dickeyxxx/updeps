package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/dickeyxxx/updeps/config"
	"github.com/dickeyxxx/updeps/pkg"
)

func main() {
	mongo, err := config.Mongo()
	if err != nil {
		panic(err)
	}
	defer mongo.Close()
	m := pkg.NewClient(mongo.DB("updeps"), config.Github())

	app := cli.NewApp()
	app.Name = "updeps server"

	app.Commands = cliCommands(m)

	app.Run(os.Args)
}

func cliCommands(pkg *pkg.Client) []cli.Command {
	return []cli.Command{
		{
			Name: "count",
			Action: func(c *cli.Context) {
				packages, err := pkg.AllPackages()
				if err != nil {
					panic(err)
				}

				fmt.Println("Number of packages:", len(packages))
			},
		},
		{
			Name: "refresh",
			Action: func(c *cli.Context) {
				pkg.RefreshPackages()
			},
		},
		{
			Name: "github",
			Action: func(c *cli.Context) {
				pkg.RefreshPackagesGithub()
			},
		},
	}
}
