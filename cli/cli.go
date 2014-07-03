package main

import (
	"fmt"
	"os"

	"labix.org/v2/mgo"

	"github.com/codegangsta/cli"
	"github.com/dickeyxxx/updeps/models"
)

func main() {
	mongo := mongoClient()
	defer mongo.Close()
	m := models.NewClient(mongo.DB("updeps"))

	app := cli.NewApp()
	app.Name = "updeps server"

	app.Commands = cliCommands(m)

	app.Run(os.Args)
}

//func fetchGithubInfo(owner string, name string, client *github.Repository) {
//repo, _, err := client.Repositories.Get(owner, name)
//return repo, err
//}

func cliCommands(models *models.Client) []cli.Command {
	return []cli.Command{
		{
			Name: "count",
			Action: func(c *cli.Context) {
				packages, err := models.AllPackages()
				if err != nil {
					panic(err)
				}

				fmt.Println("Number of packages:", len(packages))
			},
		},
		{
			Name: "refresh",
			Action: func(c *cli.Context) {
				models.RefreshPackages()
			},
		},
		{
			Name: "github",
			Action: func(c *cli.Context) {
				models.RefreshPackagesGithub()
			},
		},
	}
}

func mongoClient() *mgo.Session {
	mongo, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return mongo
}
