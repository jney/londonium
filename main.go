package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")
	page.
		MustWindowFullscreen().
		MustWaitLoad().
		MustScreenshotFullPage("a.png")
	res, _ := proto.PageCaptureSnapshot{}.Call(page)
	utils.OutputFile("test.mhtml", res.Data)
}
