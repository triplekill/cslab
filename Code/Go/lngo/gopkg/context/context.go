package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

//	简单使用context.WithTimeout的例子
func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		fmt.Println("http request failed,err:", err)
		return
	}
	// 如果请求成功了会将数据存入到管道中
	go func() {
		resp, err := client.Do(req)
		pack := Result{resp, err}
		c <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		fmt.Println("timeout!")
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("server response:%s", out)
	}
	return

}

func add(ctx context.Context, a, b int) int {
	traceId := ctx.Value("trace_id").(string)
	fmt.Printf("trace_id:%v\n", traceId)
	return a + b
}

func calc(ctx context.Context, a, b int) int {
	traceId := ctx.Value("trace_id").(string)
	fmt.Printf("trace_id:%v\n", traceId)
	//再将ctx传入到add中
	return add(ctx, a, b)
}
func main() {
	process()
	ctx := context.WithValue(context.Background(), "trace_id", "123456")
	calc(ctx, 20, 30)
}
