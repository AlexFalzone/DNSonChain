package util

import "time"

func ParseExpirationDate(dateStr string) (int64, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
