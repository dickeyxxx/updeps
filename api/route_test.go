package api

import (
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

type fakeEngine struct {
	paths []string
}

func (e *fakeEngine) GET(path string, handlers ...gin.HandlerFunc) {
	e.paths = append(e.paths, path)
}
func (e *fakeEngine) POST(path string, handlers ...gin.HandlerFunc) {
	e.paths = append(e.paths, path)
}
func (e *fakeEngine) OPTIONS(path string, handlers ...gin.HandlerFunc) {
	e.paths = append(e.paths, path)
}
func (e *fakeEngine) Use(middlewares ...gin.HandlerFunc) {}

func TestRouter(t *testing.T) {
	Convey("Given an engine", t, func() {
		client := &Client{}
		engine := &fakeEngine{}
		Convey("It has routes", func() {
			client.Route(engine)
			So(engine.paths, ShouldNotBeEmpty)
		})
	})
}

func TestOptions(t *testing.T) {
	Convey("", t, func() {
		So(2, ShouldEqual, 2)
	})
}
