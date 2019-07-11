package web

import (
	"fmt"
	"time"
)

type JsonTime time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (this *JsonTime) UnmarshalJSON(data []byte) (err error) {
	timeData := string(data)
	now, err := time.Parse(`"`+timeFormat+`"`, timeData)
	*this = JsonTime(now)
	return err
}
