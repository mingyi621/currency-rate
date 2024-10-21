package models

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func WebParser(text string) string {
	reg, _ := regexp.Compile(`(?ms)^.*?日圓 \(JPY\).*?本行現金賣出.*?>(.*?)<.*$`)
	return reg.ReplaceAllString(text, "$1")
}

type GaitameElement struct {
	Pair string  `json:"pair"`
	Bid  float64 `json:"bid"`
	Ask  float64 `json:"ask"`
}

type GaitameResponse struct {
	Status int              `json:"status"`
	Data   []GaitameElement `json:"data"`
}

func ParseGaitameUSDJPY(gaitameText string) (float64, error) {
	result := new(GaitameResponse)
	err := json.Unmarshal([]byte(gaitameText), result)
	if err != nil {
		return 0, err
	}
	for _, obj := range result.Data {
		if obj.Pair == "USDJPY" {
			return obj.Ask, nil
		}
	}
	return 0, fmt.Errorf("USDJPY not found")
}
