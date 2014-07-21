package api

import (
	"testing"

	"github.com/dickeyxxx/updeps/languages"
	. "github.com/smartystreets/goconvey/convey"
)

type fakeLanguageClient struct{}

func (c *fakeLanguageClient) Get(slug string) *languages.Language {
	return nil
}

func (c *fakeLanguageClient) List() []*languages.Language {
	return []*languages.Language{}
}

func TestLanguages(t *testing.T) {
	Convey("languageList", t, func() {
		//languages := &fakeLanguageClient{}
		//api := &Client{languages: languages}
		//context := &gin.Context{}
		//api.languageList(context)
	})
}
