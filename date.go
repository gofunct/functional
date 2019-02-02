package functional

import (
	"time"
)

func Date(fmt string, date interface{}) string {
	return DateInZone(fmt, date, "Local")
}

func HtmlDate(date interface{}) string {
	return DateInZone("2006-01-02", date, "Local")
}

func HtmlDateInZone(date interface{}, zone string) string {
	return DateInZone("2006-01-02", date, zone)
}

func DateInZone(fmt string, date interface{}, zone string) string {
	var t time.Time
	switch date := date.(type) {
	default:
		t = time.Now()
	case time.Time:
		t = date
	case int64:
		t = time.Unix(date, 0)
	case int:
		t = time.Unix(int64(date), 0)
	case int32:
		t = time.Unix(int64(date), 0)
	}

	loc, err := time.LoadLocation(zone)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}

	return t.In(loc).Format(fmt)
}

func DateModify(fmt string, date time.Time) time.Time {
	d, err := time.ParseDuration(fmt)
	if err != nil {
		return date
	}
	return date.Add(d)
}

func DateAgo(date interface{}) string {
	var t time.Time

	switch date := date.(type) {
	default:
		t = time.Now()
	case time.Time:
		t = date
	case int64:
		t = time.Unix(date, 0)
	case int:
		t = time.Unix(int64(date), 0)
	}
	// Drop resolution to seconds
	duration := time.Since(t).Round(time.Second)
	return duration.String()
}

func ToDate(fmt, str string) time.Time {
	t, _ := time.ParseInLocation(fmt, str, time.Local)
	return t
}
