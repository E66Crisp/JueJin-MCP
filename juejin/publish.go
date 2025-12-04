package juejin

import (
	"context"
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

type PublishContent struct {
	Title   string
	Content string
}

var (
	PUBLISH_URL = "https://juejin.cn/editor/drafts/new?v=2"
)

func Publish(page *rod.Page, ctx context.Context, content PublishContent) error {
	fmt.Println("publish", content)
	p := page.Context(ctx)
	p.MustNavigate(PUBLISH_URL).MustWaitLoad()

	time.Sleep(10 * time.Second)
	return nil
}
