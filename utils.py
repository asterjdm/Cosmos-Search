import urllib.parse

def encode_url(url):
    return urllib.parse.quote(url)