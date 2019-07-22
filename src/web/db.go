package web

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"
)

//mysql时间转换 类型一定要对
func DateStr(json JsonTime) string {
	return strings.Join([]string{"'", time.Time(json).Format("2006-01-02 15:04:05"), "'"}, "")
}

//拼接字符串加上''
func SpliceString(data interface{}) string {
	return strings.Join([]string{"'", data.(string), "'"}, "")
}

//支持string uint64 int64 int Float32 Float64 foreach
func Foreach(data interface{}, open string, close string, separator string) string {
	if reflect.ValueOf(data).IsNil() {
		Log.Error("data is nil")
		return ""
	}
	builder := strings.Builder{}
	if reflect.ValueOf(data).Len() > 0 {
		builder.WriteString(open)
		switch reflect.ValueOf(data).Index(0).Kind() {
		case reflect.String:
			array := data.([]string)
			for i, v := range array {
				if i > 0 {
					builder.WriteString(separator)
				}
				builder.WriteString("'")
				builder.WriteString(v)
				builder.WriteString("'")
			}
			break
		case reflect.Int:
			array := data.([]int)
			for i, v := range array {
				if i > 0 {
					builder.WriteString(separator)
				}
				builder.WriteString(strconv.Itoa(v))
			}
			break
		case reflect.Uint64:
			array := data.([]uint64)
			for i, v := range array {
				if i > 0 {
					builder.WriteString(separator)
				}
				builder.WriteString(strconv.FormatUint(v, 10))
			}
			break
		case reflect.Int64:
			array := data.([]int64)
			for i, v := range array {
				if i > 0 {
					builder.WriteString(separator)
				}
				builder.WriteString(strconv.FormatInt(v, 10))
			}
			break
		case reflect.Float32:
			array := data.([]float32)
			for i, v := range array {
				if i > 0 {
					builder.WriteString(separator)
				}
				builder.WriteString(strconv.FormatFloat(float64(v), 'E', -1, 32))
			}
			break
		case reflect.Float64:
			array := data.([]float64)
			for i, v := range array {
				if i > 0 {
					builder.WriteString(separator)
				}
				builder.WriteString(strconv.FormatFloat(v, 'E', -1, 64))
			}
			break
		}
		builder.WriteString(close)
	} else {
		Log.Error("data size is 0")
	}
	return builder.String()
}

var t *template.Template

func InitTemplate(path string, method SqlMethod) {
	fmt.Println("init sql template")
	t = template.New("template sql")
	funcMap := template.FuncMap{"str": SpliceString, "dateStr": DateStr, "foreach": Foreach}
	if method != nil {
		methods := method.AddMethod()
		for k, v := range methods {
			funcMap[k] = v
		}
	}
	t.Funcs(funcMap)
	_, err := t.ParseGlob(path)
	if err != nil {
		Log.Error(err)
	}
}

//新增template自定义方法
type SqlMethod interface {
	AddMethod() (method map[string]interface{})
}

//sql拼接
func SqlStr(name string, data interface{}) string {
	buf := new(bytes.Buffer)
	err := t.ExecuteTemplate(buf, name, data)
	if err != nil {
		Log.Error("sql template error:", err)
	}
	sql := strings.TrimSpace(strings.Replace(buf.String(), "\r\n", " ", -1))
	Log.Info(name, "[", sql, "]")
	return sql
}

//返回单条数据
func QueryOne(name string, data interface{}, result interface{}) error {
	err := DB.Get(result, SqlStr(name, data))
	return err
}

//返回多条数据
func QueryList(name string, data interface{}, result interface{}) error {
	err := DB.Select(result, SqlStr(name, data))
	return err
}
