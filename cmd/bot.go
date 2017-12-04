package cmd

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/pkg/base"
	"github.com/zhuharev/hyip/pkg/bootstrap"
	"github.com/zhuharev/hyip/pkg/buttons"
	"github.com/zhuharev/hyip/pkg/context"
	"github.com/zhuharev/hyip/pkg/setting"
	routers "github.com/zhuharev/hyip/routers/t"

	"github.com/Unknwon/com"
	"github.com/Unknwon/i18n"
	"github.com/fatih/color"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	base62 "github.com/pilu/go-base62"
	"github.com/urfave/cli"
	"github.com/zhuharev/tamework"
)

var (
	bot *tgbotapi.BotAPI

	secretAdd        = 123123
	secretMultiplyer = 3

	states = map[int64]string{}
)

var (
	// CmdBot cli tool
	CmdBot = cli.Command{
		Name:   "bot",
		Action: runBot,
	}
)

func runBot(ctx *cli.Context) {

	err := bootstrap.GlobalInit(ctx.Bool("dev"))
	if err != nil {
		log.Fatalln(err)
	}

	RunBot(ctx)
}

// RunBot inits db connections, load settings and starts bot
func RunBot(ctx *cli.Context) {

	// err := i18n.SetMessage("ru-RU", "conf/locals/locale_ru-RU.ini")
	// if err != nil {
	// 	panic(err)
	// }
	// err = i18n.SetMessage("en-US", "conf/locals/locale_en-US.ini")
	// if err != nil {
	// 	panic(err)
	// }
	// i18n.SetDefaultLang("ru-RU")

	log.Println("Setting loaded")
	log.Println(setting.App)

	log.Println(i18n.Tr("ru-RU", "my_bank"))

	//tgbotapi.APIEndpoint = "http://localhost:4000/bot%s/%s"

	tw, err := tamework.New(setting.App.Telegram.BotToken)
	if err != nil {
		panic(err)
	}

	setting.App.Telegram.BotUsername = tw.Bot().Self.UserName

	color.Cyan("%s", tgbotapi.APIEndpoint)
	tw.Bot().Debug = true

	//tw.Bot().Send(tgbotapi.NewMessage(102710272, "hello"))
	tw.Locales = append(tw.Locales, i18n.Locale{Lang: "ru-RU"}.Tr, i18n.Locale{Lang: "en-US"}.Tr)

	tw.Use(tamework.Recovery())
	tw.Use(tamework.Waiterer())
	tw.Use(context.Contexter())

	tw.Use(func(c *context.Context) {
		color.Green("[%s] %s", c.Method, c.Text)
		if c.Update().ChannelPost != nil {
			color.Cyan("%+v", c.Update().ChannelPost)
			color.Cyan("%+v", c.Update().ChannelPost.Chat.ID)
		}
	})

	tw.Prefix("/start", handleStart)

	//банк
	tw.Text(buttons.MyBank, handleBank)
	tw.CallbackQuery(buttons.Deposit, handleRefill)
	tw.CallbackQuery("deposit_btc", handleBitcoinChoose)
	//tw.CallbackQuery(buttons.Reinvest, handleReinvest)
	//tw.CallbackQuery("reinvest_submit", handleReinvestSubmit)
	tw.CallbackQuery(buttons.HistoryOfTransactions, handleOutgoingTransactionHistory)
	tw.CallbackQuery(buttons.WithdrawRequest, handleTransactionHistory)
	tw.CallbackQuery(buttons.Calc, handleCalc)
	tw.CallbackQuery("deposit_"+buttons.Qiwi, routers.Qiwi)
	tw.CallbackQuery("ps_change_qiwi", routers.ChangeQiwiWalletID)
	//tw.CallbackQuery(buttons.WeRecommendMethod, routers.Recommend)
	//tw.CallbackQuery("manual_sum", routers.ManualSum)
	tw.CallbackQuery(buttons.InvestPlanList, handlePlans)

	// Partners
	tw.Text(buttons.Partners, handlePartners)

	// About
	tw.Text(buttons.AboutService, handleAboutService)

	//Settings
	tw.Text(buttons.Settings, handleSetting)
	tw.CallbackQuery(buttons.Language, handleLanguage)
	tw.Text(buttons.LanguageButtonsRU[0], handleLanguageChoose)
	tw.Text(buttons.LanguageButtonsRU[1], handleLanguageChoose)
	tw.CallbackQuery(buttons.InviteLink, handleRefLink)
	tw.CallbackQuery(buttons.MyReferrer, handleRef)
	tw.CallbackQuery(buttons.WebToken, handleWebtoken)

	tw.Run()
}

