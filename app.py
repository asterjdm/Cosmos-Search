from flask import *
from engines.google import google_search
from engines.wiki import getWikiSummary
app = Flask(__name__)


@app.route("/", methods=["GET"])
def index():
    return render_template("index.html")


@app.route("/search/", methods=["GET"])
def search():
    query = request.args.get('query')
    selectedLang = request.args.get('lang')
    actuPage = request.args.get('page')

    if query:
        if not actuPage:
            actuPage = 0
        actuPage = int(actuPage)
        if actuPage == 0:
            wiki = getWikiSummary(query, lang=selectedLang)
        else:
            wiki = []
        results = google_search(query, lang=selectedLang, page=actuPage)
        next_url = "/search?query=%s&lang=%s&page=%d" % (query, selectedLang, actuPage+1)
        return render_template("index.html", results=results, query=query, info=wiki, next_url=next_url, isSearch=True)

    else:
        return redirect(url_for('index'))



if __name__ == "__main__":
    app.run(debug=True)
