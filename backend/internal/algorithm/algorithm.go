package algorithm

import (
	"errors"
	"log"
	"time"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/models"
	"github.com/bitcoin-trading-automation/internal/utils"
)

type Algorithm struct {
	Config config.Config
}

type IAlgorithm interface{}

func (alg *Algorithm) SMA(tickers []models.Ticker, startDate, endDate string) (float64, error) {
	log.Printf("SMA start startDate=%s endDate=%s", startDate, endDate)

	if startDate >= endDate {
		return 0, nil
	}

	endTimestamp, err := utils.GetTimestampFromDate(endDate)
	if err != nil {
		return 0, err
	}

	if time.Now().Unix() < endTimestamp {
		return 0, errors.New("endDate is future")
	}

	nDay, err := utils.DiffDays(startDate, endDate)
	if err != nil {
		return 0, err
	}
	log.Printf("SMA nDay=%d", nDay)

	closes := make(map[string]float64)
	for d := 0; d < nDay; d++ {
		date, err := utils.AddDate(startDate, d)
		if err != nil {
			return 0, err
		}

		close, err := GetDayClose(tickers, date)
		if err != nil {
			return 0, err
		}

		closes[date] = close
	}
	log.Printf("SMA closes=%v", closes)

	sum := 0.0
	for _, close := range closes {
		sum += close
	}
	log.Printf("SMA sum=%f", sum)

	sma := sum / float64(nDay)
	log.Printf("SMA sma=%f", sma)

	return sma, nil
}

func (alg *Algorithm) EMA(tickers []models.Ticker, startDate, endDate string) (float64, error) {
	// TODO
	return 0, nil
}

func RSI(tickers []models.Ticker, startDate, endDate string) (float64, error) {
	if startDate >= endDate {
		return 0, nil
	}

	endTimestamp, err := utils.GetTimestampFromDate(endDate)
	if err != nil {
		return 0, err
	}

	if time.Now().Unix() < endTimestamp {
		return 0, errors.New("endDate is future")
	}

	nDay, err := utils.DiffDays(startDate, endDate)
	if err != nil {
		return 0, err
	}

	diffs := make(map[string]float64)
	for d := 0; d < nDay; d++ {
		date, err := utils.AddDate(startDate, d)
		if err != nil {
			return 0, err
		}

		close, err := GetDayClose(tickers, date)
		if err != nil {
			return 0, err
		}

		open, err := GetDayOpen(tickers, date)
		if err != nil {
			return 0, err
		}

		diff := close - open
		diffs[date] = diff
	}

	upSum, downSum := 0.0, 0.0
	for _, diff := range diffs {
		if diff > 0 {
			upSum += diff
		}
		if diff < 0 {
			downSum += diff
		}
	}

	upAvg, downAvg := upSum/float64(nDay), downSum/float64(nDay)

	// TODO upAvg == downAvg の場合の処理(ほぼないはず)
	return 100.0 * upAvg / (upAvg - downAvg), nil
}

func GetDayOpen(tickers []models.Ticker, date string) (float64, error) {
	nextDate, err := utils.AddDate(date, 1)
	if err != nil {
		return 0, err
	}

	minTimestamp, err := utils.GetTimestampFromDate(nextDate)
	if err != nil {
		return 0, err
	}

	for _, ticker := range tickers {
		if ticker.Timestamp < minTimestamp && utils.GetDateFromTimestamp(ticker.Timestamp) == date {
			minTimestamp = ticker.Timestamp
		}
	}

	for _, ticker := range tickers {
		if ticker.Timestamp == minTimestamp {
			return ticker.Ltp, nil
		}
	}

	return 0, errors.New("no ticker")
}

func MACD(){
	// TODO
}

func BollingerBands(){
	// TODO
}

func FibonacciRetracement(){
	// TODO
}

func GetDayClose(tickers []models.Ticker, date string) (float64, error) {
	maxTimestamp, err := utils.GetTimestampFromDate(date)
	if err != nil {
		return 0, err
	}

	for _, ticker := range tickers {
		if ticker.Timestamp > maxTimestamp && utils.GetDateFromTimestamp(ticker.Timestamp) == date {
			maxTimestamp = ticker.Timestamp
		}
	}

	for _, ticker := range tickers {
		if ticker.Timestamp == maxTimestamp {
			return ticker.Ltp, nil
		}
	}

	return 0, errors.New("no ticker")
}

func GetLastLtp(tickers []models.Ticker) (float64, error) {
	var maxTimestamp int64 = 0
	for _, ticker := range tickers {
		if ticker.Timestamp > maxTimestamp {
			maxTimestamp = ticker.Timestamp
		}
	}

	for _, ticker := range tickers {
		if ticker.Timestamp == maxTimestamp {
			return ticker.Ltp, nil
		}
	}

	return 0, errors.New("no ticker")
}
