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
	"strings"
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
		result := en.Slang(m.Payload)
		var fullText string
		if len(result) != 0 {
			// If example exists, also send example text.
			if eg, exist := result["example"]; exist {
				fullText = fmt.Sprintf("<b>%s:</b>\n%s\n\n<i>\"%s\"</i>", strings.Title(m.Payload), result["definition"], eg)
			} else {
				fullText = fmt.Sprintf("<b>%s:</b>\n%s", strings.Title(m.Payload), result["definition"])
			}
			// Send definition of slang word with HTML parse mode.
			b.Send(m.Sender, fullText, tb.ModeHTML)
		} else {
			fullText = fmt.Sprintf("%q not found, try another day.", m.Payload)
			b.Send(m.Sender, fullText)
		}
	})

	// Send list result of definition and example use
	// of the word, when command /mean is issued.
	b.Handle("/mean", func(m *tb.Message) {
		// Call mean function from english package.
		// Take word from user as argument for function,
		// and return with slice of map.
		result := en.Mean(m.Payload)
		var fullText string
		// If result for the word exists, let's process.
		if len(result) != 0 {
			// Iterate list of result, and then process and
			// reply back to user.
			for _, item := range result {
				text := []string{}
				header := fmt.Sprintf("<b>%s as %s</b>", strings.Title(m.Payload), strings.Title(item.Speech))
				text = append(text, header)

				for _, word := range item.Definitions {
					definition := fmt.Sprintf("• %s", word.Mean)

					text = append(text, definition)
					if word.Example != "" {
						example := fmt.Sprintf("  <i>%q</i>", word.Example)
						text = append(text, example)
					}
				}
				fullText = strings.Join(text[:], "\n")

				b.Send(m.Sender, fullText, tb.ModeHTML)
			}

		} else {
			fullText = fmt.Sprintf("%q not found, try another day.", m.Payload)
			b.Send(m.Sender, fullText)
		}
	})

	// Send random quotes when /quote command is issued.
	b.Handle("/quote", func(m *tb.Message) {
		// Call quotes function to get random
		// map of quotes.
		quotes := en.Quotes()
		var fullText string
		// If author of quotes exists, append to bottom of text.
		if author, exist := quotes["author"]; exist {
			fullText = fmt.Sprintf("<i>%q</i>\n\n%s", quotes["text"], author)
		} else {
			fullText = fmt.Sprintf("<i>%q</i>", quotes["text"])
		}
		b.Send(m.Sender, fullText, tb.ModeHTML)
	})

	b.Start()
}
