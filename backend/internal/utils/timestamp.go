package utils

import "time"

const (
	DATE_FORMAT = "2006-01-02"
)

func GetDateFromTimestamp(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(DATE_FORMAT)
}

func DiffDays(startDate, endDate string) (int, error) {
	start, err := time.Parse(DATE_FORMAT, startDate)
	if err != nil {
		return 0, err
	}
	end, err := time.Parse(DATE_FORMAT, endDate)
	if err != nil {
		return 0, err
	}
	return int(end.Sub(start).Hours() / 24), nil
}

func GetTimestampFromDate(date string) (int64, error) {
	t, err := time.Parse(DATE_FORMAT, date)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func AddDate(date string, days int) (string, error) {
	t, err := time.Parse(DATE_FORMAT, date)
	if err != nil {
		return "", err
	}
	t = t.AddDate(0, 0, days)
	return t.Format(DATE_FORMAT), nil
}
