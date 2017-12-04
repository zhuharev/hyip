package buttons

import (
	"fmt"
	"strconv"

	"github.com/Unknwon/i18n"
	"github.com/zhuharev/tamework"
)

const (
	Partners     = "partners"
	MyBank       = "my_bank"
	AboutService = "about"
	Settings     = "settings"

	SubmitRU = "✅ Подтвердить"

	WeRecommendMethod = "we_recomend"

	ReinvestSubmit = "reinvest_submit"

	Deposit               = "deposit"
	Withdraw              = "withdraw"
	InvestPlanList        = "inves_plan_list"
	Reinvest              = "reinvest"
	HistoryOfTransactions = "history_of_transactions"
	WithdrawRequest       = "withdraw_request"
	Calc                  = "calc"

	// partners
	InviteLink = "invite_link"
	MyTeam     = "my_team"
	MyReferrer = "my_referrer"
	PromoStuff = "promo_stuff"

	// about
	Community   = "community"
	FAQ         = "faq"
	ForPartners = "for_partners"
	Support     = "support"
	Reviews     = "reviews"
	Recomend    = "recomend"
)

func BankKB(lang string) *tamework.Keyboard {
	t := i18n.Locale{Lang: lang}.Tr
	return tamework.NewKeyboard(nil).AddCallbackButton(t(Deposit), Deposit).
		AddCallbackButton(t(Withdraw), Withdraw).
		AddCallbackButton("").
		AddCallbackButton(t(InvestPlanList), InvestPlanList).
		//AddCallbackButton(t(Reinvest), Reinvest).
		//AddCallbackButton(t(HistoryOfTransactions), HistoryOfTransactions).
		AddCallbackButton("").
		//AddCallbackButton(t(WithdrawRequest), WithdrawRequest).
		AddCallbackButton(t(Calc), Calc)
}

func PartnersKB(lang string) *tamework.Keyboard {
	t := i18n.Locale{Lang: lang}.Tr
	return tamework.NewKeyboard(nil).AddCallbackButton(t(InviteLink), InviteLink).
		AddCallbackButton(t(MyTeam), MyTeam).
		AddCallbackButton("").
		AddCallbackButton(t(MyReferrer), MyReferrer) //.
	//AddCallbackButton(t(PromoStuff), PromoStuff)
}

var (
	Language          = "lang"
	PaymentRequisites = "payment_requisites"
	WebToken          = "web_token"
	SetReferrer       = "set_referrer"
)

func SettingsKB(lang string) *tamework.Keyboard {
	t := i18n.Locale{Lang: lang}.Tr
	return tamework.NewKeyboard(nil).AddCallbackButton(t(Language), Language).
		AddCallbackButton(t(PaymentRequisites), PaymentRequisites).
		AddCallbackButton("").
		AddCallbackButton(t(WebToken), WebToken) //.
	//AddCallbackButton(t(SetReferrer), SetReferrer)
}

var LanguageButtonsRU = []string{
	"🇷🇺 Русский",
	"🇺🇸 Английский",
}
var Bitcoin = "bitcoin"
var (
	AdvancedCash = "advanced_cash"
	YandexMoney  = "yandex_money"
	Qiwi         = "qiwi"
)
var PaymentMethodRU = []string{"💵 Bitcoin"}

func BicoinChooseKB(lang string) *tamework.Keyboard {
	t := i18n.Locale{Lang: lang}.Tr
	var amounts = []float64{
		0.009,
		0.024,
		0.047,
		0.095,
		0.24,
		0.47,
		0.7,
		1.5,
	}

	kb := tamework.NewKeyboard(nil)
	for i, amount := range amounts {
		if i > 0 && i%3 == 0 {
			kb.AddCallbackButton("")
		}
		kb.AddCallbackButton(fmt.Sprintf("Ƀ%s", strconv.FormatFloat(amount, 'f', -1, 32)), fmt.Sprintf("deposit_btc_%f", amount))
	}
	kb.AddCallbackButton(t("manual_input"), "deposit_manual_input")
	return kb
}

func AboutKB(lang string) *tamework.Keyboard {
	return nil
	t := i18n.Locale{Lang: lang}.Tr
	return tamework.NewKeyboard(nil).AddCallbackButton(t(Community), Community).
		AddCallbackButton(t(FAQ), FAQ).
		AddCallbackButton("").
		AddCallbackButton(t(ForPartners), ForPartners).
		AddCallbackButton(t(Support), Support).
		AddCallbackButton("").
		AddCallbackButton(t(Reviews), Reviews).
		AddCallbackButton(t(Recomend), Recomend)
}
