package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	//PART 1
	// resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=test")
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	return
	// }
	// b, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	return
	// }
	// fmt.Println(string(b))
	// resp.Body.Close()

	//PART2
	data := url.Values{} //对url进行编码转移
	urlObjs, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "张三")
	queryStr := data.Encode() //url encode之后的url
	urlObjs.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObjs.String(), nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(resp)
}
