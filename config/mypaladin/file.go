package mypaladin

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	defaultChSize = 10
)

type file struct {
	values *Map
	// 该map用于存放每次 热更配置后的内容， 然后在将该内容回写到values中
	rawVal map[string]*Value

	watchChs map[string][]chan Event
	mx       sync.Mutex
	wg       sync.WaitGroup

	base string
	done chan struct{}
}

func NewFile(base string) (Client, error) {

	// 先将路径中的分隔符 转成对应系统的
	base = filepath.FromSlash(base)
	// 读取路径中的所有文件 该paths 里都是完整的路径(不一定是绝对路径，base给的绝对路径，这里就是绝对路径)
	paths, err := readAllPaths(base)
	if err != nil {
		return nil, err
	}
	rawVal, err := loadValuesFromPaths(paths)
	if err != nil {
		return nil, err
	}

	valMap := &Map{}

	valMap.Store(rawVal)
	fc := &file{
		values:   valMap,
		rawVal:   rawVal,
		watchChs: make(map[string][]chan Event),

		base: base,
		done: make(chan struct{}, 1),
	}

	fc.wg.Add(1)
	go fc.daemon()

	return fc, nil
}

func (f *file) Get(key string) *Value {
	return f.values.Get(key)
}

func (f *file) GetAll() *Map {
	return f.values
}

func (f *file) Close() error {
	f.done <- struct{}{}
	f.wg.Wait()
	return nil
}

func (f *file) WatchEvent(ctx context.Context, keys ...string) <-chan Event {
	f.mx.Lock()
	defer f.mx.Unlock()

	ch := make(chan Event, defaultChSize)

	for _, key := range keys {
		f.watchChs[key] = append(f.watchChs[key], ch)
	}
	return ch
}

func (f *file) daemon() {
	defer f.wg.Done()

	fswatcher, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Printf("create file watcher fail! reload function will lose efficacy error: %s\n", err)
		return
	}

	if err = fswatcher.Add(f.base); err != nil {
		fmt.Printf("create fsnotify for base path %s fail %s, reload function will lose efficacy", f.base, err)
		return
	}

	fmt.Printf("start watch filepath : %s", f.base)

	for event := range fswatcher.Events {
		switch event.Op {
		case fsnotify.Create, fsnotify.Write:
			f.reloadFile(event.Name)
		case fsnotify.Chmod:
		default:
			fmt.Printf("unsupport event %s ingored\n", event.Name)
		}
	}
}

func (f *file) reloadFile(name string) {
	if isHiddenFile(name) {
		return
	}
	// 这里先等一会 以防内容还没有回写到 文件中
	time.Sleep(100 * time.Millisecond)
	key := filepath.Base(name)
	val, err := loadValue(name)
	if err != nil {
		fmt.Printf("load file %s error: %s, skipped", name, err)
		return
	}

	f.rawVal[key] = val
	f.values.Store(f.rawVal)

	f.mx.Lock()
	chs := f.watchChs[key]
	f.mx.Unlock()

	for _, ch := range chs {
		select {
		case ch <- Event{Event: EventUpdate, Key: key, Value: val.raw}:
		default:
			fmt.Printf("event channel full discard file %s update event", name)
		}
	}

}

func readAllPaths(base string) ([]string, error) {

	fi, err := os.Stat(base)

	if err != nil {
		return nil, fmt.Errorf("check local config file fail! error : %s", err)
	}

	var paths []string

	if fi.IsDir() {
		files, err := ioutil.ReadDir(base)
		if err != nil {
			return nil, fmt.Errorf("read dir %s  errors :%s", base, err)
		}

		for _, file := range files {
			if !file.IsDir() && !isHiddenFile(file.Name()) {
				paths = append(paths, path.Join(base, file.Name()))
			}

		}
	} else {
		paths = append(paths, base)
	}
	return paths, nil
}

func loadValuesFromPaths(paths []string) (map[string]*Value, error) {
	// var err error

	values := make(map[string]*Value, len(paths))

	for _, fPath := range paths {

		val, err := loadValue(fPath)
		if err != nil {
			return nil, err
		}
		values[path.Base(fPath)] = val
	}
	return values, nil
}

func loadValue(fPath string) (*Value, error) {
	data, err := ioutil.ReadFile(fPath)

	if err != nil {
		return nil, err
	}
	content := string(data)
	return &Value{val: content, raw: content}, nil

}

func isHiddenFile(name string) bool {
	// base 返回路径中的最后一个元素  比如 /config/test.go  这里就会返回test.go
	return strings.HasPrefix(filepath.Base(name), ".")
}
