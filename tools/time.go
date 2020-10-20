package tools

import "time"

const Time1 = "2006年01月02日 15:04:05"
const Time2 = "2006-01-02 15:04:05"
const Time3 = "2006/01/02 15:04:05"
const Time4 = "2006.01.02 15:04:05"

//时间的格式化
func TimeFormat(t int64,format string)string{
	return time.Unix(t,0).Format(format)
}
