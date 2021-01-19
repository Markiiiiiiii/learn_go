package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1479970937:AAFKPdrFUmUCsnF4waRr4dZmLieY9Qj76-U")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			// ignore any non-Message Updates
			continue
		}
		weather := getinfo(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, weather)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// 获取城市天气
func getinfo(cityname string) string {
	URL := "https://way.jd.com/he/freeweather?appkey=79d6586692c7cb8127acf38815f8f6c8&city=" + cityname
	queryUrl := fmt.Sprintf("%s", URL)
	resp, err := http.Get(queryUrl)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	res, err := simplejson.NewJson(body)
	qlty, _ := res.Get("result").Get("HeWeather5").GetIndex(0).Get("aqi").Get("city").Get("qlty").String()
	return qlty
}
