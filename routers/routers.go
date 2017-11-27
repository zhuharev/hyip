package routers

import (
	"pure/api/socs/telegram/trinity/pkg/middleware"
	"pure/api/socs/telegram/trinity/pkg/wait"
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
