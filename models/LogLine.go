package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/nleeper/goment"
)

type LogLine struct {
	Timestamp   *goment.Goment
	Rate        float64
	GaitameRate float64
	Multply     float64
}

func NewLogLine(rate float64, gaitameRate float64) *LogLine {
	g, _ := goment.New()
	return &LogLine{
		Timestamp:   g,
		Rate:        rate,
		GaitameRate: gaitameRate,
		Multply:     rate * gaitameRate,
	}
}

func NewLogLineFromString(str string) *LogLine {
	empty := &LogLine{}
	strs := strings.Split(str, " ")
	if len(strs) < 5 {
		log.Errorf("The str len is less than 5, str: %s", str)
		return empty
	}
	timestamp, err := goment.New(strs[0]+" "+strs[1], "YYYY/MM/DD HH:mm:ss")
	if err != nil {
		log.Errorf(err.Error())
		return empty
	}
	rate, err := strconv.ParseFloat(strs[2][0:len(strs[2])-1], 64)
	if err != nil {
		log.Errorf(err.Error())
		return empty
	}
	gaitameRate, err := strconv.ParseFloat(strs[3][0:len(strs[3])-1], 64)
	if err != nil {
		log.Errorf(err.Error())
		return empty
	}
	muliply, err := strconv.ParseFloat(strs[4], 64)
	if err != nil {
		log.Errorf(err.Error())
		return empty
	}
	return &LogLine{timestamp, rate, gaitameRate, muliply}
}

func (l *LogLine) Print() {
	if l.Rate == 0 {
		return
	}
	timestamp := l.Timestamp.Format("YYYY/MM/DD HH:mm:ss")
	fmt.Printf("%s %.4f, %.3f, %.4f\n", timestamp, l.Rate, l.GaitameRate, l.Multply)
}
