from search_engine_parser.core.engines.duckduckgo import Search as DuckduckgoSearch

def duckduckgo_search(query):
    dsearch = DuckduckgoSearch()
    results = dsearch.search(query=query)
    return results
