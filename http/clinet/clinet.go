package main

import (
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
	uslstr := url.Parse("http:///127.0.0.1:9090/xxx/")
	data.Set("name", "张三")
	urlstr := data.Encode() //url encode之后的url
	req, err := http.NewRequest("GET", urlstr, nil)
}
