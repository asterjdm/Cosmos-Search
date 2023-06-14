package engines

import(
	"github.com/trietmn/go-wiki"
	"fmt"
)

func getWiki(query string) (map[string]string, error) {
	search_result, _, err := gowiki.Search(query, 1, true)
	if err != nil {
		return nil, err
	}
	if len(search_result) <= 0 {
		return nil, nil
	}

	page, err := gowiki.GetPage(search_result[0], -1, false, true)
	if err != nil {
		return nil, err
	}
	summary, err := gowiki.Summary(search_result[0], 3, -1, false, true)
	if err != nil {
		return nil, err
	}


	wikiInfo := map[string]string{
		"Summary": summary,
	}
	return wikiInfo, nil
}
