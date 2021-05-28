package format

import "time"

//JSONTime ...
type JSONTime time.Time

const (
	//2006-01-02
	FormatDate = "2006-01-02"
	//2006-01
	FormatYearMonth = "2006-01"
	//2006
	FormatYear = "2006"
	//2006-01-02 15:04:05
	StandardFormat = "2006-01-02 15:04:05"
	//15:04:05
	FormatTime = "15:04:05"
)

//UnmarshalJSON ...
func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+StandardFormat+`"`, string(data), time.Local)
	*t = JSONTime(now)
	return
}

//MarshalJSON ...
func (t JSONTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(StandardFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, StandardFormat)
	b = append(b, '"')
	return b, nil
}

//String ...
func (t JSONTime) String() string {
	return time.Time(t).Format(StandardFormat)
}

//Format ...
func (t JSONTime) Format(fmt string) string {
	return time.Time(t).Format(fmt)
}

// StrToTime 字符串转time
func StrToTime(s string, format string, loc *time.Location) time.Time {
	t, _ := time.ParseInLocation(format, s, loc)
	return t
}

//TimePtrToString ...
func TimePtrToString(t *time.Time, format string) string {
	if format == "" {
		format = StandardFormat
	}
	if t != nil {
		return t.Format(format)
	}
	return ""
}

//TimeToString ...
func TimeToString(t time.Time, format string) string {
	if format == "" {
		format = StandardFormat
	}
	return t.Format("2006-01-02 15:04:05")
}