func getStartKeyBoard(lang string) *tamework.Keyboard {
	btn := i18n.Tr(lang, buttons.MyBank)
	btn2 := i18n.Tr(lang, buttons.Partners)
	btn3 := i18n.Tr(lang, buttons.AboutService)
	btn4 := i18n.Tr(lang, buttons.Settings)

	kb := tamework.NewKeyboard(nil).AddReplyButton(btn).
		AddReplyButton(btn2).
		AddReplyButton("").
		AddReplyButton(btn3).
		AddReplyButton(btn4)
	// tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn, btn2}, []tgbotapi.KeyboardButton{btn3, btn4})
	//kb.Selective = true
	return kb
}

//
func handleStart(c *context.Context) {
	c.Keyboard = getStartKeyBoard("ru-RU")
	c.Send("hello")
}

//
func handleBank(c *context.Context) {
	c.Keyboard = buttons.BankKB("ru-RU")
	c.Markdown(c.Render("my_bank.text", map[string]interface{}{"user": c.User}))
}

func handlePlans(c *context.Context) {
	plans, err := models.Plans.All()
	if err != nil {
		return
	}
	c.Data["plans"] = plans

	for _, v := range plans {
		c.Data["plan"] = v
		c.NewKeyboard(nil).AddCallbackButton(c.Tr("buy"), "buy_"+strconv.Itoa(int(v.ID)))
		c.Markdown(c.Render("my_bank.plan"))
	}
}

//
func handlePartners(c *context.Context) {

	refs, err := models.Users.AllReferrals(c.User.ID)
	if err != nil {
		c.Send(err.Error())
		return
	}

	c.Keyboard = buttons.PartnersKB("ru-RU")
	// TODO: remove this
	c.Data["user"] = c.User
	c.Data["counters"] = struct {
		PartnerNumber       int
		ActivePartnerNumber int
		InvestedTotalUSD    float64
		Profit1LvlUSD       float64
		ProfitLowLvlUSD     float64
		InvestedTotalBTC    float64
		Profit1LvlBTC       float64
		ProfitLowLvlBTC     float64
	}{
		len(refs),
		100,
		100,
		1000,
		2000,
		10,
		100,
		10,
	}
	_, err = c.Markdown(c.Render("partners.text"))
	log.Println(err)
}

//
func handleAboutService(c *context.Context) {
	c.Keyboard = buttons.AboutKB(c.User.LangString())
	c.Data["ProjectName"] = c.Tr("project_name")
	_, err := c.Markdown(c.Render("about.text"))
	log.Println(err)
}

//
func handleSetting(c *context.Context) {
	c.Keyboard = buttons.SettingsKB(c.User.LangString())
	c.Data["Login"] = c.User.Name
	c.Markdown(c.Render("settings.text"))
}

//
func handleLanguage(c *context.Context) {
	c.Keyboard.AddReplyButton(buttons.LanguageButtonsRU[0]).
		AddReplyButton(buttons.LanguageButtonsRU[1])
	c.Send(c.Tr("choose_lang"))
}

//
func handleLanguageChoose(c *context.Context) {
	switch c.Text {
	case buttons.LanguageButtonsRU[0]:
		c.User.Lang = 0
	case buttons.LanguageButtonsRU[1]:
		c.User.Lang = 1
	}
	err := models.Users.Save(c.User, models.UserDBSchema.Lang)
	if err != nil {
		c.Send("error")
		return
	}
	c.Keyboard = getStartKeyBoard(c.User.LangString())
	c.Send(c.Tr("choosen_lang"))
}

//
func handleRefill(c *context.Context) {
	c.Keyboard.AddCallbackButton(c.Tr(buttons.Bitcoin), "deposit_btc").
		//	AddCallbackButton(c.Tr(buttons.AdvancedCash), "deposit_"+buttons.AdvancedCash).
		//AddCallbackButton("").
		//AddCallbackButton(c.Tr(buttons.YandexMoney), "deposit_"+buttons.YandexMoney).
		AddCallbackButton(c.Tr(buttons.Qiwi), "deposit_"+buttons.Qiwi)
	c.Send(c.Tr("choose_payment_provider"))
}

