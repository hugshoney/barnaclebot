#!/usr/bin/env python3
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
"""Plugin to get random word"""
import random
import urllib.request


class Words:
    """Define word from Urban Dictionary"""
    def __init__(self, number=1):
        self.number = number

    def get_response(self):
        """Get full result and return"""
        url = 'https://svnweb.freebsd.org/csrg/share/dict/words?view=co&content-type=text/plain'
        r = urllib.request.urlopen(url)
        word = r.read().decode()
        response = word.splitlines()

        return response

    def get_word(self):
        """Get random quote"""
        list_words = self.get_response()
        word = random.choice(list_words)

        return word
