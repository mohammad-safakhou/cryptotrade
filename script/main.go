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

var Kir = 0

func Calculate(name string) (string, float64) {
	freshRecords := readCsvFile("/home/mohammad/repos/mine/cryptotrade/script/BINANCE_BTCUSDTPERP, " + name + ".csv")[1:]
	//freshRecords = append(freshRecords, readCsvFile("/home/mohammad/repos/mine/cryptotrade/script/2.csv")[1:]...)

	var records []Records
	for _, freshRecord := range freshRecords {
		recordTime, _ := time.Parse("2006-01-02T15:04:05Z", freshRecord[0])
		//if recordTime.Year() < 2022 {
		//	continue
		//}
		//if recordTime.Month() < time.Now().Month() - 1 {
		//	continue
		//}
		if time.Now().Sub(recordTime).Hours() > 30*24 {
			continue
		}
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
	var prevIndex int
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

	Kir = 0
	var currentStatus int
	if prevRecord.Buy == 1 {
		currentStatus = 1
	} else if prevRecord.Sell == 1 {
		currentStatus = -1
	}
	win := 0
	loss := 0
	var outcome float64
	for index, record := range records {
		if currentStatus == 0 {
			if record.Buy == 0 && record.Sell == 0 {
				continue
			} else {
				prevRecord = record
				prevIndex = index
				if record.Buy == 1 {
					currentStatus = 1
				} else if record.Sell == 1 {
					currentStatus = -1
				}
				continue
			}
		}
		if record.Buy == 1 && record.Sell == 1 {
			Kir += 1
			if prevRecord.Buy == 1 {
				record.Buy = 0
			} else {
				record.Sell = 0
			}
		}

		if index-prevIndex > 2 {
			if prevRecord.Buy == 1 {
				//if record.Low < prevRecord.Close*10/10 {
				if false {
					diff := ((record.Open + record.Close) / 2) - (prevRecord.Open+prevRecord.Close)/2
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
					currentStatus = 0
				}
			} else if false {
				if record.High > prevRecord.Close*10/10 {
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
					currentStatus = 0
				}
			}
		}

		if record.Buy == 1 || record.Sell == 1 {
			var diff float64

			//recordAvg := (record.Close + record.Open) / 2
			//prevrecordAvg := (prevRecord.Close + prevRecord.Open) / 2
			//var an float64
			//if prevRecord.Close < prevRecord.Open {
			//	an = (prevRecord.Close + prevrecordAvg) / 2
			//} else {
			//	an = (prevRecord.Open + prevrecordAvg) / 2
			//}
			//
			//var an2 float64
			//if record.Close < record.Open {
			//	an2 = (record.Close + recordAvg) / 2
			//} else {
			//	an2 = (record.Open + recordAvg) / 2
			//}
			//
			//diff = an2 - an
			//Kir += 1
			diff = record.Close - prevRecord.Close



			if prevRecord.Buy == 1 {
				if diff > 0 {
					// taking profit
					win += 1
					outcome = outcome + diff
				} else {
					// losing profit
					loss += 1
					outcome = outcome + diff
				}
			} else {
				if diff > 0 {
					// losing profit
					loss += 1
					outcome = outcome - diff
				} else {
					// taking profit
					win += 1
					outcome = outcome - diff
				}
			}
			prevRecord = record
			prevIndex = index
		}

	}

	//fmt.Println(Kir)
	fmt.Println(fmt.Sprintf("\nstrategy with %s timeframe: \nwinrate = %d, lossrate = %d\n", name, win*100/(win+loss), loss*100/(win+loss)))

	return name, outcome
}
func main() {
	fmt.Println(Calculate("1"))
	fmt.Println(Calculate("3"))
	fmt.Println(Calculate("5"))
	fmt.Println(Calculate("15"))
	fmt.Println(Calculate("30"))
	fmt.Println(Calculate("60"))
	fmt.Println(Calculate("120"))
	fmt.Println(Calculate("180"))
	fmt.Println(Calculate("240"))
	fmt.Println(Calculate("1D"))
	fmt.Println(Calculate("1W"))
	fmt.Println(Calculate("1M"))
}
