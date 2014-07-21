package main

import (
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRouter(t *testing.T) {
	Convey("Given an engine", t, func() {
		engine := gin.New()
		Convey("It doesnt crash", func() {
			Route(engine)
		})
	})
}
