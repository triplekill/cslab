package main

import (
			"fmt"
	"reflect"
)

func main() {
	//var tss = make(map[string]string)
	//tss["key"] = "value"
	//println(tss)
	//tmapany(&tss)
	//
	//var b bytes.Buffer // A Buffer needs no initialization.
	//b.Write([]byte("Hello "))
	//fmt.Fprintf(&b, "world!")
	//b.WriteTo(os.Stdout)
	//
	//l := list.New()
	//e4 := l.PushBack(4)
	//e1 := l.PushFront(1)
	//l.InsertBefore(3, e4)
	//l.InsertAfter(2, e1)
	//
	//// list
	//// Iterate through list and print its contents.
	//for e := l.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}
	//
	//str := `test"' 'IOIJwo ijo jj tes"`
	//println(str)
	//fmt.Println(str)
	//Pr := "stdout"
	//pr := "stderr"
	//println(pr)
	//fmt.Println(Pr)
	//fmt.Println(reflect.TypeOf(println))
	fmt.Println(reflect.TypeOf(fmt.Println))
}

func tmapany(m interface{}) {
	fmt.Println(m)
}
