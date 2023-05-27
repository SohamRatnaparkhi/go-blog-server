package utils

import (
	"net/http"
)

func GetCookie(header http.Header, key string) (string, error) {
	value := header.Get(key)

	if value == "" {
		return "", http.ErrNoCookie
	}

	return value, nil
}
