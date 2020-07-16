## Behind the Source Code
I always want to learn English but I'm so unmotivated, and I want to learn programming too, but I'm so lazy. Because of that, I'm trying to learn programming while making a Telegram bot to help myself learn English. Now, my English and programming skill that I have still sucks :sweat_smile:.

## Quick Start Setup
1. Clone the repository. `git clone https://github.com/hugshoney/barnacle-bot.git`
2. Open clone directory, and create Python environment. `python -m venv venv`
3. Activate your environment, and install dependencies. `pip install -r requirements.txt`
4. [Create bot](https://core.telegram.org/bots#6-botfather), and then add your Telegram bot token to `config.ini`.
5. Start the bot with `python main.py`.

## List Bot Command
- `/mean <word>`: Get definition of the word.
- `/slang <word>`: Get slang definition of the word.
- `/synonym <word>`: Get synonyms of the word.
- `/quote`: Get random quotes.
- `/random`: Get random word.

## Credit
- [Telebot](https://github.com/tucnak/telebot/) `MIT License`. Telegram bot framework that I use in this project.
- [Urban Dictionary](https://www.urbandictionary.com/). I use their API to get definition of sland word.
- [(unofficial) Google Dictionary API](https://dictionaryapi.dev/). I use this for `/mean` and `/synonym` command.
- [Type.fit](https://type.fit/). I ~~steal~~ use their API to get random quotes.
