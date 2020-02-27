package tool

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const DEBUG = true

type H map[string]interface{}

const (
	length = 63
)

var (
	idx int64
)

func IntsToLongToint(src ...int64) (dst []int64) {
	for _, dstInfo := range src {
		dst = append(dst, dstInfo)
	}
	return
}

func IntToLong(src []int64) (dst []int64) {
	idx = 1
	if len(src) < 1 {
		return
	}
	//求取最大数
	max := src[0]
	for _, val := range src {
		if val > max {
			max = val
		}
	}
	idx = max / length
	var idxString [][length]string
	var dstString []string
	if max%length == 0 {
		dstString = make([]string, idx)
		idxString = make([][length]string, idx)
		dst = make([]int64, idx)

	} else {
		dstString = make([]string, idx+1)
		idxString = make([][length]string, idx+1)
		dst = make([]int64, idx+1)
	}
	for _, val := range src {
		if 0 == val {
			continue
		}
		idy := val / length
		location := val % length
		if 0 == location {
			location = length
			idy -= 1
		}
		idxString[idy][location-1] = "1"
	}
	for index, val := range idxString {
		for _, locationVal := range val {
			if "" == locationVal {
				dstString[index] += "0"
				continue
			}
			dstString[index] += locationVal
		}
	}

	for index, val := range dstString {
		d, _ := strconv.ParseInt(val, 2, 64)
		dst[index] = d
	}
	return
}

/**
@params permissions []int64 角色权限集
@params power int64 等待验证的权限
*/
func Check(permissions []int64, power int64) (result bool) {
	result = false
	if power == 0 {
		return
	}
	idx := power / length
	if power%length == 0 {
		idx -= 1
	}
	if idx > int64(len(permissions)) {
		return
	}
	if permissions[idx]&IntsToLongToint(power)[idx] == 0 {
		return
	}
	result = true
	return
}

// 判断obj是否在target中，target支持的类型arrary,slice,map
func Contain(target interface{}, obj interface{}) bool {

	targetValue := reflect.ValueOf(target)

	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsNotBlank(str string) bool {
	if len(str) == 0 {
		return false
	}

	if strings.TrimSpace(str) == "" {
		return false
	}

	return true
}

func IsBlank(str string) bool {

	if len(str) == 0 {
		return true
	}

	if len(strings.TrimSpace(str)) == 0 {
		return true
	}

	return false
}

//过滤空白字符
func ReplaceBlank(str string) string {

	if len(str) == 0 {
		return str
	}

	charactors := make([]rune, 0)

	for _, v := range str {
		if unicode.IsSpace(v) {
			charactors = append(charactors, ' ')
		} else {
			charactors = append(charactors, v)
		}
	}

	return string(charactors)
}

func CompileTemplate(template, variable, sign string) string {

	if IsBlank(template) {
		return ""
	}

	template = ReplaceBlank(template) //去掉\t\n等特殊字符

	if IsBlank(variable) {
		return template + "【" + sign + "】"
	}

	variable = strings.TrimSpace(variable)

	variable = ReplaceBlank(variable) //去掉\t\n等特殊字符

	jo := make(map[string]interface{})

	err := json.Unmarshal([]byte(variable), &jo)
	if err != nil {
		log.Println("json.Unmarshal error: variable ==>", variable)
		fmt.Println("json.Unmarshal error: variable ==>", variable)
		return template + "【" + sign + "】"
	}

	for k, v := range jo {

		var str string

		switch v.(type) {
		case string:
			str = fmt.Sprintf("%s", v)
		case int:
			str = fmt.Sprintf("%d", v)
		case int32:
			str = fmt.Sprintf("%d", v)
		case int64:
			str = fmt.Sprintf("%d", v)
		case float64:
			str = fmt.Sprintf("%.2f", v)
		case float32:
			str = fmt.Sprintf("%.2f", v)
		}

		template = strings.Replace(template, "{"+k+"}", str, -1)
	}

	return template + "【" + sign + "】"
}

// Return part of a string
//
// Examples:
// string   start length return
// "abcdef" 0     1      a
// "abcdef" 1     2      bc
// "abcdef" -2    1      e
// "abcdef" -2    0      ef
// "abcdef" 1     -2     bcd
// "abcdef" -3    -2     d
// "abcdef" -2    -4     ""
// "abcdef" -20   -4     ab
// "abcdef" -20   -10    ""
func SubStr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl + start
		if start < 0 {
			start = 0
		}
	}

	if 0 == length {
		end = rl
	} else if 0 > length {
		end = rl + length
		if end < 0 {
			end = 0
		}
	} else {
		end = start + length
		if end > rl {
			end = rl
		}
	}

	if start > end {
		return ""
	}

	return string(rs[start:end])
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {

	str := time.Time(t).Format("2006-01-02 15:04:05")

	return []byte(str), nil
}

func Truncation(t float64) float64 {
	t, _ = strconv.ParseFloat(strconv.FormatFloat(t, 'f', 1, 64), 64)
	return t
}

func RamdomCode() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(9999999999999999))
}
