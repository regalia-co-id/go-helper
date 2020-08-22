package helpers

import "time"

func CompareTime(data1 string, data2 string) float64 {
	time1, _ := time.Parse("15:04:05", data1)
	time2, _ := time.Parse("15:04:05", data2)
	difference := time1.Sub(time2).Seconds()
	return difference
}

func inTimeSpan(start, end, check time.Time) bool {
	return (check.Equal(start) || check.After(start)) && check.Before(end)
}

func BetweenDate(request string, startDate string, endDate string) bool {
	date, _ := time.Parse("2006-01-02", request)
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	if inTimeSpan(start, end, date) {
		return true
	}

	return false
}

func BetweenDatetime(request string, startDate string, endDate string) bool {
	date, _ := time.Parse("2006-01-02 15:04:05", request)
	start, _ := time.Parse("2006-01-02 15:04:05", startDate)
	end, _ := time.Parse("2006-01-02 15:04:05", endDate)

	if inTimeSpan(start, end, date) {
		return true
	}

	return false
}
