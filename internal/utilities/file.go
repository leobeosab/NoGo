package utilities

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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

func DownloadFile(download_url string, output_path string) {
	file, err := os.Create(output_path)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}

	resp, err := client.Get(download_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()
	fmt.Printf("Downloaded asset \nFrom: %s\nTo: %s\nSize: %d", download_url, output_path, size)
}

func MakeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModeDir|0755)
	}
	return nil
}

func WriteStringToFile(content string, filePath string, fileName string) error {
	if err := MakeDirectoryIfNotExists(filePath); err != nil {
		return err
	}

	f, err := os.Create(filePath + fileName)
	if err != nil {
		log.Println("Cannot open file: ", filePath)
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)
	return err
}
