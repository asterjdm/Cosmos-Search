import wikipedia

def getWikiSummary(query, lang="en"):
    try:
        if(type(lang) == None.__class__):
            lang = "en"
        wikipedia.set_lang(lang)
        search_result = wikipedia.search(query, results=1)
        search_result = search_result[0]
        page = wikipedia.page(search_result, auto_suggest=False)
        title = page.title
        url = page.url
        summary = wikipedia.summary(title, sentences=2, auto_suggest=False)
        if len(page.images) >= 1:
            image = page.images[0]
        else:
            image = "https://upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/800px-Wikipedia-logo-v2.svg.png"
        return [{"title": title, "url": url, "summary": summary, "image": image}]
    except wikipedia.exceptions.DisambiguationError as e:
        return []
    except wikipedia.exceptions.PageError as e:
        return []