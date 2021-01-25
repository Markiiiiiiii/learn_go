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

//net的HTTP
func main() {
	http.HandleFunc("/post/", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
