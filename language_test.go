package main

import (
	"testing"

	"github.com/dickeyxxx/updeps/languages"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

type fakeLanguageClient struct{}

func (c *fakeLanguageClient) List() []languages.Language {
	return []languages.Language{}
}

func TestLanguages(t *testing.T) {
	Convey("languageList", t, func() {
		client := &fakeLanguageClient{}
		context := &gin.Context{}
		context.Set("LanguageClient", client)
		//languageList(context)
	})
}
