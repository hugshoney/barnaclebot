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
	// Initialize settings
	var botSettings tb.Settings
	// Add telegram token to settings.
	botSettings.Token = os.Getenv("TELEGRAM_TOKEN")
	// Check if Public URL exist in envrionment variable.
	publicURL := os.Getenv("PUBLIC_URL")
	// If public url exist, use webhook, else use longpoller.
	if publicURL != "" {
		port := os.Getenv("PORT")
		webhook := &tb.Webhook{
			Listen:   ":" + port,
			Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
		}
		botSettings.Poller = webhook
	} else {
		botSettings.Poller = &tb.LongPoller{Timeout: 10 * time.Second}
	}

	// Create bot.
	b, err := tb.NewBot(botSettings)
	if err != nil {
		panic(err)
	}

	// Closure function to create not found message.
	notFound := func(word string) string {
		return fmt.Sprintf("Result for %q is not found.\nTry again next time.", word)
	}

	// Send a message when command /start is issued.
	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Hi, buddy.")
	})

	// Send command help information when command /help is issued.
	b.Handle("/help", func(m *tb.Message) {
		helpText := "<b>List command:</b>\n<code>/mean &lt;word&gt;</code>: Get definition of the word.\n<code>/slang &lt;word&gt;</code>: Get slang definition of the word.\n<code>/synonym &lt;word&gt;</code>: Get synonyms of the word.\n<code>/quote</code>: Get random quotes."
		b.Send(m.Sender, helpText, tb.ModeHTML)
	})

	// Send definition and example use of slang word,
	// when command /slang is issued.
	b.Handle("/slang", func(m *tb.Message) {
		// Call slang function and take user word as argument and
		// return top result of definition (def) and example (eg).
		result := en.Slang(m.Payload)
		// Create full message variable.
		var fullText string
		// If definition result of slang word exist, process text of message.
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
			fullText = notFound(m.Payload)
			b.Send(m.Sender, fullText)
		}
	})

	// Send list result of definition and example use
	// of the word, when command /mean is issued.
	b.Handle("/mean", func(m *tb.Message) {
		// Call Dict function from english package.
		// Take word from user as argument for function,
		// and return with []Dictionary struct.
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
			fullText = notFound(m.Payload)
			b.Send(m.Sender, fullText)
		}
	})

	// Send list synonym of the word, when command /synonym is issued.
	// Full Disclosure: HARDCODE.
	b.Handle("/synonym", func(m *tb.Message) {
		// Call Dict function from english package.
		// Take word from user as argument for function,
		// and return with []Dictionary struct.
		result := en.Dict(m.Payload)
		// Create full message variable.
		var fullText string
		// Let's the hardcode begin.
		// If result for the word exists, let's process.
		if len(result) != 0 {
			text := []string{}
			// Create header for message.
			header := fmt.Sprintf("<b>%s</b>", strings.Title(m.Payload))
			// Apend header to slice message.
			text = append(text, header)
			// Closure function to check if string already exist in slice.
			contain := func(word string, list []string) bool {
				for _, item := range list {
					if word == item {
						return true
					}
				}
				return false
			}

			// Iterate list of item in Mean struct.
			for _, item := range result {
				// Iterate list of synonyms if exists.
				for _, word := range item.Definitions {
					if len(word.Synonym) != 0 {
						// Create speech of word for message.
						speech := fmt.Sprintf("\n<i>%s</i>", item.Speech)
						// Append this thing only if speech is not exist
						// in slice message.
						if !contain(speech, text) {
							// Append text header to text.
							text = append(text, speech)
						}

						// Join of all synonyms into one string.
						synonyms := fmt.Sprintf("• %s", strings.Join(word.Synonym[:], ", "))
						// Append synonym to slice message.
						text = append(text, synonyms)
					}
				}
			}

			// Stupid way to check if word synonyms exist in result.
			// If length of slice message more than or equal to 3, so
			// the result for sure have synonyms.
			if len(text) >= 3 {
				// Join all text as full message.
				fullText = strings.Join(text[:], "\n")
			} else {
				// If less than 3, just create not found message.
				fullText = notFound(m.Payload)
			}
			// Send full message to user.
			b.Send(m.Sender, fullText, tb.ModeHTML)
		} else {
			// If Synonym of the word not exist, send this text to inform user.
			fullText = notFound(m.Payload)
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
