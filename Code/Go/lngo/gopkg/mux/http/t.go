package main

import (
	"net/http"
	"fmt"
)



func main(){
	http.HandleFunc("/hi", hi)
	http.ListenAndServe(":8080", nil)
	fmt.Println("test http")
}

func hi (res http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Fprintf(res, "hi")
}
