package clients

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://rate.bot.com.tw/xrt?Lang=zh-TW"

func GetWebPage() (string, error) {
	response, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("the statusCode is %d", response.StatusCode)
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
