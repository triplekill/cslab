package main

import (
	"io"
	"fmt"
	"bytes"
	"os"
	"strconv"
)

func main(){
	//r := strings.NewReader("some io.Reader stream to be read\n")
	//
	//if _, err := io.Copy(os.Stdout, r); err != nil {
	//	log.Fatal(err)
	//}

	//r := strings.NewReader("some io.Reader stream to be read\n")
	//
	//var buf1, buf2 bytes.Buffer
	//w := io.MultiWriter(&buf1, &buf2, os.Stdout, os.Stderr)
	//w.Write([]byte("teste\n"))
	//fmt.Print(buf1.String())
	//fmt.Print(buf2.String())
	//if _, err := io.Copy(w, r); err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Print(buf1.String())
	//fmt.Print(buf2.String())
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some text to be read\n")
		w.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	fmt.Print(buf.String())
	io.WriteString(os.Stdout, "Hello World")

	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}
}
