package main

import (
	"strconv"
)

func generateShortUrl(hashedUrl *int) string {
	shortLink := "https://shortlink/" + strconv.Itoa(*hashedUrl)
	return shortLink
}
