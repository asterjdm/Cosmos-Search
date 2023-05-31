import wikipedia
import requests
import json
from app.utils import encode_url

def get_wiki_main_image(title):
    url = 'https://en.wikipedia.org/w/api.php'
    data = {
        'action': 'query',
        'format': 'json',
        'formatversion': 2,
        'prop': 'pageimages|pageterms',
        "pithumbsize": 500,
        'titles': encode_url(title)
    }
    response = requests.get(url, params=data)
    json_data = json.loads(response.text)
    try:
        return json_data['query']['pages'][0]['thumbnail']['source']
    except KeyError:
        return "https://upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/150px-Wikipedia-logo-v2.svg.png"


def getWikiSummary(query):
    try:
        search_result = wikipedia.search(query, results=1)
        if len(search_result) >= 1:
            search_result = search_result[0]
            page = wikipedia.page(search_result, auto_suggest=False)
            title = page.title
            url = page.url
            summary = wikipedia.summary(title, sentences=2, auto_suggest=False)
            image = get_wiki_main_image(search_result)

            return [{"title": title, "url": url, "summary": summary, "image": image}]
        else:
            return []
    except (wikipedia.exceptions.DisambiguationError, wikipedia.exceptions.PageError):
        return []
