package juejin

import (
	"fmt"

	"github.com/go-rod/rod"
)

func (j *JueJin) Login() error {
	page := rod.New().MustPages().First().MustNavigate("https://juejin.cn/")
	page.MustWaitLoad()

	fmt.Println(page.MustInfo().URL)
	return nil
}
