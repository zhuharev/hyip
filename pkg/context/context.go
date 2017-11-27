package context

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/pkg/buttons"
	"github.com/zhuharev/hyip/pkg/setting"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	base62 "github.com/pilu/go-base62"
	"github.com/zhuharev/tamework"

	"github.com/Unknwon/i18n"
)

type Context struct {
	*tamework.Context

	User *models.User

	i18n.Locale
}

func Contexter() tamework.Handler {
	return func(c *tamework.Context) {
		ctx := &Context{
			Context: c,
			Locale:  i18n.Locale{Lang: "ru-RU"},
		}

		user, err := models.Users.GetByTelegramID(c.UserID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// create new user
				// send registration message

				color.Green(c.Text)

				referID := base62.Decode(c.Text) / setting.App.SecretNumber

				user := new(models.User)
				user.Ref1 = uint(referID)

				err = models.Users.Create(user)
				if err != nil {
					c.Send(err.Error())
					return
				}

				err = models.BindTelegramID(user.ID, c.UserID)
				if err != nil {
					c.Send(err.Error())
					return
				}

				c.Keyboard.AddReplyButton(buttons.LanguageButtonsRU[0]).
					AddReplyButton(buttons.LanguageButtonsRU[1])
				c.Send("Выберите ваш язык\n\nChoose your language:")

				c.Exit()
				return
			}
		}
		ctx.User = user
		ctx.Locale.Lang = user.LangString()
		c.Map(ctx)
	}
}
