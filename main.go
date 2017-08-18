package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"todoBot/config"
	"todoBot/messages"

	"github.com/MartinSahlen/go-cloud-fn/shim"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/melvinmt/firebase"
)

// Item is the item struct
type Item struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {

	file, err := config.Asset("config.yaml")
	if err != nil {
		panic(err)
	}
	c := config.FromYAML(file)

	bot := messages.NewBot(c.TelegramKey)
	bot.Client.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	ch := make(chan tgbotapi.Update, bot.Client.Buffer)
	bytes, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		panic(err)
	}

	ch <- update

	for update := range ch {

		if update.Message == nil {
			continue
		}

		url := "https://https://lyrical-beach-175121.firebaseio.com/todolist/" + strconv.FormatInt(update.Message.Chat.ID, 10) + "/item"
		ref := firebase.NewReference(url)

		r, _ := regexp.Compile("/new ")
		if r.MatchString(update.Message.Text) {
			bot.SendMessage(update, "This string matches")
			item := Item{
				Name: r.ReplaceAllString(update.Message.Text, ""),
			}

			err = ref.Write(item)
			if err != nil {
				panic(err)
			}

			itemURL := "https://https://lyrical-beach-175121.firebaseio.com/todolist/" + strconv.FormatInt(update.Message.Chat.ID, 10)

			itemRef := firebase.NewReference(itemURL).Export(false)

			newItem := Item{}

			if err = itemRef.Value(newItem); err != nil {
				panic(err)
			}

			bot.SendMessage(update, newItem.Name)

			continue
		}

		switch update.Message.Text {
		case "/hello":
			bot.SendMessage(update, "Hello, I am the todobot, you can send me the following commands.")
		case "/list":
			bot.SendMessage(update, "Here is your list")
		case "/wipe":
			bot.SendMessage(update, "Everything was wiped")
		case "help":
			bot.SendMessage(update, "I will help you")
		default:
			bot.SendMessage(update, "I didn't understand")
		}

	}

}

func main() {
	shim.ServeHTTP(handler)
}
