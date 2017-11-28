// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package advcash

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"time"

	"github.com/go-resty/resty"
)

func makeToken(pass string) string {
	now := time.Now()
	now = now.UTC()
	date := now.Format("20060102:15")

	untoken := pass + ":" + date

	hasher := sha256.New()
	hasher.Write([]byte(untoken))
	token := hex.EncodeToString(hasher.Sum(nil))
	return token
}

type TimeWithZero time.Time

func (t TimeWithZero) MarshalXML(e *xml.Encoder, se xml.StartElement) error {
	if time.Time(t).IsZero() {
		return nil
	}
	return e.EncodeElement(time.Time(t), se)
}

type HistoryRequest struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	XMLNS   string   `xml:"xmlns:soapenv,attr"`
	XMLNS2  string   `xml:"xmlns:wsm,attr"`

	// Header struct{} `xml:"soapenv:Header"`

	Body struct {
		Request struct {
			Arg0 struct {
				APIName string `xml:"apiName"`
				Token   string `xml:"authenticationToken"`
				Email   string `xml:"accountEmail"`
			} `xml:"arg0"`
			Arg1 struct {
				// ALL
				TransactionName string `xml:"transactionName,omitempty"`
				//<transactionStatus>COMPLETED</transactionStatus>
				TransactionStatus string       `xml:"transactionStatus,omitempty"`
				StartTimeFrom     TimeWithZero `xml:"startTimeFrom,omitempty"`
				StartTimeTo       TimeWithZero `xml:"startTimeTo,omitempty"`
				Count             int          `xml:"count,omitempty"`
			} `xml:"arg1"`
		} `xml:"wsm:history"`
	} `xml:"soapenv:Body"`
}

// Transaction was auto-generated from WSDL.
type AdvTransaction struct {
	Id                 string    `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	AccountName        string    `xml:"accountName,omitempty" json:"accountName,omitempty" yaml:"accountName,omitempty"`
	ActivityLevel      int       `xml:"activityLevel,omitempty" json:"activityLevel,omitempty" yaml:"activityLevel,omitempty"`
	Amount             float64   `xml:"amount,omitempty" json:"amount,omitempty" yaml:"amount,omitempty"`
	AmountInUSD        float64   `xml:"amountInUSD,omitempty" json:"amountInUSD,omitempty" yaml:"amountInUSD,omitempty"`
	Comment            string    `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	Currency           string    `xml:"currency,omitempty" json:"currency,omitempty" yaml:"currency,omitempty"`
	Direction          string    `xml:"direction,omitempty" json:"direction,omitempty" yaml:"direction,omitempty"`
	FullCommission     float64   `xml:"fullCommission,omitempty" json:"fullCommission,omitempty" yaml:"fullCommission,omitempty"`
	OrderId            string    `xml:"orderId,omitempty" json:"orderId,omitempty" yaml:"orderId,omitempty"`
	ReceiverEmail      string    `xml:"receiverEmail,omitempty" json:"receiverEmail,omitempty" yaml:"receiverEmail,omitempty"`
	Sci                bool      `xml:"sci,omitempty" json:"sci,omitempty" yaml:"sci,omitempty"`
	SenderEmail        string    `xml:"senderEmail,omitempty" json:"senderEmail,omitempty" yaml:"senderEmail,omitempty"`
	StartTime          time.Time `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
	Status             string    `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	TransactionName    string    `xml:"transactionName,omitempty" json:"transactionName,omitempty" yaml:"transactionName,omitempty"`
	UpdatedTime        time.Time `xml:"updatedTime,omitempty" json:"updatedTime,omitempty" yaml:"updatedTime,omitempty"`
	VerificationStatus string    `xml:"verificationStatus,omitempty" json:"verificationStatus,omitempty" yaml:"verificationStatus,omitempty"`
	WalletDestId       string    `xml:"walletDestId,omitempty" json:"walletDestId,omitempty" yaml:"walletDestId,omitempty"`
	WalletSrcId        string    `xml:"walletSrcId,omitempty" json:"walletSrcId,omitempty" yaml:"walletSrcId,omitempty"`
	TypeAttrXSI        string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace      string    `xml:"xmlns:objtype,attr,omitempty"`
}

// <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
//   <soap:Body>
//      <ns2:historyResponse xmlns:ns2="http://wsm.advcash/">
type HistoryResponse struct {
	XMLName xml.Name `xml:"Envelope"`

	Body struct {
		Ns struct {
			Transactions []AdvTransaction `xml:"return"`
		} `xml:"historyResponse"`
	} `xml:"Body"`
}

func (a Advcash) hist(lastFetchTime time.Time) (_ []AdvTransaction, err error) {
	var s HistoryRequest
	s.XMLNS = "http://schemas.xmlsoap.org/soap/envelope/"
	s.XMLNS2 = "http://wsm.advcash/"
	s.Body.Request.Arg0.APIName = a.apiName

	s.Body.Request.Arg0.Token = makeToken(a.apiPassword)
	s.Body.Request.Arg0.Email = a.email
	s.Body.Request.Arg1.TransactionStatus = "COMPLETED"
	s.Body.Request.Arg1.Count = 10
	if !lastFetchTime.IsZero() {
		s.Body.Request.Arg1.StartTimeFrom = TimeWithZero(lastFetchTime)
		s.Body.Request.Arg1.StartTimeTo = TimeWithZero(time.Now())
	}

	bts, err := xml.Marshal(s)
	if err != nil {
		return
	}

	link := "https://wallet.advcash.com/wsm/merchantWebService?wsdl"

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "text/xml; charset=\"utf-8\"").
		SetHeader("SOAPAction", "http://schemas.xmlsoap.org/soap/envelope/").
		SetHeader("User-Agent", "gowsdl/0.1").
		SetBody(bts).
		Post(link)
	if err != nil {
		return
	}

	bts = resp.Body()

	var res HistoryResponse

	err = xml.Unmarshal(bts, &res)
	if err != nil {
		return
	}

	return res.Body.Ns.Transactions, nil
}
