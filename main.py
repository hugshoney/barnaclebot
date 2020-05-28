"""
Jessabelle is bot that created to help myself learn English.
"""

import logging
import sys
from configparser import ConfigParser

from telegram.ext import Updater, CommandHandler, MessageHandler, Filters

# Import plugins
sys.path.append('plugins')
from urbandictionary import UrbanDictionary
from dictionaryapi import DictionaryAPI

# Read and parse configuration file.
parser = ConfigParser()
parser.read('config.ini')

# Enable logging
logging.basicConfig(
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
    level=logging.INFO)

logger = logging.getLogger(__name__)


def start(update, context):
    """Send a message when the command /start is issued."""
    update.message.reply_text('Hello')


def help(update, context):
    """Send a message when the command /help is issued."""
    update.message.reply_text('Help!')


def echo(update, context):
    """Echo the user message."""
    update.message.reply_text(update.message.text)


def error(update, context):
    """Log Errors caused by Updates."""
    logger.warning('Update "%s" caused error "%s"', update, context.error)


def urban_dictionary(update, context):
    """Send a defined slang words when the command /slang is issued."""
    words = " ".join(context.args)
    dict = UrbanDictionary(words)
    result = f'<b>Top definition of <a href="https://www.urbandictionary.com/define.php?term={words}">{words.title()}</a>:</b>\n{dict.get_definition()}\n\n<i>Example:</i>\n{dict.get_example()}'
    update.message.reply_text(text=result,
                              parse_mode='html',
                              disable_web_page_preview=True)


def dictionary_api(update, context):
    """Send a meaning of the word when the command /mean is issued."""
    words = " ".join(context.args)
    dict = DictionaryAPI(words)
    for meaning in dict.get_meaning():
        speech = meaning['speech']
        definition = meaning['definition']
        example = meaning['example']

        if example is None:
            text = f'*Meaning of "{words.title()}" as {speech.title()}:*\n' + f'{definition}\n\n'
        else:
            text = f'*Meaning of "{words.title()}" as {speech.title()}:*\n' + f'{definition}\n\n' + f'_Example:_\n{example}'
        update.message.reply_text(text=text, parse_mode='markdown')


def dictionary_synonym(update, context):
    words = " ".join(context.args)
    dict = DictionaryAPI(words)
    list_synonyms = dict.get_synonyms()
    result = f'*Synonyms of "{words.title()}":*\n_{", ".join(list_synonyms)}_.'
    update.message.reply_text(text=result, parse_mode='markdown')


def main():
    """Start the bot."""
    updater = Updater(parser.get('core', 'token'), use_context=True)

    # Get the dispatcher to register handlers
    dp = updater.dispatcher

    # on different commands - answer in Telegram
    dp.add_handler(CommandHandler("start", start))
    dp.add_handler(CommandHandler("help", help))
    dp.add_handler(CommandHandler("slang", urban_dictionary))
    dp.add_handler(CommandHandler("mean", dictionary_api))
    dp.add_handler(CommandHandler("synonym", dictionary_synonym))

    # on noncommand i.e message - echo the message on Telegram
    dp.add_handler(MessageHandler(Filters.text, echo))

    # log all errors
    dp.add_error_handler(error)

    # Start the Bot
    updater.start_polling()

    # Run the bot until you press Ctrl-C or the process receives SIGINT,
    updater.idle()


if __name__ == '__main__':
    main()
