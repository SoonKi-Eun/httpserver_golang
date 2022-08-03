package lib_module

import (
	"fmt"
	"time"
)

// ref : https://golangbyexample.com/current-timestamp-in-golang/
// ref : https://jeonghwan-kim.github.io/dev/2019/01/14/go-time.html
// ref : https://www.geeksforgeeks.org/fmt-sscanf-function-in-golang-with-examples/
// ref : https://stackoverflow.com/questions/40260599/difference-between-two-time-time-objects

func Get_CurrentTimeStamp() string {

	// 현재 시간 , 2020-01-24 09:43:42
	t := time.Now()
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func Get_UTC_CurrnetTime() string {

	// 현재 시간(2020-01-24 09:43:42) UTC , 2020-01-24 00:43:42
	t := time.Now().UTC()
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func Get_UTC_CurrnetTime_SecNone() string {

	// 현재 시간(2020-01-24 09:43:42) UTC , 2020-01-24 00:43:42
	t := time.Now().UTC()
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0)
}

func Get_Current_Unix() int64 {

	// 2020-01-24 09:43:42 -> 1579839222
	t := time.Now()
	return t.Unix()

}

func Get_Current_UnixNano() int64 {

	// 2020-01-24 09:43:42 -> 1579839222196901000
	t := time.Now()
	return t.UnixNano()
}

func Get_Current_UnixMilli() int64 {

	// 2020-01-24 09:43:42 -> 1579839222196
	t := time.Now()
	UnixMilli := int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
	return UnixMilli

}

func Get_Current_UnixMicro() int64 {

	// 2020-01-24 09:43:42 -> 1579839222196901
	t := time.Now()
	tUnixMicro := int64(time.Nanosecond) * t.UnixNano() / int64(time.Microsecond)
	return tUnixMicro

}

func TimeStamp_to_TimeST(timestamp string) time.Time {

	// 2020-01-24 09:43:42 -> time.Time = 2020-01-24 09:43:42 +0900 KST

	var (
		year   int = 0
		month  int = 0
		day    int = 0
		hour   int = 0
		minute int = 0
		second int = 0
		t      time.Time
	)

	_, err := fmt.Sscanf(timestamp, "%04d-%02d-%02d %02d:%02d:%02d", &year, &month, &day, &hour, &minute, &second)

	if err != nil {
		Logger.Error(err)
	} else {
		t = time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
	}

	return t

}

func TimeStamp01_to_TimeST(timestamp string) time.Time {
	// 20200124094342 -> time.Time = 2020-01-24 09:43:42 +0900 KST

	var (
		year   int = 0
		month  int = 0
		day    int = 0
		hour   int = 0
		minute int = 0
		second int = 0
		t      time.Time
	)

	_, err := fmt.Sscanf(timestamp, "%04d%02d%02d%02d%02d%02d", &year, &month, &day, &hour, &minute, &second)

	if err != nil {
		Logger.Error(err)
	} else {
		t = time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
	}

	return t

}

func TimeStamp_AddTime(timestamp string, add_sec int64) string {

	// 2020-01-24 09:43:42 , 60 -> 2020-01-24 09:44:42
	t := TimeStamp_to_TimeST(timestamp).Add(time.Second * time.Duration(add_sec))
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

}

func TimeStamp_SubTime(timestamp string, sub_sec int64) string {

	// 2020-01-24 09:43:42 , 60 -> 2020-01-24 09:42:42
	t := TimeStamp_to_TimeST(timestamp).Add(-1 * time.Second * time.Duration(sub_sec))
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

}

func Get_TimeGapNanoSec(s_time string, e_time string) int64 {

	//2020-01-24 09:42:42, 2020-01-24 09:43:42 -> 60000000000
	st := TimeStamp_to_TimeST(s_time)
	et := TimeStamp_to_TimeST(e_time)

	t_gap := et.Sub(st).Nanoseconds()

	return t_gap
}

func Get_TimeGapSec(s_time string, e_time string) int {

	//2020-01-24 09:42:42, 2020-01-24 09:43:42 -> 60

	t_gap := Get_TimeGapNanoSec(s_time, e_time) / 1000000000
	return int(t_gap)
}

func Convert_AWSTimeFormat(time string) string {

	//2020-01-24 09:42:42 -> 2020-01-24T09:42:42

	var (
		year          int = 0
		month         int = 0
		day           int = 0
		hour          int = 0
		minute        int = 0
		second        int = 0
		aws_time_form string
	)

	_, err := fmt.Sscanf(time, "%04d-%02d-%02d %02d:%02d:%02d", &year, &month, &day, &hour, &minute, &second)

	if err != nil {
		Logger.Error(err)
	} else {
		aws_time_form = fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d", year, month, day, hour, minute, second)
	}

	return aws_time_form
}

func Convert_ESTimeFormat(time string) string {

	//2020-01-24 09:42:42 -> 2020-01-24T09:42:42.000Z

	var (
		year         int = 0
		month        int = 0
		day          int = 0
		hour         int = 0
		minute       int = 0
		second       int = 0
		es_time_form string
	)

	_, err := fmt.Sscanf(time, "%04d-%02d-%02d %02d:%02d:%02d", &year, &month, &day, &hour, &minute, &second)

	if err != nil {
		Logger.Error(err)
	} else {
		es_time_form = fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d.000Z", year, month, day, hour, minute, second)
	}

	return es_time_form

}

func Convert_TimeST_To_ESTimeFormat(t_st time.Time) string {

	//2022-04-19 11:28:07 +0900 KST (time.Time) -> 2020-04-19T11:28:08.000Z
	t := t_st.UTC()
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d.000Z", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

}

func UnixMilli_to_TimeST(unix_mili float64) time.Time {

	// 1.650335287e+12 ->  2022-04-19 11:28:07 +0900 KST (time.Time)
	return time.Unix(int64(unix_mili)/1000, 0)

}
