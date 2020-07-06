// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Barnacle Bot is Telegram bot that I use to help myself
// learn English while learning Go and vice versa.
// My English is still sucks, and my Go programming skill
// is still horrible.
package main

import (
	"fmt"
	"os"
	"time"

    en "github.com/hugshoney/barnaclebot/english"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Just general main function to start Telegram Bot.
// Copied from telebot README with some modification.
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

	// Send definition and example use of slang word,
	// when command /slang is issued.
	b.Handle("/slang", func(m *tb.Message) {
		// Call slang function and take user word as argument and
		// return top result of definition (def) and example (eg).
		def, eg := en.Slang(m.Payload)
		// Formating definition text with adding bold header on top of it.
		fullText := fmt.Sprintf("<b>Definition of %q:</b>\n%s", m.Payload, def)
		// Send definition of slang word with HTML parse mode.
		b.Send(m.Sender, fullText, tb.ModeHTML)
		// If example exists, also send example text.
		if eg != "" {
			fullEg := fmt.Sprintf("<b>Example of %q:</b>\n%s", m.Payload, eg)
			b.Send(m.Sender, fullEg, tb.ModeHTML)
		}
	})

	b.Start()
}
