from flask import *
from engines.google import google_search
app = Flask(__name__)


@app.route("/", methods=["GET"])
def index():
    return render_template("index.html", query="")


@app.route("/search/", methods=["GET"])
def search():
    query = request.args.get('query')
    if query:
        results = google_search(query)
        return render_template("index.html", results = results, query = query)

    else:
        return redirect(url_for('index'))



if __name__ == "__main__":
    app.run(host="127.0.0.1", port=8080, debug=True)
