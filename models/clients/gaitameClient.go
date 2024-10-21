package clients

import (
	"fmt"
	"io"
	"net/http"
)

const RATE_URL = "https://navi.gaitame.com/v3/info/prices/rate"

func GetRateData() (string, error) {
	response, err := http.Get(RATE_URL)
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