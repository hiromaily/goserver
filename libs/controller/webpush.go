package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/hiromaily/go-server/libs/parse"
	tm "github.com/hiromaily/go-server/libs/template"
	lg "github.com/hiromaily/golibs/log"
	//"github.com/googlechrome/push-encryption-go/webpush"
	conf "github.com/hiromaily/go-server/libs/config"
	webpush "github.com/sherclockholmes/webpush-go"
)

//const PrivateKey = "zmbOxbcxb_M7NsB7DNIXjp0rbOtLyTdNl8MSKJasbSk"
type WebPush struct {
	Subscription    *Subscription `json:"subscription"`
	ApplicationKeys *AppKeys      `json:"applicationKeys"`
	Data            *string       `json:"data"`
}

type Subscription struct {
	Endpoint       string     `json:"endpoint"`
	ExpirationTime *time.Time `json:"expirationTime"`
	Keys           Key        `json:"keys"`
}
type Key struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

type AppKeys struct {
	Public string `json:"public"`
}

//GET
func GetWebPush(res http.ResponseWriter, req *http.Request) {
	lg.Info("[WebPush]")
	//lg.Debugf("[req]%#v\n", req)

	//index
	tm.Execute(res, "webpush", nil)
}

//POST
func PostWebPush(res http.ResponseWriter, req *http.Request) {
	lg.Info("[PostWebPush]")

	var wp WebPush
	err := parse.ParseJson(req.Body, &wp)

	if err != nil {
		lg.Error(err)
		http.Error(res, err.Error(), 500)
		return
	}

	//debug
	lg.Debugf("Subscription: %v", wp.Subscription)
	if wp.Subscription != nil {
		lg.Debugf("Subscription.Endpoint: %s", wp.Subscription.Endpoint)
		lg.Debugf("Subscription.ExpirationTime: %v", wp.Subscription.ExpirationTime)
		lg.Debugf("Subscription.Keys.Auth: %s", wp.Subscription.Keys.Auth)
		lg.Debugf("Subscription.Keys.P256dh: %s", wp.Subscription.Keys.P256dh)
	}
	lg.Debugf("ApplicationKeys.Public: %s", wp.ApplicationKeys.Public)
	lg.Debugf("Data: %s", *wp.Data) //Data is not necessary

	//TODO: Save user data

	//TODO: when subscription is null, it means cancel subscription
	if wp.Subscription == nil {
		return
	}

	// send
	err = sendNotification(wp)
	if err != nil {
		lg.Error(err)
		http.Error(res, err.Error(), 500)
		return
	}
}

func sendNotification(wp WebPush) error {
	//sb, _ := json.Marshal(wp.Subscription)
	//sub, err := webpush.SubscriptionFromJSON(sb)
	//if err != nil {
	//	lg.Error(err)
	//	http.Error(res, err.Error(), 500)
	//	return
	//}
	//webpush.Send(nil, sub, "Yay! Web Push!", conf.GetConf().WebPush.PrivateKey)

	// Decode subscription
	s := webpush.Subscription{}

	sb, _ := json.Marshal(wp.Subscription)
	if err := json.NewDecoder(bytes.NewBuffer(sb)).Decode(&s); err != nil {
		return err
	}

	// make messages
	msg := "TestTestTest"

	// Send Notification
	_, err := webpush.SendNotification([]byte(msg), &s, &webpush.Options{
		Subscriber:      "",
		TTL:             0,
		VAPIDPrivateKey: conf.GetConf().WebPush.PrivateKey,
	})

	return err
}
