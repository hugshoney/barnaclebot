# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
"""Plugin to define words using Urban Dictionary"""
import requests


class UrbanDictionary:
    """Define word from Urban Dictionary"""
    def __init__(self, word):
        self.word = word

    def get_response(self):
        """Get full result and return"""
        url = f'http://api.urbandictionary.com/v0/define?term={self.word}'
        r = requests.get(url)
        response = r.json()

        return response

    def get_definition(self):
        """Get definition for words"""
        dict_result = self.get_response()
        definition = dict_result['list'][0]['definition']

        return definition

    def get_example(self):
        """Get example for words"""
        dict_result = self.get_response()
        example = dict_result['list'][0]['example']

        return example
