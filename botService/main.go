package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/Diflapuna/fuckingSlave/factsService/models"
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
						sendTextMessage("https://static.1000.menu/img/content-v2/a5/13/42853/pelmeni-domashnie-klassicheskie_1580382950_16_max.jpg", update.Message.Chat.ID, bot)
					case "joke":
						sendTextMessage(GetJoke(URL), update.Message.Chat.ID, bot)
					case "eblan":
						sendTextMessage("https://www.youtube.com/watch?v=dQw4w9WgXcQ", update.Message.Chat.ID, bot)
					default:
						sendTextMessage("Ты че еблан? Нет такой команды", update.Message.Chat.ID, bot)
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
	joke := models.Joke{}
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	json.Unmarshal(resBody, &joke)

	return joke.Joke
}
