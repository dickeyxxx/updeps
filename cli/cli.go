package main

import (
	"fmt"

	"labix.org/v2/mgo"

	"github.com/dickeyxxx/updeps/models"
)

func main() {
	mongo, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer mongo.Close()
	m := models.NewClient(mongo.DB("updeps"))
	packages, err := m.AllPackages()
	if err != nil {
		panic(err)
	}

	fmt.Println("Number of packages:", len(packages))
}

//func fetchGithubInfo(owner string, name string, client *github.Repository) {
//repo, _, err := client.Repositories.Get(owner, name)
//return repo, err
//}
