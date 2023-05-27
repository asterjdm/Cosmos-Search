from flask import *
from engines.google import google_search
app = Flask(__name__)


@app.route("/", methods=["GET"])
def index():
    return render_template("index.html")


@app.route("/search/", methods=["GET"])
def search():
    query = request.args.get('query')
    if query:
        results = google_search(query)
        resultHtml = ""
        for i in range(len(results)):
            link = results[i]["links"]
            title = results[i]["title"]
            description = results[i]["descriptions"]
            resultHtml += render_template("search_result.html", link=link, title=title, description=description)
        return resultHtml

    else:
        return redirect(url_for('index'))



if __name__ == "__main__":
    app.run(host="127.0.0.1", port=8080, debug=True)
