package helper

import (
	"kassa360/kassa360_go_dynamic_service/config"
	"time"
)

func TimeNow() string {
	return time.Now().UTC().Format(config.TimeStampLayout)
}

func DateNow() string {
	return time.Now().UTC().Format(config.DateLayout)
}

func TimeNowWithLocation(location string) string {
	loc := getLocation(location)
	return time.Now().In(loc).Format(config.TimeStampLayout)
}

func DateNowWithLocation(location string) string {
	loc := getLocation(location)
	return time.Now().In(loc).Format(config.TimeStampLayout)
}

func getLocation(location string) *time.Location {
	if location != "" {
		location = "Asia/Tashkent"
	}

	loc, err := time.LoadLocation(location)
	if err != nil {
		location = "Asia/Tashkent"
		loc, _ = time.LoadLocation(location)
	}

	return loc
}

func ToLocationTime(date, location string) (string, error) {
	loc := getLocation(location)

	t, err := time.ParseInLocation(config.TimeStampLayout, date, time.UTC)
	if err != nil {
		t, err = time.ParseInLocation(config.DateLayout, date, time.UTC)
		if err != nil {
			return date, err
		}

		return t.In(loc).Format(config.DateLayout), nil
	}

	return t.In(loc).Format(config.TimeStampLayout), nil
}

func ToUTC(date, location string) (string, error) {
	loc := getLocation(location)

	t, err := time.ParseInLocation(config.TimeStampLayout, date, loc)
	if err != nil {
		t, err = time.ParseInLocation(config.DateLayout, date, loc)
		if err != nil {
			return date, err
		}
	}

	return t.In(time.UTC).Format(config.TimeStampLayout), nil
}
