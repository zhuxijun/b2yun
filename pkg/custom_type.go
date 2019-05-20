package root

import (
	"fmt"
	"time"
)

// JSONTime 满足json格式化的time类型
type JSONTime time.Time

// MarshalJSON 自定义json转换
func (jt JSONTime) MarshalJSON() ([]byte, error) {
	t := time.Time(jt)
	str := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(str), nil
}

// Marshal 格式化为字符串
func (jt JSONTime) Marshal() string {
	t := time.Time(jt)
	return fmt.Sprintf("%s", t.Format("2006-01-02 15:04:05"))
}
