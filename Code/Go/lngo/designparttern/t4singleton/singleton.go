package t4singleton

import (
	"sync"
)

// One
type singleton struct {
	A int
}

var (
	s  *singleton
	mu sync.Mutex
)

func New() *singleton {
	if s == nil {
		mu.Lock()
		if s == nil { // 双重检查加锁
			s = &singleton{A: 10}
		}
		mu.Unlock()
	}
	return s
}

// Two
var Singleton = singleton{}

// Three init() 中进行初始化
var is *singleton

func init() {
	s = &singleton{A: 10}
}

// Four 在cache中获得缓存，不设置失效时间，在get的时候进行一次初始化，之后进行复用

// Five onece函数进行初始化
