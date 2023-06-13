package web

import (
	"net/url"
)

function EncodeUrl(toEncode string) string{
	return url.QueryEscape(toEncode)
}