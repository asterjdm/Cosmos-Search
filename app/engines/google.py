import app.utils as utils
import requests
from bs4 import BeautifulSoup

def google_search(query, page=0):
    try:
        url_encode_query = utils.encode_url(query)
        url = f"https://www.google.com/search?q={url_encode_query}&start={(page)*10}"
        headers = utils.get_random_header()

        with requests.Session() as s:
            s.post(url, headers=headers)
            response = s.get(url, headers=headers)
            response.raise_for_status()  # Raise an exception for non-successful responses

        soup = BeautifulSoup(response.text, 'html.parser')

        links_container = soup.find("div", {"id": "search"})
        if links_container is None:
            return {"error": "Failed to find links container in the HTML"}

        link_containers = links_container.find_all("div", {"class": "g"})
        if not link_containers:
            return {"error": "Failed to find link containers in the HTML"}

        descriptions_parents = soup.find_all("div", {"data-sncf": "1"})

        links = []
        titles = []
        descriptions = []

        for a in link_containers:
            link_tag = a.find("a", href=True)
            url = link_tag.get("href")
            title_tag = link_tag.find("h3")
            if title_tag:
                title = title_tag.getText()
            else:
                title = "No title"
            links.append(url)
            titles.append(title)

        for description_parent in descriptions_parents:
            try:
                description = description_parent.find("div").find("span").text
            except Exception as e:
                description = "No description"
                print(f"Error retrieving description: {str(e)}")

            descriptions.append(description)

        results_dict = [{"title": titles[i], "description": descriptions[i], "link": links[i]} for i in range(len(descriptions)) if links[i].startswith("http")]
        return results_dict

    except requests.RequestException as e:
        return {"error": f"Error making the HTTP request: {str(e)}"}
    except Exception as e:
        return {"error": f"Error occurred: {str(e)}"}

    return {"error": "Unknown error occurred"}  # Return an error dictionary in case of unknown errors