// TODO: rename this func
func handleBitcoinChoose(c *context.Context) {

	msg := fmt.Sprintf("%s: *Ƀ%s*\n\n%s:\n\n`%s`",
		c.Tr("minimal_amount_for_deposit"),
		base.FmtAmount(10),
		c.Tr("transfer_amount_to_address"),
		"todo")
	//c.Keyboard = buttons.BicoinChooseKB(c.User.LangString())
	//log.Println(c.Keyboard.InlineKeyboardMarkup())
	_, err := c.Markdown(msg)
	if err != nil {
		color.Red("%s", err)
	}
}

//
// func handleWithdraw(update tgbotapi.Update) {
// 	textRU := "Какую сумму Вы хотите вывести на Ваш Bitcoin кошелек?"
// 	markup := new3KB(buttons.BitcoinValues)
//
// 	kbmsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, textRU)
// 	kbmsg.ReplyMarkup = &markup
// 	_, err := bot.Send(kbmsg)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
//
func handleTransactionHistory(c *context.Context) {
	c.Send(c.Tr("zero_count_transactions"))
}

func handleRef(c *context.Context) {
	id := c.User.Ref1
	var text string
	if id == 0 {
		text = c.Tr("you_dont_have_referrer")
	} else {
		text = c.Tr("your_referrer", id)
	}

	c.Markdown(text)
}

func handleRefLink(c *context.Context) {
	color.Green("%d", int(c.User.ID)*int(setting.App.SecretNumber))
	key := base62.Encode(int(c.User.ID) * int(setting.App.SecretNumber))
	textRU := fmt.Sprintf("Ваша пригласительная ссылка:\n`https://t.me/%s?start=%s`",
		setting.App.Telegram.BotUsername, key)

	c.Markdown(textRU)
}

func handleOutgoingTransactionHistory(c *context.Context) {
	// TODO:
	c.Send(c.Tr("zero_count_transactions"))
}

func handleReinvest(c *context.Context) {
	c.Keyboard.AddCallbackButton(c.Tr("submit"), "reinvest_submit")
	c.Send(c.Tr("reinvest.text"))
}

func handleReinvestSubmit(c *context.Context) {
	err := models.Users.Move(models.BalanceUSD, models.ReinvestUSD, c.User.ID)
	if err != nil {
		log.Println(err)
	}
	err = models.Users.Move(models.BalanceBTC, models.ReinvestBTC, c.User.ID)
	if err != nil {
		log.Println(err)
	}
	c.Answer("reinvested")
}

func handleCalc(c *context.Context) {
	c.Send("Введите сумму депозита:")
	update, done := c.Wait("cancel", 60*time.Second)
	if !done {
		return
	}
	sum := com.StrTo(update.Text()).MustInt64()
	c.Send("Введите количество дней депозита:")
	update, done = c.Wait("cancel", 60*time.Second)
	if !done {
		return
	}
	days := com.StrTo(update.Text()).MustInt64()

	simple := fmt.Sprint((float64(sum) * 10 / 100.0) * float64(days))
	//SUM = X * (1 + %)n
	hard := fmt.Sprint(base.FmtAmount(float64(sum) * math.Pow(1.0+10/100, float64(days))))
	text := fmt.Sprintf("Вы заработаете: %s\n\nС реинвестициями: %s", simple, hard)
	c.Send(text)

}

func handleSupport(update tgbotapi.Update) {
	textRU := `👔 Техническая поддержка
✅ Для решения любого технического или финансового вопроса Вам необходимо  написать администрации на почту trinity3mebot@gmail.com
Описать свою проблему, приложить скриншот и сообщить свой id (можете его увидеть нажав на кнопку Партнеры).
Все письма обрабатываются в порядке очереди, ожидайте и Вам ответят.
Все вопросы разрешимы! Давайте дружить, а Trinity ❤️ поможет нам в этом!
📢 Любую общую консультацию по работе проекта Вы может получить обратившись к модераторам чата`
	kbmsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, textRU)
	_, err := bot.Send(kbmsg)
	if err != nil {
		log.Println(err)
	}
}

func handleWebtoken(c *context.Context) {
	newPass, err := models.Users.SetRandomPassword(c.User.ID)
	if err != nil {
		color.Red("Err update password: %s", err)
		return
	}
	c.Markdown(c.Tr("your_new_password") + ": `" + newPass + "`")
}
