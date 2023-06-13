package web

import (
	"net/url"
)

func UrlEncode(toEncode string) string{
	return url.QueryEscape(toEncode)
}