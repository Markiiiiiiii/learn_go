package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
)

const (
	APIURL = "https://way.jd.com/he/freeweather?appkey=79d6586692c7cb8127acf38815f8f6c8&city="
)

var (
	city, URL string
)

// type Weathdata struct {
// 	Code string `json:"code"`
// }

func main() {
	fmt.Println("请输入查询城市名：")
	fmt.Scanln(&city)
	for {
		if city != "" {
			getinfo(city)
			break
		} else {
			fmt.Println("未收到具体城市名请再次输入")
			break
		}
	}
	// main()
}

func getinfo(cityname string) {
	URL := APIURL + cityname
	queryUrl := fmt.Sprintf("%s", URL)
	resp, err := http.Get(queryUrl)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := simplejson.NewJson(body)
	infomap, _ := res.Get("result").Get("HeWeather5").GetIndex(0).Get("aqi").Get("city").Get("qlty").String()
	fmt.Println(infomap)
	// codeen, _ := res.Get("code").String()
	// fmt.Println(citys)
}
