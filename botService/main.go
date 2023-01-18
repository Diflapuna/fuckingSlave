package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const URL = "http://127.0.0.1:6969/jokes"

func main() {

	bot, err := tgbotapi.NewBotAPI("5885447329:AAEXR6We27f4cu5An9bMAOzdTYa5ysaFOU4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for {
		select {
		case update := <-updates:
			if update.Message != nil {
				if update.Message.IsCommand() {
					switch update.Message.Command() {
					case "start":
						sendTextMessage("Пошел ты нахуй шутник)))", update.Message.Chat.ID, bot)
					case "joke":
						sendTextMessage(GetJoke(URL), update.Message.Chat.ID, bot)
					}
				}
			}
		}
	}
}

func sendTextMessage(text string, id int64, bot *tgbotapi.BotAPI) {
	m := tgbotapi.NewMessage(id, text)
	bot.Send(m)

}

func GetJoke(url string) string {
	type Joke struct {
		Joke string `json:"joke"`
	}
	joke := &Joke{}
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	resBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(resBody, joke)

	return joke.Joke
}
