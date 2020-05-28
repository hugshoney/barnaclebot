"""Plugin to show meaning of the word using Dictionary API"""
import requests


class DictionaryAPI:
    """Show meaning of the word from Urban Dictionary"""
    def __init__(self, word):
        self.word = word

    def get_response(self, version=2):
        """Get full result and return"""
        url = f'https://api.dictionaryapi.dev/api/v{version}/entries/en/{self.word}'
        r = requests.get(url)
        response = r.json()

        return response

    def get_phonetic(self):
        """Get phonetic text for the word"""
        dict_result = self.get_response()
        phonetic = dict_result[0]['phonetic']

        return phonetic

    def get_synonyms(self):
        """Get synonyms of the word"""
        dict_result = self.get_response(1)
        synonyms = dict_result[0]['meaning']['noun'][0]['synonyms']

        return synonyms

    def get_meaning(self, example=True):
        """Get meaning of the word"""
        dict_result = self.get_response()
        meanings = []
        if example == True:
            for meaning in dict_result[0]['meanings']:
                speech = meaning['partOfSpeech']
                definition = meaning['definitions'][0]['definition']
                if 'example' in meaning['definitions'][0]:
                    example = meaning['definitions'][0]['example']
                else:
                    example = None
                result = {
                    'speech': speech,
                    'definition': definition,
                    'example': example
                }
                meanings.append(result)
        else:
            for meaning in dict_result[0]['meanings']:
                speech = meaning['partOfSpeech']
                definition = meaning['definitions'][0]['definition']
                result = {
                    'speech': speech,
                    'definition': definition,
                }
                meanings.append(result)

        return meanings
