package helpers

import(
	"os"
	"strings"
)

func EnforceHTTP(url string) string {

	if url[:5] != "https" {
		return "https://" + url
	}

	return url

}

func RemoveDomainError(url string) bool {

	if url == os.Getenv("Domain"){
		return false
	}

	newURL := strings.Replace(url, "https://", "", 1)
	newURL = strings.Replace(newURL, "http://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN"){
		return false
	}

	return true
}