package model

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

type UrlInfo struct {
	Url      string `json:"url"`
}

func (info *UrlInfo) GetLink() string {

	url := info.Url

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	return url
}

func GetUrl(minUrl string) (bool, *UrlInfo) {
	urlInfo := &UrlInfo{}

	_, err := bucket.Get(minUrl, urlInfo)
	if err != nil {
		return false, nil
	}

	return true, urlInfo
}

func PutUrl(urlInfo *UrlInfo) string {

	key, _ := GenerateRandomString()

	if _, err := bucket.Insert(key, urlInfo, 0); err != nil {
		panic(err.Error())
		return ""
	}

	return key
}

func GenerateRandomString() (string, error) {
	b, err := GenerateRandomBytes(7)
	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}