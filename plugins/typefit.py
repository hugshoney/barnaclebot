"""Plugin to get random quotes"""
import requests
import random


class Quotes:
    """Define word from Urban Dictionary"""
    def __init__(self, number=1):
        self.number = number

    def get_response(self):
        """Get full result and return"""
        url = f'https://type.fit/api/quotes'
        r = requests.get(url)
        response = r.json()

        return response

    def get_quote(self):
        """Get random quote"""
        list_quotes = self.get_response()
        quotes = []
        for _ in range(self.number):
            quotes.append(random.choice(list_quotes))
        return quotes
