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
		// Create full message variable.
		var fullText string
		// If result defintion of slang word exist, process
		// text of message.
		if len(result) != 0 {
			def := result["definition"]
			// If example exists, also include example text.
			if eg, exist := result["example"]; exist {
				fullText = fmt.Sprintf("<b>%s</b>\n%s\n\n<i>%s</i>", strings.Title(m.Payload), def, eg)
			} else {
				// If not exist, use only definition
				// of slang word as message.
				fullText = fmt.Sprintf("<b>%s</b>\n%s", strings.Title(m.Payload), def)
			}
			// Send definition of slang word with HTML parse mode.
			b.Send(m.Sender, fullText, tb.ModeHTML)
		} else {
			// If definition of slang word not exist,
			// send this text to inform user.
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
		result := en.Dict(m.Payload)
		// Create full message variable.
		var fullText string
		// If result for the word exists, let's process.
		if len(result) != 0 {
			// Iterate list mean from result, and then process text
			// and reply back to user.
			for _, item := range result {
				text := []string{}
				// Create header of message.
				header := fmt.Sprintf("<b>%s (%s)</b>", strings.Title(m.Payload), item.Speech)
				// Append text header to text.
				text = append(text, header)
				// Iterate list definitions from previous result.
				for _, word := range item.Definitions {
					// Process text for definition of the word.
					definition := fmt.Sprintf("• %s", word.Mean)

					// Append definition to text message.
					text = append(text, definition)
					// If example exists, add eg to message.
					if word.Example != "" {
						// Process example text for definition,
						// add two space in front of the word
						// for stupid way to formating text.
						example := fmt.Sprintf("  <i>%q</i>", word.Example)
						// Append example to text message.
						text = append(text, example)
					}
				}
				// Join all text as full message.
				fullText = strings.Join(text[:], "\n")
				// Send full message to user.
				b.Send(m.Sender, fullText, tb.ModeHTML)
			}

		} else {
			// If definition of the word not exist,
			// send this text to inform user.
			fullText = fmt.Sprintf("%q not found, try another day.", m.Payload)
			b.Send(m.Sender, fullText)
		}
	})

	// Send random quotes when /quote command is issued.
	b.Handle("/quote", func(m *tb.Message) {
		// Call quotes function to get random
		// map of quotes.
		quotes := en.Quotes()
		// Create full message variable.
		var fullText string
		text := quotes["text"]
		// If author of quotes exists, append to bottom of text.
		if author, exist := quotes["author"]; exist {
			fullText = fmt.Sprintf("<i>%q</i>\n\n%s", text, author)
		} else {
			// If author not exist, send quotes only.
			fullText = fmt.Sprintf("<i>%q</i>", text)
		}
		// Send full message to user.
		b.Send(m.Sender, fullText, tb.ModeHTML)
	})

	b.Start()
}
