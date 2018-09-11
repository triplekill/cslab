package main

import (
	"fmt"
	"log"
	"sort"
)

type Any map[string]interface{}

func Unmarshal(v interface{}) {
	switch t := v.(type) {
	case map[string]interface{}:
		any := Any(t)
		log.Println(any)
	default:
		log.Printf("%T %v", t, t)

	}
}

// http://docscn.studygolang.com/doc/effective_go.html#%E6%8E%A5%E5%8F%A3%E4%B8%8E%E7%B1%BB%E5%9E%8B
func main() {
	var val map[string]interface{}
	val = map[string]interface{}{"A": "a"}
	Unmarshal(val)
	any := Any(val)
	log.Printf("%T %v", any, any)
}

// type
type Sequence []int

// Methods required by sort.Interface.
// sort.Interface 所需的方法。
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing.
// 用于打印的方法 - 在打印前对元素进行排序。
func (s Sequence) String() string {
	sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
	// return fmt.Sprint([]int(s))  // 类型转换
}

func assertType() string {
	var value interface{}
	value = Sequence([]int{2, 3, 1, 4})
	if str, ok := value.(string); ok {
		return str
	} else if str, ok := value.(fmt.Stringer); ok {
		return str.String()
	}
	return ""
}
