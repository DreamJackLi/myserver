package log

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

var patternMap = map[string]func(map[string]interface{}) string{

	"T": longTime,
	"t": shortTime,
	"D": longData,
	"d": shortData,
	"L": keyFactory(_level),
	"f": keyFactory(_source),
	"M": message,
}

func newPatternRender(format string) Render {
	p := &pattern{
		bufPool: sync.Pool{New: func() interface{} { return &bytes.Buffer{} }},
	}
	// 保存了 所有的保留字符
	b := make([]byte, 0, len(format))
	/*
		Format:格式都是 %T 等这类的
		若有别的字符  都要保留到输出中

		这里先去检查是不是 % 若不是则认为是保留字符，加入b中
		若是  则判断当前i是不是最后一个，因为这里要去取i+1，i+1就是patternFunc,若是最后一个则也加入b中
		然后去取 i+1 ，将i+1到patternMap中取，若取不到对应的 patternFunc 则认为是保留字符 加入到b中
			若可以取到则稍等
		然后将b中的内容 用textFactory封装 加入到p.funcs中 ,清空b
		然后在将 i+i取到的 patternFunc 加入到p.func中
		i++  跳过当前字符，当前字符其实就是 % , 然后就到了for中的i++，由于i++就是 i+1，而i+1我们已经处理过了所以跳过也没问题


		这里先处理 b  然后在将处理p  是为了保证 pattern前面的保留字符可以先执行 ,不会乱了格式
	*/
	for i := 0; i < len(format); i++ {
		if format[i] != '%' {
			b = append(b, format[i])
			continue
		}

		if i+1 >= len(format) {
			b = append(b, format[i])
		}

		f, ok := patternMap[string(format[i+1])]

		if !ok {
			b = append(b, format[i])
		}

		if len(b) != 0 {
			p.funcs = append(p.funcs, textFactory(string(b)))
			b = b[:0]
		}

		p.funcs = append(p.funcs, f)
		i++
	}

	if len(b) != 0 {
		p.funcs = append(p.funcs, textFactory(string(b)))
	}

	return p
}

type pattern struct {
	funcs   []func(map[string]interface{}) string
	bufPool sync.Pool
}

func (p *pattern) Render(w io.Writer, d map[string]interface{}) error {
	buf := p.bufPool.Get().(*bytes.Buffer)

	defer func() {
		buf.Reset()
		p.bufPool.Put(buf)
	}()

	for _, f := range p.funcs {
		buf.WriteString(f(d))
	}

	_, err := buf.WriteTo(w)
	return err
}

func (p *pattern) RenderString(d map[string]interface{}) string {
	buf := p.bufPool.Get().(*bytes.Buffer)

	defer func() {
		buf.Reset()
		p.bufPool.Put(buf)
	}()

	for _, f := range p.funcs {
		buf.WriteString(f(d))
	}

	return buf.String()
}

func longTime(map[string]interface{}) string {
	return time.Now().Format("15:04:05.000")
}

func shortTime(map[string]interface{}) string {
	return time.Now().Format("15:04")
}

func longData(map[string]interface{}) string {
	return time.Now().Format("2006/01/02")
}

func shortData(map[string]interface{}) string {
	return time.Now().Format("01/02")
}

func keyFactory(key string) func(map[string]interface{}) string {
	return func(d map[string]interface{}) string {
		if v, ok := d[key]; ok {
			if s, ok := v.(string); ok {
				return s
			}
			return fmt.Sprint(v)
		}

		return ""
	}
}

func textFactory(text string) func(map[string]interface{}) string {
	return func(map[string]interface{}) string {
		return text
	}
}

func message(d map[string]interface{}) string {
	// var m string
	var s []string
	for k, v := range d {
		if isInternalKey(k) {
			continue
		}
		s = append(s, fmt.Sprintf("%s=%v", k, v))
	}
	// s = append(s, m)
	return strings.Join(s, " ")
}

func isInternalKey(k string) bool {
	switch k {
	case _level, _time, _source, _levelValue:
		return true
	}
	return false
}
