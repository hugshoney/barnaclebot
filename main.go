package main

import (
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
	}

	// Send a message when command /start is issued.
	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Hi, buddy.")
	})

	// Send definition and example of slang word.
	b.Handle("/slang", func(m *tb.Message) {
		def, eg := slang(m.Payload)
		b.Send(m.Sender, def)
		b.Send(m.Sender, eg)
	})

	b.Start()
}
