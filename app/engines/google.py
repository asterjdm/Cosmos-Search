import app.utils as utils
import requests
from bs4 import BeautifulSoup

def google_search(query, page = 0):
    url_encode_query = utils.encode_url(query)
    url = "https://www.google.com/search?q=%s&start=%d" % (url_encode_query, (page)*10)
    print(url)
    headers = utils.get_random_header()
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
        titleTag = linkTag.find("h3")
        if titleTag:
            title = titleTag.getText()
        else:
            title = "No title"
        links.append(url)
        titles.append(title)



    descriptionsParents = soup.find_all("div", {"data-sncf": "1"})
    descriptions = []
    for desc in descriptionsParents:
        try:
            descriptions.append(desc.find("div").find("span").text)
        except:
            descriptions.append("No description")

    resultsDict = []
    for i in range(0, len(descriptions)):
        try:
            if links[i].startswith("http"):               
                    resultsDict.append({"title": titles[i], "description": descriptions[i], "link": links[i]})
        except Exception as e:
            print("Error: " + str(e))
    
    return resultsDict

