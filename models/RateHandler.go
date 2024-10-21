package models

import (
	"currency-rate/models/clients"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/nleeper/goment"
)

func LoadLatestRate() {
	// STEP 1: Get history file from the folder
	historyFileNames := GetHistoryFileNames()
	fileName := historyFileNames.GetLatestFileName()
	f, err := os.OpenFile("history/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	// STEP 2: Get the latest line
	logFile := NewLogFileWithOsFile(f)
	lastLine, _ := logFile.GetLastLine()

	// STEP 3: Save the log rate
	LastLogLine = NewLogLineFromString(lastLine)
}

func SaveIfDifferent() {
	// STEP 1: Get the rate
	updatedRate, gaitameRate, err := getRateIfDifferent(LastLogLine.Rate)
	if err != nil || updatedRate == 0 {
		return
	}

	// STEP 2: Open file logger
	g, _ := goment.New()
	date := g.Format("YYYYMMDD")
	fileName := fmt.Sprintf("history/%s.log", date)
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil || logFile == nil {
		return
	}

	// STEP 3: Print and output log
	fileLogger := log.NewWithOptions(logFile, log.Options{ReportTimestamp: true})
	outputString := fmt.Sprintf("%.4f, %.3f, %.4f", updatedRate, gaitameRate, updatedRate*gaitameRate)
	fileLogger.Printf(outputString)
	log.Printf(outputString)

	// STEP 4: Update last rate
	LastLogLine = NewLogLine(updatedRate, gaitameRate)
}

func getRateIfDifferent(lastRate float64) (float64, float64, error) {
	// STEP 1: Get bot web page
	webPage, err := clients.GetWebPage()
	if err != nil {
		return 0, 0, err
	}

	// STEP 2: Parse to JPYTWD rate
	strRate := WebParser(webPage)

	// STEP 3: Parse to float
	rate, err := strconv.ParseFloat(strRate, 64)
	if err != nil {
		return 0, 0, err
	}

	// STEP 4: If the same, do nothing
	if rate == lastRate {
		return 0, 0, nil
	}

	// STEP 5: Get Gaitame text
	gaitameText, err := clients.GetRateData()
	if err != nil {
		return 0, 0, err
	}

	// STEP 6: Parse USDJPY
	gaitameRate, err := ParseGaitameUSDJPY(gaitameText)
	if err != nil {
		return 0, 0, err
	}

	// STEP 7: Print all
	return rate, gaitameRate, nil
}
