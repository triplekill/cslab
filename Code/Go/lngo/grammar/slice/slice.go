package main

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"unsafe"
)

// SortFloat64FastV2 Fast version float sort: conv float to int, then sort
func SortFloat64FastV2() {
	a := []float64{2.2, 1.1, 3.3}
	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	// 以int方式给float64排序
	sort.Ints(c)
	fmt.Println(a)
}

func app() {
	var a []int
	a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	fmt.Printf("%#v\n", a)
	a = []int{1, 2, 3}
	a = append([]int{0}, a...)          // 在开头添加1个元素
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	fmt.Printf("%#v\n", a)

	i := 2
	a = append(a[:i], append([]int{10, 11, 12}, a[i:]...)...) // 在第i个位置插入切片
	fmt.Printf("%#v\n", a)

	x := 13
	a = append(a, 0)     // 切片扩展1个空间
	copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置
	a[i] = x             // 设置新添加的元素
	fmt.Printf("%#v\n", a)
}

func slicet() {
	h := []string{"a", "b", "c"}
	t := h[0:0]
	log.Println(t)
}

func main() {
	app()
	SortFloat64FastV2()
	slicet()
}
