package routers

import (
	"fmt"
	"log"

	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/pkg/context"
	"github.com/zhuharev/hyip/pkg/payment_system/qiwi"

	"github.com/zhuharev/boltutils"
)

func Qiwi(c *context.Context) {
	walletID, err := models.PaymentSystemGet(c.User.ID, models.Qiwi)
	if err != nil {
		if err == boltutils.ErrNotFound {
			ChangeQiwiWalletID(c)
			Qiwi(c)
			return
		}
		log.Println(err)
		c.Send(err.Error())
		return
	}

	paymentURL := qiwi.MakePaymentURL(walletID, c.User.ID)
	c.Keyboard.AddURLButton("Перейти на qiwi.com", paymentURL).
		AddCallbackButton("").
		AddCallbackButton("Изменить кошелёк", "ps_change_qiwi")

	format := c.Tr("my_bank.qiwi_deposit")
	text := fmt.Sprintf(format, walletID, c.User.ID)
	c.Markdown(text)
}

func ChangeQiwiWalletID(c *context.Context) {
	c.Markdown("Введите ваш кошелёк (без +, например `70000000000`):")
	upd, done := c.Wait("cancel")
	if !done {
		c.Send("Время истекло")
		return
	}
	err := models.PaymentSystemSave(c.User.ID, models.Qiwi, upd.Text())
	if err != nil {
		log.Println(err)
		c.Send("error")
		return
	}
	c.Send("Кошелёк сохранён!")
}
