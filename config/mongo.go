package config

import (
	"os"

	"labix.org/v2/mgo"
)

func Mongo() (*mgo.Session, error) {
	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		mongoUri = "mongodb://localhost/updeps"
	}
	return mgo.Dial(mongoUri)
}
