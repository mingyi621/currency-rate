package main

import (
	"currency-rate/models"
	"time"
)

func main() {
	models.LoadLatestRate()
	models.LastLogLine.Print()
	for {
		models.SaveIfDifferent()
		time.Sleep(60 * time.Second)
	}
}
