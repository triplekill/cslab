package main

import "fmt"

type Name string

func (n Name) String() string {
	return "String:" + string(n)
}

type A struct {
	a string
}

func (a A) String() string {
	return fmt.Sprintf("A.a:%v", a.a)
}

func main() {

	开心 := "中文开心"
	fmt.Println(开心)
	s := Name("test")
	a := A{}
	a.a = string(s)
	fmt.Printf("s type:%T, val:%v\n", s, s)
	fmt.Println(s + s)
	fmt.Printf("%T\n", s+s)

	utf8string := "张u太"
	fmt.Println(utf8string == "张u太")
	fmt.Println("获得字符串索引的字符的正确方法", string([]rune(utf8string)[0]))
	fmt.Println("获得字符串索引的字符的错误方法", string(utf8string[0]))
	for _, i := range utf8string {
		fmt.Printf("%T %5v %5v\n", i, i, string(i))
	}
	fmt.Println(string(100) == "d")
	print(string(utf8string[0])) //获得的不是第一个字符
	print("\U00010B90\n")
	tmap := map[string]int{}
	println(tmap["te"])
	println(len(utf8string))
	println(len([]rune(utf8string)))
}
