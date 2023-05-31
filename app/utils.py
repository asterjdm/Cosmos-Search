import urllib.parse
import json
import random

def encode_url(url):
    return urllib.parse.quote(url)

def get_random_header():
    UA_LIST = [
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36"
    ]

    HEADER = {
        "referer": "referer: https://www.google.com",
        "user-agent": random.choice(UA_LIST)
    }
    return HEADER

