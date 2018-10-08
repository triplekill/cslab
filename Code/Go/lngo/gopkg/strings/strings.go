package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())

	bm := make([]byte, 0)
	var tmp string = "abcdef"
	bm = append(bm, tmp...)
	fmt.Println(bm)
	fmt.Println(reflect.TypeOf(tmp))
	fmt.Println(strings.Compare("e", "cbbb"))
	neadsplit := "  foo bar  baz   "
	fmt.Printf("Fields are: %q\n", strings.Fields(neadsplit))
	fmt.Printf("%q\n", strings.Split(neadsplit, " "))
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))

	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))

	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))

}
