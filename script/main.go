package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

type Records struct {
	Time   time.Time `json:"time"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume float64   `json:"volume"`
	Buy    int       `json:"buy"`
	Sell   int       `json:"sell"`
}

func main() {
	freshRecords := readCsvFile("/home/mohammad/repos/mine/cryptotrade/script/1.csv")[1:]
	freshRecords = append(freshRecords, readCsvFile("/home/mohammad/repos/mine/cryptotrade/script/2.csv")[1:]...)

	var records []Records
	for _, freshRecord := range freshRecords {
		recordTime, _ := time.Parse("2006-01-02T15:04:05Z", freshRecord[0])
		recordOpen, _ := strconv.ParseFloat(freshRecord[1], 10)
		recordHigh, _ := strconv.ParseFloat(freshRecord[2], 10)
		recordLow, _ := strconv.ParseFloat(freshRecord[3], 10)
		recordClose, _ := strconv.ParseFloat(freshRecord[4], 10)
		recordVolume, _ := strconv.ParseFloat(freshRecord[9], 10)
		recordBuy, _ := strconv.Atoi(freshRecord[7])
		recordSell, _ := strconv.Atoi(freshRecord[8])
		records = append(records, Records{
			Time:   recordTime,
			Open:   recordOpen,
			High:   recordHigh,
			Low:    recordLow,
			Close:  recordClose,
			Volume: recordVolume,
			Buy:    recordBuy,
			Sell:   recordSell,
		})
	}

	//var start = 0
	var startIndex = 0
	var prevRecord Records
	for index, record := range records {
		if record.Buy == 1 {
			//start = 1
			startIndex = index
			prevRecord = record
			break
		}
		if record.Sell == 1 {
			//start = 0
			startIndex = index
			prevRecord = record
			break
		}
	}

	records = records[startIndex+1:]

	var outcome float64
	for _, record := range records {
		if record.Buy == 1 && record.Sell == 1 {
			if prevRecord.Buy == 1 {
				record.Buy = 0
			} else {
				record.Sell = 0
			}
		}

		diff := record.Close - prevRecord.Close
		if prevRecord.Buy == 1 {
			if diff > 0 {
				// taking profit
				outcome = outcome + diff
			} else {
				// losing profit
				outcome = outcome + diff
			}
		} else {
			if diff > 0 {
				// losing profit
				outcome = outcome - diff
			} else {
				// taking profit
				outcome = outcome - diff
			}
		}

		prevRecord = record
	}

	fmt.Println(outcome)
}
