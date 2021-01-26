package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	str, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	w.Write(str)
}
func f2(w http.ResponseWriter, r *http.Request) {
	//对于GET请求（query param），请求体中没有数据
	queryParam := r.URL.Query()
	// for k, v := range queryParam {
	// 	fmt.Println(k)
	// 	fmt.Println(v)
	// }
	fmt.Println(queryParam.Get("name"))
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

//net的HTTP
func main() {
	http.HandleFunc("/post/", f1)
	http.HandleFunc("/xxx/", f2)

	http.ListenAndServe("127.0.0.1:9090", nil)
}
