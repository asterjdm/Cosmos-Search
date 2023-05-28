from flask import *
from engines.google import google_search
import config

app = Flask(__name__)


@app.route("/", methods=["GET"])
def index():
    return render_template("index.html", langCodes=config.langCodes)


@app.route("/search/", methods=["GET"])
def search():
    query = request.args.get('query')
    selectedLang = request.args.get('lang')
    if query:
        results = google_search(query)
        return render_template("index.html", results = results, query = query, langCodes=config.langCodes, lang=selectedLang)

    else:
        return redirect(url_for('index'))



if __name__ == "__main__":
    app.run(host="localhost", port=8080, debug=True)
