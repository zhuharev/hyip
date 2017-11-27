package routers

import (
	"github.com/zhuharev/hyip/pkg/middleware"
	"github.com/zhuharev/hyip/pkg/wait"
)

func Recommend(c *middleware.Context) {
	c.Send("мы реккломендуем")
}

func ManualSum(c *middleware.Context) {
	c.Send("введите ёпта")
	text, done := wait.Wait(c.ChatID)
	if done {
		c.Send(text)
	}
}
