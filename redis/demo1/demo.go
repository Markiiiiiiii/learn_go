package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisDb *redis.Client

func initRedis() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "35.201.213.247:6379",
		Password: "",
		DB:       0,
	})

	str, err := redisDb.Ping().Result()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(str)
	return
}
func redisDemo() {
	err := redisDb.Set("name", "张三", 0).Err()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	val, err := redisDb.Get("name").Result()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(val)
	//取一个不存在的键值错误处理
	val2, err := redisDb.Get("age").Result()
	if err == redis.Nil {
		fmt.Println("this Key is not ")
	} else if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(val2)
}
func redisZsetDemo() (err error) {
	zsetkey := "rank"
	languages := []redis.Z{
		redis.Z{Score: 99, Member: "PHP"},
		redis.Z{Score: 98, Member: "Golang"},
		redis.Z{Score: 96, Member: "C"},
		redis.Z{Score: 92, Member: "Python"},
		redis.Z{Score: 82, Member: "JS"},
	}
	redisDb.ZAdd(zsetkey, languages...)
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// redisDemo()

	redisZsetDemo()

}
