package utilities

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

func GetFileNameFromURL(url string) (string, error) {
	r, err := regexp.Compile("(?:\\/)([^\\?\\/]*)+(?:\\?|$)+")
	if err != nil {
		return "", err
	}

	result := r.FindStringSubmatch(url)
	if len(result) == 0 {
		return "", errors.New("Could not find a match for file")
	}

	return result[1], nil
}

func GetAssetPath(file string, pageID string) string {
	return strings.Replace(strings.Replace(os.Getenv("ASSET_PATH"), "$PAGE_URI$", pageID, -1), "$FILE_NAME$", file, -1)
}
