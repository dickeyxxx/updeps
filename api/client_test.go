package api

import (
	"testing"

	"labix.org/v2/mgo"

	. "github.com/smartystreets/goconvey/convey"
)

type fakeMongo struct{}

func (f *fakeMongo) C(name string) *mgo.Collection {
	return &mgo.Collection{}
}

func TestNewClient(t *testing.T) {
	Convey("It creates an API client", t, func() {
		db := &fakeMongo{}
		client := NewClient(db)
		So(client, ShouldNotBeNil)
	})
}
