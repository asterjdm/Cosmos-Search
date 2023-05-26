import utils
import requests
from bs4 import BeautifulSoup
import config

def search(query):
    url_encode_query = utils.encode_url(query)
    url = "https://www.google.com/search?q=" + url_encode_query
    headers = config.headers
    with requests.Session() as s:
        s.post(url, headers=headers)
        response = s.get(url, headers=headers)
    soup = BeautifulSoup(response.text, 'html.parser')
    
    linksContainer = soup.find("div", {"id": "search"})
    linksContainers = linksContainer.find_all("div", {"class": "g"})
    links = []
    titles = []
    for a in linksContainers:
        linkTag = a.find("a", href=True)
        url = linkTag.get("href")
        title = linkTag.text
        links.append(url)
        titles.append(title)



    descriptionsParents = soup.find_all("div", {"data-sncf": "1"})
    descriptions = []
    for desc in descriptionsParents:
        descriptions.append(desc.find("div").find("span").text)

    resultsDict = []




search("linux")