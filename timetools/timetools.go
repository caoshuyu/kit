package timetools

import "time"

var BeijingLocation = time.FixedZone("Asia/Shanghai", 8*60*60)

const
(
	RFC3339     = "2006-01-02T15:04:05+08:00"
	TT          = "2006-01-02 15:04:05"
	TTDAY       = "2006-01-02"
	NOSPLITT    = "20060102150405"
	NOSPLITTDAY = "20060102"
)

const (
	ONE_HOURS_SECOND = 3600
	ONE_DAY_SECOND   = 86400
	ONE_WEEK_SECOND  = 604800
)

//获取当前时间
func GetNowDateTime() time.Time {
	return time.Now().In(BeijingLocation)
}

//获取当前时间之前时间
func GetBeforeTime(beforeSecond int) time.Time {
	return time.Unix(time.Now().In(BeijingLocation).Unix()-int64(beforeSecond), 0)
}

//获取当前时间之后时间
func GetNextTime(nextSecond int) time.Time {
	return time.Unix(time.Now().In(BeijingLocation).Unix()+int64(nextSecond), 0)
}

//时间类型转时间
func TimeStamp(layout string, value string) (time.Time, error) {
	t, err := time.ParseInLocation(layout, value, BeijingLocation)
	if nil != err {
		return time.Time{}, err
	}
	return t, nil
}

//时间类型转时间戳
func TimeStampUnix(layout string, value string) (int64, error) {
	t, err := TimeStamp(layout, value)
	if nil != err {
		return 0, err
	}
	return t.Unix(), nil
}

//获取零晨时间
func GetDayStart(t time.Time) time.Time {
	tm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, BeijingLocation)
	return tm
}

//获取一天最后一刻时间
func GetDayEnd(t time.Time) time.Time {
	tm := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 59, BeijingLocation)
	return tm
}

//时间戳转时间
func SecondToDateTime(second int64) time.Time {
	return time.Unix(second, 0).In(BeijingLocation)
}

//格式化时间
func TimeFormat(layout string, t time.Time) string {
	return t.Format(layout)
}
